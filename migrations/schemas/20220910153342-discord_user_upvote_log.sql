
-- +migrate Up
CREATE TABLE IF NOT EXISTS discord_user_upvote_logs
(
    discord_id TEXT NOT NULL,
    source TEXT NOT NULL,
    latest_upvote_time timestamp default now(),
    created_at timestamp default now(),
    UNIQUE(discord_id, source)
);
-- +migrate Down
DROP TABLE IF EXISTS discord_user_upvote_logs;