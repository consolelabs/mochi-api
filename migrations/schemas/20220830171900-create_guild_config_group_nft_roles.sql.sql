
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_group_nft_roles (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	group_name TEXT,
	guild_id TEXT NOT NULL,
	role_id TEXT NOT NULL,
	number_of_tokens INTEGER NOT NULL,
	"created_at" timestamp with time zone default now(),
	"updated_at" timestamp with time zone default now(),
	UNIQUE (guild_id, role_id)
);
ALTER TABLE guild_config_nft_roles ADD group_id UUID REFERENCES guild_config_group_nft_roles(id);
ALTER TABLE guild_config_nft_roles DROP COLUMN role_id;
ALTER TABLE guild_config_nft_roles DROP COLUMN guild_id;
ALTER TABLE guild_config_nft_roles DROP COLUMN token_id;
-- +migrate Down
DROP TABLE IF EXISTS guild_config_group_nft_roles;
ALTER TABLE guild_config_nft_roles DROP group_id;
ALTER TABLE guild_config_nft_roles ADD COLUMN role_id TEXT;
ALTER TABLE guild_config_nft_roles ADD COLUMN guild_id TEXT;
ALTER TABLE guild_config_nft_roles ADD COLUMN token_id TEXT;