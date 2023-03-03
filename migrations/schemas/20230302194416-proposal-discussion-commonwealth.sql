
-- +migrate Up
ALTER TABLE commonwealth_latest_data
    ADD COLUMN name TEXT,
    ADD COLUMN description TEXT,
    ADD COLUMN website TEXT,
    ADD COLUMN icon_url TEXT;

CREATE TABLE IF NOT EXISTS commonwealth_discussion_subscriptions (
    id SERIAL PRIMARY KEY,
    discord_thread_id TEXT UNIQUE,
    discussion_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS commonwealth_discussion_subs_discussion_index on commonwealth_discussion_subscriptions(discussion_id);

-- +migrate Down
ALTER TABLE commonwealth_latest_data 
    DROP COLUMN name, 
    DROP COLUMN description, 
    DROP COLUMN website, 
    DROP COLUMN icon_url;

DROP TABLE IF EXISTS commonwealth_discussion_subscriptions;