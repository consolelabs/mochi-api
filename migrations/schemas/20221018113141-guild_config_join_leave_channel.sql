
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_join_leave_channels (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	guild_id TEXT NOT NULL,
	channel_id TEXT NOT NULL,
    unique(guild_id)
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_join_leave_channels;
