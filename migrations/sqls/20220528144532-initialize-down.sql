/* Replace with your SQL commands */
-- +migrate Down
ALTER TABLE guild_config_invite_trackers DROP CONSTRAINT unique_guild_id;
