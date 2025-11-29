-- migrate:up
CREATE TABLE
  auth_refresh_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    auth_profile_id INT NOT NULL,
    refresh_token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

-- migrate:down
DROP TABLE IF EXISTS auth_refresh_tokens;
