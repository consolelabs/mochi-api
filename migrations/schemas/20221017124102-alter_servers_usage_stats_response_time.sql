
-- +migrate Up
ALTER TABLE servers_usage_stats ADD COLUMN IF NOT EXISTS execution_time_ms integer;

-- +migrate Down
ALTER TABLE servers_usage_stats DROP COLUMN IF EXISTS execution_time_ms;