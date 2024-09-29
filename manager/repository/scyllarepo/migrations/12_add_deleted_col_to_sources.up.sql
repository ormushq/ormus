ALTER TABLE sources ADD deleted boolean;
CREATE INDEX IF NOT EXISTS deleted_idx ON sources (deleted);
