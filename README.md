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
| 9   | Event Service   | event-service      | localhost | 8007 |             |
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
- 
#### Database Schema Desgin 
- 
### User Service
#### Goal
- User Infos
- User settings
#### API Design
- 
#### Database Schema Desgin 
- 
#### Techstack: 
- Golang, Mux golang
### Ticket Service 
#### Goal
- User Infos
- User settings
#### API Design
- 
#### Database Schema Desgin 
- 
#### Techstack: 
- Golang, Mux golang
### Booking Service 
#### Goal
- User Infos
- User settings
#### API Design

#### Techstack: 
- Golang, Mux golang
### Payment Service 
#### Goal
- Payments
#### API Design
- 
#### Database Schema Desgin 
- 

#### Techstack: 
- Golang, Mux golang
### Noti Service 
#### Goal
- Deliver email, push, and in-app notifications
- Persist notification jobs with retries
#### API Design
- `GET /health`
#### Database Schema Design 
- `notification_jobs`
- `notification_templates` (future)
#### Techstack: 
- Golang, Gin, Kafka, MySQL
### Event Service
#### API Design
- 
#### Database Schema Desgin 
- 
# CDC Service
- Considering ...
# Saga 
## Choreography Pattern
The ticketing workflow follows a choreography saga driven by Kafka events:

1. **Booking Service**
   - Persists the booking and publishes a `BookingOrdered` event to `booking.events`.
   - Listens to `payments.events` and updates the booking status to `confirmed` or `failed`.
2. **Payment Service**
   - Consumes `booking.events` to create payment records.
   - Emits `PaymentSucceeded` or `PaymentFailed` events to `payments.events`.
3. (Future extension) Ticket and notification services can join the saga by consuming the payment events.

### Event Types
- Booking Service  
  - `BookingOrdered`
  - `BookingCancelled`
- Payment Service  
  - `PaymentSucceeded`
  - `PaymentFailed`


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
