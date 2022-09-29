
-- +migrate Up
ALTER TABLE discord_guilds ADD COLUMN IF NOT EXISTS active BOOLEAN DEFAULT TRUE;

-- +migrate Down
ALTER TABLE discord_guilds DROP COLUMN IF EXISTS active;
