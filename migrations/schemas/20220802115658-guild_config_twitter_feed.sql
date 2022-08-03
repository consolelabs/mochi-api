
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_twitter_feeds
(
    id uuid DEFAULT uuid_generate_v4(),
    guild_id TEXT UNIQUE NOT NULL,
    twitter_consumer_key TEXT,
    twitter_consumer_secret TEXT,
    twitter_access_token TEXT,
    twitter_access_token_secret TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_twitter_feeds;