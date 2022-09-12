
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_welcome_channels
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    guild_id TEXT UNIQUE NOT NULL,
    channel_id TEXT NOT NULL,
    welcome_message TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_welcome_channels;