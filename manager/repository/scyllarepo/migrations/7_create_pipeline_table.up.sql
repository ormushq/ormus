CREATE TABLE IF NOT EXISTS pipelines (
    id TEXT PRIMARY KEY,
    integrations SET<TEXT>
);