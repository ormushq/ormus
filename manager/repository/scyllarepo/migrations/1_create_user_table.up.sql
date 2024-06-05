CREATE TABLE IF NOT EXISTS users (
    id          VARCHAR PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    email       VARCHAR,
    password    VARCHAR,
    is_active BOOLEAN,
);