CREATE TABLE IF NOT EXISTS connections (
    id TEXT PRIMARY KEY,
    pipes LIST<TEXT>,
    integrations SET<TEXT>
);