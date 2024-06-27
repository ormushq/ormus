CREATE TABLE IF NOT EXISTS destination_metadata (
    id TEXT PRIMARY KEY,
    name VARCHAR,
    slug VARCHAR,
    categories SET<VARCHAR>
);
