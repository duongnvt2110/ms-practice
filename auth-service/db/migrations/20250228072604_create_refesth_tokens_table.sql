-- migrate:up
CREATE TABLE refresh_tokens (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    user_id       INT  NOT NULL,
    token         TEXT NOT NULL,
    expires_at    TIMESTAMP NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS refresh_tokens;
