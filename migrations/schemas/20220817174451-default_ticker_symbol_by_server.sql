
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_default_ticker(
	id SERIAL NOT NULL PRIMARY KEY,
	guild_id TEXT NOT NULL REFERENCES discord_guilds(id),
	query TEXT NOT NULL,
	default_ticker TEXT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_guild_config_default_ticker ON guild_config_default_ticker(guild_id, query);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_default_ticker;