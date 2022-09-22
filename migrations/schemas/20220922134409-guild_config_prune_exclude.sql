
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_whitelist_prunes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	guild_id TEXT NOT NULL,
	role_id TEXT NOT NULL,
    unique(guild_id,role_id)
);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_whitelist_prunes;
