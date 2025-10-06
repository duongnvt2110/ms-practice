-- +goose Up
-- +goose StatementBegin
CREATE TABLE bookings (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    user_id       INT  NOT NULL,
    event_id      INT NOT NULL,
    status        VARCHAR(255) NOT NULL
    quantity      INT NOT NULL
    prices        FLOAT NOT NULL 
    idempotency_key VARCHAR(255) NOT NULL
    logs TEXT NOT NULL
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings;
-- +goose StatementEnd
