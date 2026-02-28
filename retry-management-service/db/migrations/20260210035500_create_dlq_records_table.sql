-- migrate:up
CREATE TABLE
  dlq_records (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(255) NOT NULL,
    partition_id INT NOT NULL,
    offset_id BIGINT NOT NULL,
    `key` BLOB DEFAULT NULL,
    headers LONGTEXT DEFAULT NULL,
    payload LONGBLOB NOT NULL,
    payload_json LONGTEXT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

-- migrate:down
DROP TABLE IF EXISTS dlq_records;
