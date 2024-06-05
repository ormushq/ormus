CREATE TABLE IF NOT EXISTS integrations (
    id TEXT PRIMARY KEY,
    source_id TEXT,
    name VARCHAR,
    connection_type VARCHAR,
    enabled BOOLEAN,
    config MAP<TEXT, TEXT>,
    created_at TIMESTAMP
);