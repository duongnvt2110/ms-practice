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
| 5   | Order Service   | order-serivce      | localhost | 8004 |             |
| 6   | Payment Service | payment-serivce    | localhost | 8005 |             |
| 8   | Noti Service    | noti-serivce       | localhost | 8005 |             |
| 9   | FrontEnd        | Frontend           | localhost | 8888 |             |
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
### Event Service 
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
- Events (events)
  - id 
  - cate_id
  - event_name
  - event_price
  - event_date
  - total_slot
  - available_slot
  - reserved_slot
  - status
  - created_at
  - updated_at
- CategoryEvent (cate_events)
  - id 
  - cate_name
    - workshop
  - cate_type
    - free
    - paid
  - created_at
  - updated_at
#### Techstack: 
- Golang, Mux golang
### Order Service 
#### Goal
- User Infos
- User settings
#### API Design
- [GET] `v1/orders`
  - Description
    - Get list
  - Request
    - Query params 
      - next_token 
      - previous_token
  - Response
    - list user
- [GET] `v1/orders/{id}`
  - Description
    - Get user by id
  - Response
    - list user
- [POST] `v1/orders`
  - Request body:
    - event_id
    - number_slot
#### Database Schema Desgin 
- Tickets (tickets) 
  - id 
  - user_id 
  - event_id 
  - status 
  - created_at
  - updated_at
- Orders (Orders)
  - id 
  - user_id
  - ticket_id 
  - status 
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
- Payments (payments)
  - id 
  - user_id
  - order_id
  - amount
  - status
  - created_at
  - updated_at
- PaymentHistory (payment_histories)
  - pay_id
  - status
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

# CDC Service
- Considering ...
# Messeage Queue
## Kafka 
### Topic 
| Topic Name | Message Type | Producer | Consumer | Description |
| ---------- | ------------ | -------- | -------- | ----------- |
|            |              |          |          |             |
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
