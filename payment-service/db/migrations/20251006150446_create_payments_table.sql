-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idempotency_key VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    booking_id INT NOT NULL,
    transaction_id VARCHAR(255) DEFAULT NULL,
    payment_code INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    prices INT NOT NULL,
    paid_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings;

-- +goose StatementEnd