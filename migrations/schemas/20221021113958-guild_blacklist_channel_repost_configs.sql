
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_blacklist_channel_repost_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    guild_id TEXT,
    channel_id TEXT,
    created_at timestamptz DEFAULT now()
);
CREATE UNIQUE INDEX guild_blacklist_channel_repost_configs_uindex on guild_blacklist_channel_repost_configs (channel_id);
-- +migrate Down
DROP INDEX guild_blacklist_channel_repost_configs_uindex;
DROP TABLE IF EXISTS guild_blacklist_channel_repost_configs;