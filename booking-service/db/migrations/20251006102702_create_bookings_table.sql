-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  bookings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idempotency_key VARCHAR(255) NOT NULL UNIQUE,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    booking_code VARCHAR(255) NOT NULL UNIQUE,
    holded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMP NULL,
    status VARCHAR(255) NOT NULL,
    number_seats INT NOT NULL DEFAULT 0,
    total_price INT NOT NULL,
    logs JSON NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings;

-- +goose StatementEnd
