
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_dao_trackers (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	space TEXT NOT NULL,
	guild_id TEXT NOT NULL,
	channel_id TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(guild_id, space)
);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_dao_trackers;
