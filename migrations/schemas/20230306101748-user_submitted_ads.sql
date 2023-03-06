
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_submitted_ads(
    id SERIAL PRIMARY KEY,
    creator_id TEXT NOT NULL,
    ad_channel_id TEXT NOT NULL,
    status TEXT DEFAULT 'pending',
    introduction TEXT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS user_submitted_ads;