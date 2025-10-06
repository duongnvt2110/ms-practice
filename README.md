# Architecture 
## Ticketing System
https://systemdesignschool.io/problems/ticketmaster/solution
## Requirements 
### Functional Requirements
- Users can search for and book resources such as (concert, sports,....)
- Availability is updated in real-time.
- Payments are processed securely.
- Notifications are sent promptly.
### Non-Functional Requirements

# Services 
## List
| No. | Service Name    | Directory Name     | Host      | Port | Description |
| --- | --------------- | ------------------ | --------- | ---- | ----------- |
| 1   | API Gateway     | api-gatewa-service | localhost | 8000 |             |
| 2   | Auth Service    | auth-service       | localhost | 8001 |             |
| 3   | User Service    | user-serivce       | localhost | 8002 |             |
| 4   | Ticket Service  | ticket-service     | localhost | 8003 |             |
| 5   | Booking Service | booking-serivce    | localhost | 8004 |             |
| 6   | Payment Service | payment-serivce    | localhost | 8005 |             |
| 8   | Noti Service    | noti-serivce       | localhost | 8005 |             |
| 9   | Event Service   | event-service      | localhost | 8003 |             |
| 10  | FrontEnd        | Frontend           | localhost | 8888 |             |
### Overrall the flow 
- User authenticate though by the AuthService and get user's information by User Service.
- User books a the ticket via the booking service and the booking services store the booking information after send an booking event to the booking.events. the payment service consumes the evvent from that to process payment or handle other exception
- If the payment succeeded this emit an event to the payment.events and the booking services pull this event to generate the ticket. It also send the noti to the users 
- The booking information will get from the event services.
## Detail
### API Gateway
- Proxy
- Forward the request to the correct services 
- Loadbalancer
#### Techstack: 
- Golang, Echo framework
### Auth Service
#### Goal 
- Authen the system 
- Middlware 
- Validate token 
- Rotate token 
#### API Design
- [POST] `v1/login`
  - Request 
    - Body:
      - username|email
      - password
  - Response
    - access_token
    - refresh_token  
- [POST] `v1/logout`
- [POST] `v1/register`
  - Request
    - Body:
      - email 
      - user_name
      - first_name
      - last_name
      - password
      - mobile_phone
#### Database Schema Desgin 
- AuthProfiles (auth_profiles)
  - id 
  - email
  - username
  - password (hashing)
  - created_at
  - updated_at
### User Service
#### Goal
- User Infos
- User settings
#### API Design
- [GET] `v1/users`
  - Description
    - Get list users
  - Request
    - Query params 
      - next_token 
      - previous_token
  - Response
    - list user
- [GET] `v1/users/{id}`
  - Description
    - Get user by id
  - Response
    - list user
- [PUT] `v1/users/{id}`
  - Request body:
    - user_name
    - first_name
    - last_name
    - birth_day
    - mobile_phone
#### Database Schema Desgin 
- Users (users)
  - id 
  - user_name
  - first_name
  - last_name
  - birth_day
  - mobile_phone
  - created_at
  - updated_at
#### Techstack: 
- Golang, Mux golang
### Ticket Service 
#### Goal
- User Infos
- User settings
#### API Design
- [GET] `v1/tickets`
  - Description
    - Get list tickets
  - Request
    - Query params 
      - next_token 
      - previous_token
  - Response
    - list user
- [GET] `v1/tickets/{id}`
  - Description
    - Get ticket details
  - Response
    - list user
- [PUT] `v1/tickets/{id}` -> CMS
  - Description
    - Update ticket info
  - Request body:
#### Database Schema Desgin 
- Tickets (tickets) 
  - id 
  - user_id 
  - booking_id
  - booking_item_id
  - event_type_id
  - event_id
  - issued_at
  - code
  - qr_payload
  - status 
  - created_at
  - updated_at
#### Techstack: 
- Golang, Mux golang
### Booking Service 
#### Goal
- User Infos
- User settings
#### API Design
- [GET] `v1/bookings`
  - Description
    - Get list
  - Request
    - Query params 
      - next_token 
      - previous_token
  - Response
    - list user
