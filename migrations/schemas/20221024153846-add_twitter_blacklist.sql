
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_twitter_blacklist (
	id SERIAL NOT NULL PRIMARY KEY,
	guild_id TEXT NOT NULL,
	twitter_username TEXT NOT NULL,
	twitter_id TEXT NOT NULL,
	created_by TEXT NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	FOREIGN KEY (guild_id) REFERENCES discord_guilds(id),
	UNIQUE (guild_id, twitter_username),
	UNIQUE (guild_id, twitter_id)
);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_twitter_blacklist;
