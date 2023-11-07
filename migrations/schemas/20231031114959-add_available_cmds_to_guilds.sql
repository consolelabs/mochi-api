
-- +migrate Up

ALTER TABLE discord_guilds ADD COLUMN available_cmds JSONB DEFAULT NULL;

-- +migrate Down

ALTER TABLE discord_guilds DROP COLUMN IF EXISTS available_cmds;
