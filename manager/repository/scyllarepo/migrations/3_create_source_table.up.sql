CREATE TABLE IF NOT EXISTS sources (
    id TEXT PRIMARY KEY,
    write_key TEXT,
    name VARCHAR,
    description VARCHAR,
    project_id TEXT,
    owner_id TEXT,
    status TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);