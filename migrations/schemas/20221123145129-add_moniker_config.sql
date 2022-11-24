
-- +migrate Up
CREATE TABLE IF NOT EXISTS moniker_configs (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    moniker TEXT NOT NULL,
    plural TEXT,
    guild_id TEXT NOT NULL REFERENCES discord_guilds (id),
    token_id UUID NOT NULL REFERENCES offchain_tip_bot_tokens (id),
    amount float8 NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    UNIQUE (guild_id, moniker)
);
-- +migrate Down
DROP TABLE IF EXISTS moniker_configs;
