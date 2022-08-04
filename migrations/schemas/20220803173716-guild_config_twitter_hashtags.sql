
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_twitter_hashtags
(
    guild_id TEXT NOT NULL,
    channel_id TEXT NOT NULL,
    hashtag TEXT NOT NULL,
    UNIQUE (guild_id)
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_twitter_hashtags;