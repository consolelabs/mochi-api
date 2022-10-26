
-- +migrate Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS is_migrate_bal bool DEFAULT FALSE;

-- +migrate Down
ALTER TABLE users DROP COLUMN IF EXISTS is_migrate_bal;
