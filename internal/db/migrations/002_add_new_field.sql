-- internal/db/migrations/002_add_new_field.sql
CREATE TABLE IF NOT EXISTS schema_migrations (
    version INTEGER PRIMARY KEY,
    applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Check if priority column exists before adding it
INSERT OR IGNORE INTO schema_migrations (version) VALUES (2);
