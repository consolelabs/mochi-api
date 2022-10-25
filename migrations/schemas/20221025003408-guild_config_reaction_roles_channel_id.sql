
-- +migrate Up
ALTER TABLE guild_config_reaction_roles ADD COLUMN IF NOT EXISTS channel_id text;

-- +migrate Down
ALTER TABLE guild_config_reaction_roles DROP COLUMN IF EXISTS channel_id;