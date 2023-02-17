
-- +migrate Up
ALTER TABLE guild_config_dao_trackers ADD COLUMN source TEXT;
-- +migrate Down
ALTER TABLE guild_config_dao_trackers DROP COLUMN source;
