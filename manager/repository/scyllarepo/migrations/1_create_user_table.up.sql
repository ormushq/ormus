CREATE TABLE IF NOT EXISTS users (
    id          VARCHAR ,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    email       VARCHAR,
    password    VARCHAR,
    is_active BOOLEAN,
    PRIMARY KEY(email,is_active)
);