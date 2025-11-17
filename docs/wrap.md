## Requirements
### Functional Requirements
- Users
 - Search and view the events (Liveshow music, concerts, sports, films and others)
 - Booking the events
   - Choose number of seats
   - Choose the position of seats (optional)
 - Checkout
   - Hold the seats in 15 minutes
   - Confirm information
   - Confirm the payment method
   - Proceed payment
 - Send notify
   - Ticket to the email
   - Noti in the website
- Organizer
 - Create/Delete/Edit the events



User -> Choose the events -> choose the position of seats (optinal) ->  input the number of seats -> proceed booking the tickets -> confirm the booking information -> select the payment method -> payment proceed.


- Auth Service: Authorize the users to use the system
- User service: 
 - User information
 - User settings
- Booking service
 - Create the booking
 - Hold the seats
 - Emit the booking.confirmed event to the `bookings.topics`
 - Receive the payment event from payment.topics and update the booking status
- Payment service
 - Receive the event from Booking Service
 - Proceed payment
 - Emit the payment.succeeded for the `payments.topics`
- Noti Service
 - Receive the payment event from payment.topics and send notify
- Ticket Service
 - Receive the payment event from payment.topics to generate ticket and send the ticket to users email.
- Catalog Service
 - Create/Update/Edit the events


# Auth Service
## API Design
- [POST] `v1/auths/login`
```
{
 email: string
 password: string
}
```
- [POST] `v1/auths/refresh_token`
```
{
 refresh_token: string
}
```
- [POST] `v1/auths/logout`
## DB Schema
`auth_profiles`
 - id
 - email
 - password
 - created_at
 - updated_at
`auth_refresh_tokens` 
- id
- auth_profile_id
- refresh_token
- expired_at
- creatd_at
- updated_at
- 
# User Service
## API Design
- [GET] `v1/users/me`
 - id
 - email
 - username
 - avatar
 - brithday
 - user_settings
## DB Schema
`users`
 - id
 - email
 - username
 - avatar
 - brithday
 - created_at
 - updated_at
r
# Booking Service
## API Design
- [POST] `v1/bookings/hold`
```
{
 event_id
 event_type_id
 qty
}
```
- [POST] `v1/bookings/{id}/information`
```
{
  "email": "user@example.com",
  "phone": "+84...",
  "address": "..."
}
```
## DB Schema
`bookings`
- id 
- event_id
- holded_at
- expired_at
- booking_code
- total_price
- status
- number_seats
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
- number_phone
- address
- created_at
- updated_at

# Event Service
## API Design
- [GET] `v1/events`
```
[
 {
   id: int
   name: string
   start_at: time
   end_at: time
   banner: string
   location
   status
 }
]
```
- [GET] `v1/events/{:id}`
```
{
 id: int
 name: string
 title: sttring
 start_at: time
 end_at: time
 banner: string
 location: string
 status: string
 ticket_types [
   {
     id: string
     position: int
     name: string
     description: string
     image_url: string
     status: string
     number_seats: int
     price: int
   }
 ]
}
```
## DB Schema
`events`
- id
- name
- title
- start_at
- end_at
- banner
- location
- status
- created_at
- updated_at

`ticket_types`
- id
- event_id
- position
- name
- description
- imageUrl
- status
- qty
- price
- currency
- sale_at
- sale_end
- created_at
- updated_at
- 
# Payment Service
## API Design
- [GET] `v1/payments/{payment_id}`
```
 {
   id: int
   bookings: 
   events
   status
 }
```
- [POST] `v1/payments/webhook?payment_code=xxx&booking_id=xxx&transaction_id=xxx`
## DB Schema
`payments` 
- id 
- payment_code
- user_id
- booking_id
- transaction_id (3rd)
- price
- status
- provider
- paid_at
- created_at
- updated_at

# Ticket Service
## API Design
- [GET] `v1/tickets?limit=xx&next_token=xxx`
```
[
 {
   id: int
   bookings: 
   events
   payments
   qr_url
 }
]
```
- [GET] `v1/ticket/{id}`
```
  {
   id: int
   bookings: 
   events
   payments
   qr_url
 }
```
## DB Schema
`tickets` 
- id 
- user_id
- booking_id
- payment_id
- qr_url
- status
- created_at
- updated_at