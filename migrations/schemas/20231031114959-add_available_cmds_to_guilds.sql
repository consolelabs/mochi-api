
-- +migrate Up

ALTER TABLE discord_guilds ADD COLUMN available_cmds TEXT NOT NULL DEFAULT '[]';

-- +migrate Down

ALTER TABLE discord_guilds DROP COLUMN IF EXISTS available_cmds;
