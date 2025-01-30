# Architecture 
## Booking Ticket System
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
| 2   | Ticket Service  | ticket-service     | localhost | 8001 |             |
| 3   | Payment Service | payment-serivce    | localhost | 8002 |             |
| 4   | Booking Service | order-service      | localhost | 8003 |             |
| 5   | User Service    | user-serivce       | localhost | 8004 |             |
| 6   | Noti Service    | Noti-serivce       | localhost | 8005 |             |
| 8   | Auth Service    | Noti-serivce       | localhost | 8005 |             |
| 9   | Web             | web                | localhost | 8888 |             |
## Detail
### API Gateway
#### Techstack: 
- Golang, Echo framework
### User Service
#### Techstack: 
- Golang, Mux golang
### Auth Service 
#### Techstack: 
- Golang, Gin framework


# CDC Service
- Considering ...
# Messeage Queue
## Kafka 
### Topic 
| Topic Name | Message Type | Producer | Consumer | Description |
| ---------- | ------------ | -------- | -------- | ----------- |
|            |              |          |          |             |
### Usecases 
#### CLI for creating and deleting topic 
#### Format payload mesage

# Data Models 
### User
  - id
  - username 
  - encrypted_password
  - email
### Booking
  - id 
  - user_id 
  - name
  - date
  - ticket_id
### Ticket
  - id 
  - name
  - type 
  - price
  - tax
### Payment
  - id
  - user_id
  - booking_id
### Noti 


# 