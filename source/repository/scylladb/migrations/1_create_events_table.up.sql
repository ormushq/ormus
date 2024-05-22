CREATE TABLRE IF NOT EXISTS `events` {
    `id`    UUID PRIMARY KEY
    `key`   TEXT,
    `value` TEXT,
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP
}