- [GET] `v1/bookings/{id}`
  - Description
    - Get user by id
  - Response
    - list user
- [POST] `v1/bookings`
  - Request body:
    - event_id
    - number_slot
#### Database Schema Desgin 
- Bookings (bookings)
  - id 
  - user_id
  - event_id 
  - status 
  - quantity
  - prices
  - idempotency_key
  - status
  - logs
  - created_at
  - updated_at
- BookingItems 
  - id
  - booking_id
  - event_type_id 
  - qty 
  - unit_price
  - currency 
  - created_at
  - updated_at 
  - 
#### Techstack: 
- Golang, Mux golang
### Payment Service 
#### Goal
- User Infos
- User settings
#### API Design
- [Get] `v1/payments`
  - Description
    - Get list users
  - Request
    - Query params 
      - next_token 
      - previous_token
  - Response
    - list user
- [GET] `v1/payments/{id}`
  - Description
    - Get user by id
  - Response
    - list user
- [POST] `v1/payments`
  - Request body:
    - order_id
    - price
#### Database Schema Desgin 
- Payments (payments)
  - id 
  - user_id
  - booking_id
  - prices
  - method
  - idempotency_key
  - status
  - created_at
  - updated_at
- PaymentHistory (payment_histories)
  - payment_id
  - status
  - logs
  - paid_at
  - created_at
  - updated_at
#### Techstack: 
- Golang, Mux golang
### Noti Service 
#### Goal
- User Infos
- User settings
#### API Design
#### Database Schema Desgin 
#### Techstack: 
- Golang, Mux golang
### Event Service
- Events (events)
  - id 
  - category_id
  - name
  - starts_at
  - ends_at
  - max_per_user
  - status
  - created_at
  - updated_at
- EventCategory (event_categories)
  - id 
  - name
    - workshop
  - type
    - free
    - paid
  - created_at
  - updated_at
- EventTypes
  - id 
  - event_id
  - name 
  - prices
  - currency
  - status 
  - capacity  
# CDC Service
- Considering ...
# Saga 
## Choreography Pattern
### Useccase
- A customer places an order in BookingService.
<!-- - OrderService saves the order and emits an OrderPlacedEvent.
- InventoryService listens for OrderPlacedEvent, and once it catches this event, it checks and reserves the stock. If stock is reserved successfully, it emits a StockReservedEvent.
- If stock isn't available, it emits a StockUnavailableEvent.
PaymentService listens for StockReservedEvent. Once it catches this event, it charges the customer.
- If payment is successful, it emits a PaymentSuccessEvent.
- If payment fails, it emits a PaymentFailedEvent.
OrderService listens for PaymentSuccessEvent and PaymentFailedEvent to update the order status accordingly.
- NotificationService listens to various events to notify the customer at different stages. -->
### Event 
#### BookingrService 
- BookingOrdered
- BookingConfirmed
- BokkingFailed
#### PaymentSerivce
- PaymentSucceeded
- PaymenFailed
#### TicketService 
- TicketPending 
- TicketCreated
- TicketRejected
### NotificationService 
- NotiPending 
- NotiCreated


# Messeage Queue
## Kafka 
### Topic 
| Topic Name | Message Type | Producer | Consumer | Description |
| ---------- | ------------ | -------- | -------- | ----------- |
| OrderEvent |              |          |          |             |
#### Usecases 
#### CLI for creating and deleting topic 
#### Format payload mesage
# gRPC command 
## Install 
### Setup protoc
```
https://protobuf.dev/installation/
```
### Setup golang binary
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Command 
#### 1. user-service 
```
protoc --go_out=./proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative \
    --proto_path=proto \
    ./proto/user.proto
```

## Testing 
#### GRPC
https://ghz.sh/ 
Example:
```
ghz --insecure \
  --proto ./proto/user.proto \
  --call gen.UserService.TestGracefulShutdown \
  -c 1 -n 1 \
  0.0.0.0:50001
```