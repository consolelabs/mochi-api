
-- +migrate Up
ALTER TABLE guild_config_twitter_hashtags 
ADD COLUMN twitter_username TEXT,
ADD COLUMN rule_id TEXT;

-- +migrate Down
ALTER TABLE guild_config_twitter_hashtags 
DROP COLUMN twitter_username,
DROP COLUMN rule_id;