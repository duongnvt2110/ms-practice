# Event Booking MVP ERD (Sample)

```mermaid
erDiagram
  USER {
    uuid id PK
    string email
    string full_name
    string phone
    datetime created_at
  }

  ORGANIZER {
    uuid id PK
    uuid user_id FK
    string display_name
    datetime created_at
  }

  VENUE {
    uuid id PK
    string name
    string address
    string city
    string country
    int capacity
  }

  EVENT {
    uuid id PK
    uuid organizer_id FK
    uuid venue_id FK
    string title
    string description
    string event_type  // music | sports | films | festival | concert
    datetime starts_at
    datetime ends_at
    string status    // draft | published | cancelled
    datetime created_at
  }

  SEAT_MAP {
    uuid id PK
    uuid event_id FK
    string name
    string version
  }

  SEAT {
    uuid id PK
    uuid seat_map_id FK
    string section
    string row
    string number
    string status // available | held | booked
  }

  PRICE_TIER {
    uuid id PK
    uuid event_id FK
    string name
    int price_cents
    string currency
    string inventory_type // fixed_seat | general_admission
    string seat_section // optional mapping for fixed_seat
    int ga_capacity // optional for general_admission
  }

  BOOKING {
    uuid id PK
    uuid user_id FK
    uuid event_id FK
    string status // pending | paid | cancelled | refunded
    int total_amount_cents
    string currency
    datetime created_at
    datetime expires_at // seat hold expiry
  }

  BOOKING_ITEM {
    uuid id PK
    uuid booking_id FK
    uuid price_tier_id FK
    int quantity
    int unit_price_cents
  }

  BOOKING_SEAT {
    uuid id PK
    uuid booking_id FK
    uuid seat_id FK
    int price_cents
  }

  SEAT_HOLD {
    uuid id PK
    uuid booking_id FK
    uuid seat_id FK
    datetime held_at
    datetime expires_at
    string status // active | released | converted
  }

  PAYMENT {
    uuid id PK
    uuid booking_id FK
    string provider // card
    string status // initiated | succeeded | failed
    string provider_ref
    int amount_cents
    string currency
    datetime created_at
  }

  TICKET {
    uuid id PK
    uuid booking_id FK
    uuid booking_item_id FK
    uuid booking_seat_id FK
    uuid seat_id FK
    string code
    string delivery_status // sent | failed
    datetime issued_at
  }

  NOTIFICATION {
    uuid id PK
    uuid user_id FK
    uuid booking_id FK
    string channel // email | web
    string type // ticket_issued | refund_update
    string status // pending | sent | failed
    datetime created_at
  }

  REFUND_REQUEST {
    uuid id PK
    uuid booking_id FK
    uuid user_id FK
    uuid organizer_id FK
    string status // requested | approved | rejected
    string reason
    datetime created_at
    datetime decided_at
  }

  USER ||--o{ BOOKING : makes
  USER ||--o{ NOTIFICATION : receives
  USER ||--o{ REFUND_REQUEST : submits
  ORGANIZER ||--o{ EVENT : creates
  ORGANIZER ||--o{ REFUND_REQUEST : decides
  VENUE ||--o{ EVENT : hosts
  EVENT ||--o{ PRICE_TIER : has
  EVENT ||--o{ BOOKING : for
  EVENT ||--o{ SEAT_MAP : uses
  SEAT_MAP ||--o{ SEAT : contains
  BOOKING ||--o{ BOOKING_ITEM : includes
  PRICE_TIER ||--o{ BOOKING_ITEM : priced_by
  BOOKING ||--o{ BOOKING_SEAT : includes
  SEAT ||--o{ BOOKING_SEAT : assigned
  BOOKING ||--o{ SEAT_HOLD : reserves
  SEAT ||--o{ SEAT_HOLD : held
  BOOKING ||--o{ PAYMENT : paid_by
  BOOKING ||--o{ TICKET : issues
  BOOKING_ITEM ||--o{ TICKET : issues
  BOOKING_SEAT ||--o{ TICKET : issues
  BOOKING ||--o{ NOTIFICATION : triggers
  BOOKING ||--o{ REFUND_REQUEST : may_have
```
