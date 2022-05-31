/* Replace with your SQL commands */

-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_level_roles (
	guild_id TEXT NOT NULL REFERENCES discord_guilds(id),
	level INTEGER NOT NULL REFERENCES config_xp_levels(level),
	role_id TEXT NOT NULL,
	CONSTRAINT guild_config_levels_roles_pkey PRIMARY KEY (guild_id, level)
);