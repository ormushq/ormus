CREATE TABLE IF NOT EXISTS sources (
    id TEXT PRIMARY KEY,
    write_key TEXT,
    name VARCHAR,
    description VARCHAR,
    project_id TEXT,
    owner_id TEXT,
    status TEXT,
    create_at TIMESTAMP,
    update_at TIMESTAMP,
    delete_at TIMESTAMP
);