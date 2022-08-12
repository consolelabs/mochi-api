
-- +migrate Up
ALTER TABLE guild_config_tokens ADD COLUMN is_default BOOLEAN DEFAULT FALSE;

-- +migrate Down
ALTER TABLE guild_config_tokens DROP COLUMN is_default;