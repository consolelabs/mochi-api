
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_default_currencies
(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    guild_id TEXT NOT NULL,
    tip_bot_token_id uuid NOT NULL,
    updated_at timestamptz DEFAULT now(),
    created_at timestamptz DEFAULT now(),
    UNIQUE(guild_id),
    FOREIGN KEY (tip_bot_token_id) REFERENCES offchain_tip_bot_tokens(id)
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_default_currencies;