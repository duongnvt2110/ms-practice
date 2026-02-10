-- =========================
-- USERS / ORGANIZERS
-- =========================

CREATE TABLE users (
  id CHAR(36) PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  phone VARCHAR(50),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_users_email (email)
) ENGINE=InnoDB;

CREATE TABLE organizers (
  id CHAR(36) PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  display_name VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_organizers_user (user_id),
  CONSTRAINT fk_organizers_user
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB;

-- =========================
-- VENUES / EVENTS
-- =========================

CREATE TABLE venues (
  id CHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  city VARCHAR(120) NOT NULL,
  country VARCHAR(120) NOT NULL,
  capacity INT NOT NULL,
  CHECK (capacity >= 0)
) ENGINE=InnoDB;

CREATE TABLE events (
  id CHAR(36) PRIMARY KEY,
  organizer_id CHAR(36) NOT NULL,
  venue_id CHAR(36) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  event_type VARCHAR(30) NOT NULL,
  starts_at DATETIME NOT NULL,
  ends_at DATETIME NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK (event_type IN ('music','sports','films','festival','concert')),
  CHECK (status IN ('draft','published','cancelled')),
  CHECK (ends_at > starts_at),
  KEY idx_events_organizer (organizer_id),
  KEY idx_events_venue (venue_id),
  CONSTRAINT fk_events_organizer
    FOREIGN KEY (organizer_id) REFERENCES organizers(id),
  CONSTRAINT fk_events_venue
    FOREIGN KEY (venue_id) REFERENCES venues(id)
) ENGINE=InnoDB;

-- =========================
-- SEAT MAPS / SEATS
-- =========================

CREATE TABLE seat_maps (
  id CHAR(36) PRIMARY KEY,
  event_id CHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  version VARCHAR(50) NOT NULL,
  KEY idx_seat_maps_event (event_id),
  CONSTRAINT fk_seat_maps_event
    FOREIGN KEY (event_id) REFERENCES events(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE seats (
  id CHAR(36) PRIMARY KEY,
  seat_map_id CHAR(36) NOT NULL,
  section VARCHAR(60) NOT NULL,
  row_label VARCHAR(30) NOT NULL,
  seat_number VARCHAR(30) NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'available',
  CHECK (status IN ('available','held','booked')),
  UNIQUE KEY uk_seats_position (seat_map_id, section, row_label, seat_number),
  KEY idx_seats_status (status),
  CONSTRAINT fk_seats_map
    FOREIGN KEY (seat_map_id) REFERENCES seat_maps(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- PRICE TIERS
-- =========================

CREATE TABLE price_tiers (
  id CHAR(36) PRIMARY KEY,
  event_id CHAR(36) NOT NULL,
  name VARCHAR(120) NOT NULL,
  price_cents INT NOT NULL,
  currency CHAR(3) NOT NULL,
  inventory_type VARCHAR(30) NOT NULL,
  seat_section VARCHAR(60),
  ga_capacity INT,
  CHECK (price_cents >= 0),
  CHECK (inventory_type IN ('fixed_seat','general_admission')),
  CHECK (ga_capacity IS NULL OR ga_capacity >= 0),
  KEY idx_price_tiers_event (event_id),
  CONSTRAINT fk_price_tiers_event
    FOREIGN KEY (event_id) REFERENCES events(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- BOOKINGS
-- =========================

CREATE TABLE bookings (
  id CHAR(36) PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  event_id CHAR(36) NOT NULL,
  status VARCHAR(20) NOT NULL,
  total_amount_cents INT NOT NULL DEFAULT 0,
  currency CHAR(3) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at DATETIME,
  CHECK (status IN ('pending','paid','cancelled','refunded')),
  CHECK (total_amount_cents >= 0),
  KEY idx_bookings_user (user_id),
  KEY idx_bookings_event (event_id),
  CONSTRAINT fk_bookings_user
    FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_bookings_event
    FOREIGN KEY (event_id) REFERENCES events(id)
) ENGINE=InnoDB;

CREATE TABLE booking_items (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  price_tier_id CHAR(36) NOT NULL,
  quantity INT NOT NULL,
  unit_price_cents INT NOT NULL,
  CHECK (quantity > 0),
  CHECK (unit_price_cents >= 0),
  KEY idx_booking_items_booking (booking_id),
  CONSTRAINT fk_booking_items_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_booking_items_price_tier
    FOREIGN KEY (price_tier_id) REFERENCES price_tiers(id)
) ENGINE=InnoDB;

CREATE TABLE booking_seats (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  seat_id CHAR(36) NOT NULL,
  price_cents INT NOT NULL,
  CHECK (price_cents >= 0),
  UNIQUE KEY uk_booking_seat (seat_id),
  KEY idx_booking_seats_booking (booking_id),
  CONSTRAINT fk_booking_seats_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_booking_seats_seat
    FOREIGN KEY (seat_id) REFERENCES seats(id)
) ENGINE=InnoDB;

CREATE TABLE seat_holds (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  seat_id CHAR(36) NOT NULL,
  held_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at DATETIME NOT NULL,
  status VARCHAR(20) NOT NULL,
  CHECK (status IN ('active','released','converted')),
  CHECK (expires_at > held_at),
  KEY idx_seat_holds_seat (seat_id),
  CONSTRAINT fk_seat_holds_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_seat_holds_seat
    FOREIGN KEY (seat_id) REFERENCES seats(id)
) ENGINE=InnoDB;

-- =========================
-- PAYMENTS
-- =========================

CREATE TABLE payments (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  provider VARCHAR(30) NOT NULL,
  status VARCHAR(20) NOT NULL,
  provider_ref VARCHAR(255),
  amount_cents INT NOT NULL,
  currency CHAR(3) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK (provider IN ('card')),
  CHECK (status IN ('initiated','succeeded','failed')),
  CHECK (amount_cents >= 0),
  KEY idx_payments_booking (booking_id),
  CONSTRAINT fk_payments_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

-- =========================
-- TICKETS
-- =========================

CREATE TABLE tickets (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  booking_item_id CHAR(36),
  booking_seat_id CHAR(36),
  seat_id CHAR(36),
  code VARCHAR(120) NOT NULL,
  delivery_status VARCHAR(20) NOT NULL,
  issued_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK (delivery_status IN ('sent','failed')),
  UNIQUE KEY uk_tickets_code (code),
  CONSTRAINT fk_tickets_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_tickets_booking_item
    FOREIGN KEY (booking_item_id) REFERENCES booking_items(id)
    ON DELETE SET NULL,
  CONSTRAINT fk_tickets_booking_seat
    FOREIGN KEY (booking_seat_id) REFERENCES booking_seats(id)
    ON DELETE SET NULL,
  CONSTRAINT fk_tickets_seat
    FOREIGN KEY (seat_id) REFERENCES seats(id)
    ON DELETE SET NULL
) ENGINE=InnoDB;

-- =========================
-- NOTIFICATIONS / REFUNDS
-- =========================

CREATE TABLE notifications (
  id CHAR(36) PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  booking_id CHAR(36),
  channel VARCHAR(20) NOT NULL,
  type VARCHAR(40) NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK (channel IN ('email','web')),
  CHECK (type IN ('ticket_issued','refund_update')),
  CHECK (status IN ('pending','sent','failed')),
  CONSTRAINT fk_notifications_user
    FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_notifications_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id)
    ON DELETE SET NULL
) ENGINE=InnoDB;

CREATE TABLE refund_requests (
  id CHAR(36) PRIMARY KEY,
  booking_id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  organizer_id CHAR(36) NOT NULL,
  status VARCHAR(20) NOT NULL,
  reason TEXT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  decided_at DATETIME,
  CHECK (status IN ('requested','approved','rejected')),
  CONSTRAINT fk_refunds_booking
    FOREIGN KEY (booking_id) REFERENCES bookings(id),
  CONSTRAINT fk_refunds_user
    FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_refunds_organizer
    FOREIGN KEY (organizer_id) REFERENCES organizers(id)
) ENGINE=InnoDB;
