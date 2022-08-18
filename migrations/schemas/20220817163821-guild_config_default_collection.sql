
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_default_collections
(
    guild_id TEXT NOT NULL,
    symbol TEXT,
    address TEXT NOT NULL,
    chain_id TEXT,
    updated_at timestamptz DEFAULT now(),
    created_at timestamptz DEFAULT now(),
    UNIQUE(guild_id, address)
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_default_collections;