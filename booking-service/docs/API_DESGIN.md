# API Design (Booking Service)

## Current HTTP APIs
These routes map directly to the handlers in `pkg/handler/http/booking`.

### POST `/v1/bookings`
- **Headers**
  - `Idempotency-Key` *(optional)* – if provided we treat duplicate requests as a no-op and return the existing booking.
- **Request Body**
  ```json
  {
    "user_id": 123,
    "event_id": 456,
    "status": "pending",            // optional; defaults to pending
    "total_price": 120000,          // optional; recomputed when zero
    "logs": "{}",                   // optional
    "items": [
      {
        "ticket_type_id": 1,
        "qty": 2,
        "price": 60000
      }
    ]
  }
  ```
- **Behavior**
  - Server fills missing defaults (booking code, idempotency key, total price, number of seats, hold/expiry timestamps).
  - After persistence we emit a `BookingOrdered` event to Kafka topic `booking.events`.
- **Response**
  ```json
  {
    "status": "success",
    "data": 789   // booking ID
  }
  ```

### GET `/v1/bookings`
- **Query params**
  - `user_id` *(optional)* filters bookings by owner.
- **Response**
  Returns an array of bookings with nested `items` and `booking_users` as stored in the database.

### GET `/v1/bookings/:id`
- Returns a single booking with nested line items and booking users.
- `404` when the record is not found.

## Event Flow (Current State)
- When a booking is created we publish `BookingOrdered` payloads to `booking.events`.
- Payment service consumes the topic and (in a separate service) publishes `PaymentSucceeded` / `PaymentFailed` events to `payments.events`.
- **Gap:** booking service does not yet listen to `payments.events`; status remains “pending” until we implement the consumer.

## Planned Extensions
The following APIs/state transitions are referenced in requirements but are **not** implemented yet:

| Planned Feature | Description | Notes |
| --------------- | ----------- | ----- |
| `POST /v1/bookings/hold` | Reserve seats before payment, return hold expiry (15 min SLA). | Needs seat inventory integration + cron/worker to release expired holds. |
| `POST /v1/bookings/{id}/information` | Capture traveller/contact info. | Persist into `booking_users`, validate contacts. |
| `POST /v1/bookings/{id}/confirm` | Finalize booking after payment success. | Should verify PaymentSucceeded event, update status, emit ticket creation event. |
| `POST /v1/bookings/{id}/cancel` | Cancel hold/booking (manual or automatic). | Update status + emit `BookingCancelled`; refund flow TBD. |
| Payment consumer | Subscribe to `payments.events` to move booking status to `confirmed`/`failed`. | Required before exposing confirm/cancel APIs. |

## Database Schema Design
`bookings`
  - id 
  - idempotency_key
  - event_id
  - user_id
  - holded_at
  - expired_at
  - total_price
  - status
  - log
  - created_at
  - updated_at
`booking_items` 
  - id
  - booking_id
  - ticket_type_id
  - qty
  - price 
  - created_at
  - updated_at
`booking_users`
  - id 
  - booking_id
  - user_id
  - email 
  - mobile_number
  - address
  - created_at
  - updated_at

> TODO (tracked here for visibility): add ER diagram, document enum values for `status`, and describe clean-up job that expires holds once the workflow is implemented.
