
-- +migrate Up
ALTER TABLE discord_guilds ADD COLUMN IF NOT EXISTS joined_at timestamptz DEFAULT now();
ALTER TABLE discord_guilds ADD COLUMN IF NOT EXISTS left_at timestamptz DEFAULT NULL;
-- +migrate Down
ALTER TABLE discord_guilds DROP COLUMN IF EXISTS joined_at;
ALTER TABLE discord_guilds DROP COLUMN IF EXISTS left_at;
