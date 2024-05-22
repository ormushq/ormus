CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    name VARCHAR,
    type VARCHAR,
    event VARCHAR,
    received_at TIMESTAMP,
    send_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);