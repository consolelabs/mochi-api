
-- +migrate Up
ALTER TABLE servers_usage_stats ADD COLUMN created_at timestamptz DEFAULT now();

-- +migrate Down
ALTER TABLE servers_usage_stats DROP COLUMN created_at;