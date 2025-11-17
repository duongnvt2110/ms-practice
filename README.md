# Architecture 
## Ticketing System
https://systemdesignschool.io/problems/ticketmaster/solution
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
| 8   | Noti Service    | noti-serivce       | localhost | 8006 |             |
| 9   | Catalog Service | catalog-service    | localhost | 8007 |             |
| 11  | FrontEnd        | Frontend           | localhost | 8888 |             |
### Overral the flow 
- User authenticate though by the AuthService and get user's information by User Service.
- User -> Choose the events -> choose the position of seats (optinal) ->  input the number of seats -> proceed booking the tickets -> confirm the booking information -> select the payment method -> payment proceed.
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
- [POST] `v1/auths/login`
  - Request 
    - Body:
      - username|email
      - password
  - Response
    - access_token
    - refresh_token  
- [POST] `v1/auths/refresh_token`
  - refresh_token
- [POST] `v1/logout`
  - refresh_token
- [POST] `v1/register`
  - Request
    - Body:
      - email 
      - username
      - password
      - mobile_number
#### Database Schema Desgin 
- AuthProfiles (auth_profiles)
  - id 
  - email
  - username
  - password (hashing)
  - created_at
  - updated_at
- auth_refresh_tokens
  - id
  - auth_profile_id
  - token
  - expired_at
  - creatd_at
  - updated_at
### User Service
#### Goal
- User Infos
- User settings
#### API Design
- [GET] `v1/users/me`
 - id
 - email
 - username
 - avatar
 - brithday
 - user_settings
#### Database Schema Desgin 
- Users (users)
  - id 
  - email
  - username
  - birth_day
  - avatar
  - mobile_number
  - created_at
  - updated_at
- `user_settings`
   - id
   - allow_noti
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
  - payment_id
  - ticket_type_id
  - code
  - qr_url
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
  "email": "user@example.com",
  "phone": "+84...",
  "address": "..."
```
- [POST] `v1/bookings/{id}/confirm`
```
  "email": "user@example.com",
  "phone": "+84...",
  "address": "..."
```
- [POST] `v1/bookings/{id}/cancel`
- [GET] /v1/bookings/{id}
#### Database Schema Desgin 
`bookings`
  - id 
  - idempotency_key
  - event_id
  - user_id
  - holded_at
  - expired_at
  - booking_code
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
#### API Design
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
#### Database Schema Desgin 
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
- BookingFailed
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
#### 2. payment-service 
```
protoc --go_out=./proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative \
    --proto_path=proto \
    ./proto/payment.proto
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