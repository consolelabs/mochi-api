
-- +migrate Up
ALTER TABLE guild_config_twitter_hashtags ADD COLUMN updated_at timestamptz DEFAULT now();

-- +migrate Down
ALTER TABLE guild_config_twitter_hashtags DROP COLUMN updated_at;