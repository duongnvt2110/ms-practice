-- +goose Up
-- +goose StatementBegin
CREATE TABLE bookings (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    booking_id INT NOT NULL,
    event_type_id INT NOT NULL, 
    qty  INT NOT NULL, 
    unit_price FLOAT NOT NULL, 
    currency  VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS booking_items;
-- +goose StatementEnd