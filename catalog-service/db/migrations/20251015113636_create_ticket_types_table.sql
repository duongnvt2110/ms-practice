-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  ticket_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    position INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    image_url VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    qty INT NOT NULL DEFAULT 1,
    price INT NOT NULL DEFAULT 0,
    location TEXT DEFAULT NULL,
    sale_at TIMESTAMP DEFAULT NULL,
    sale_end TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket_types;

-- +goose StatementEnd