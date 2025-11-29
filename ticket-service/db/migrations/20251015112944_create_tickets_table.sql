-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  tickets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    booking_id INT NOT NULL,
    payment_id INT NOT NULL,
    ticket_type_id INT NOT NULL,
    code VARCHAR(255) NOT NULL,
    qr_url TEXT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;

-- +goose StatementEnd
