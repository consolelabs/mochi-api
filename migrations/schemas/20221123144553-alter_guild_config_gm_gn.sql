
-- +migrate Up
ALTER TABLE guild_config_gm_gn ADD COLUMN IF NOT EXISTS msg text;
ALTER TABLE guild_config_gm_gn ADD COLUMN IF NOT EXISTS emoji text;
ALTER TABLE guild_config_gm_gn ADD COLUMN IF NOT EXISTS sticker text;
-- +migrate Down
ALTER TABLE guild_config_gm_gn DROP COLUMN IF EXISTS msg;
ALTER TABLE guild_config_gm_gn DROP COLUMN IF EXISTS emoji;
ALTER TABLE guild_config_gm_gn DROP COLUMN IF EXISTS sticker;