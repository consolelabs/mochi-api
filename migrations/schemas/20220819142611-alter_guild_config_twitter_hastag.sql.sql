
-- +migrate Up
ALTER TABLE guild_config_twitter_hashtags ADD COLUMN from_twitter TEXT;
-- +migrate Down
ALTER TABLE guild_config_twitter_hashtags DROP COLUMN from_twitter;