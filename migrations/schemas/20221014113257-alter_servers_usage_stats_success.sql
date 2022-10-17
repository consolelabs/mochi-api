
-- +migrate Up
ALTER TABLE servers_usage_stats ADD COLUMN IF NOT EXISTS success BOOLEAN DEFAULT true;

-- +migrate Down
ALTER TABLE servers_usage_stats DROP COLUMN success IF EXISTS;