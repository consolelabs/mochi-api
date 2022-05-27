
-- +migrate Up
ALTER TABLE discord_guilds ADD COLUMN global_xp BOOLEAN DEFAULT FALSE;

-- +migrate Down
ALTER TABLE discord_guilds DROP COLUMN global_xp;
