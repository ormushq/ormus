CREATE TABLE IF NOT EXISTS sourceMetadata (
    source_id TEXT PRIMARY KEY,
    metadata_id TEXT,
    metadata_name VARCHAR,
    metadata_slug VARCHAR,
    metadata_category TEXT
);
