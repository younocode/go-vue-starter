CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    id            serial PRIMARY KEY,
    email         CITEXT UNIQUE NOT NULL, -- 忽略大小写
    password      VARCHAR(255)  NOT NULL, --  hashed password
    refresh_token VARCHAR(255),           -- jwt flush token
    created_at    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP
);