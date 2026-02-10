-- migrate:up
CREATE TABLE
  payments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    idempotency_key VARCHAR(255) NULL UNIQUE,
    user_id INT NOT NULL,
    booking_id INT NOT NULL,
    transaction_id VARCHAR(255) DEFAULT NULL,
    payment_code VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    provider VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    paid_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- migrate:down
DROP TABLE IF EXISTS payments;
