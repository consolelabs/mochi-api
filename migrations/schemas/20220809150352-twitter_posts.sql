
-- +migrate Up
CREATE TABLE IF NOT EXISTS twitter_posts
(
    id uuid DEFAULT uuid_generate_v4(),
    twitter_id TEXT,
    twitter_handle TEXT,
    tweet_id TEXT,
    guild_id TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS twitter_posts;