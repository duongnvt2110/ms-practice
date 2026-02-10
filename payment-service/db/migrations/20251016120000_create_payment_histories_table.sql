-- migrate:up
CREATE TABLE
  payment_histories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    payment_id INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    logs JSON NULL,
    paid_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- migrate:down
DROP TABLE IF EXISTS payment_histories;
