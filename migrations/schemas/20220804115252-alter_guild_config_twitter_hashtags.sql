
-- +migrate Up
ALTER TABLE guild_config_twitter_hashtags 
ADD COLUMN created_at timestamptz DEFAULT now(),
ADD COLUMN user_id TEXT;

-- +migrate Down
ALTER TABLE guild_config_twitter_hashtags 
DROP COLUMN created_at,
DROP COLUMN user_id;