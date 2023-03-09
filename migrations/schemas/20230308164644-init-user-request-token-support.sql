
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_token_support_requests (
    id SERIAL PRIMARY KEY,
    user_discord_id TEXT NOT NULL,
    channel_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    token_name TEXT NOT NULL,
    token_address TEXT NOT NULL,
    token_chain_id INTEGER NOT NULL REFERENCES chains (id),
    status TEXT DEFAULT 'pending',
    updated_at timestamptz DEFAULT now(),
    created_at timestamptz DEFAULT now(),
    UNIQUE (token_address, token_chain_id)
);

-- +migrate Down
DROP TABLE IF EXISTS user_token_support_requests;