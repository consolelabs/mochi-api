
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_levelup_messages
(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    guild_id TEXT NOT NULL UNIQUE,
    message TEXT NOT NULL,
    image_url TEXT NULL,
    channel_id TEXT NULL,
    updated_at timestamptz DEFAULT now(),
    created_at timestamptz DEFAULT now()
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_levelup_messages;