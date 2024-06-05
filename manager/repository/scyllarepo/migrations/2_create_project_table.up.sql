CREATE TABLE IF NOT EXISTS projects (
    id VARCHAR PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR,
    description VARCHAR,
    user_id VARCHAR
);
