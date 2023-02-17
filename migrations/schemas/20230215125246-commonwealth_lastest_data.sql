
-- +migrate Up
CREATE TABLE IF NOT EXISTS commonwealth_latest_data(
    id SERIAL PRIMARY KEY,
    community_id TEXT UNIQUE NOT NULL,
    post_count INTEGER,
    latest_at timestamptz
);

-- +migrate Down
DROP TABLE IF EXISTS commonwealth_latest_data;