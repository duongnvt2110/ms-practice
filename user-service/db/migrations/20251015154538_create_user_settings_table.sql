-- migrate:up
CREATE TABLE
  user_settings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    allow_noti BOOLEAN,
  );

-- migrate:down
DROP TABLE IF EXISTS user_settings;