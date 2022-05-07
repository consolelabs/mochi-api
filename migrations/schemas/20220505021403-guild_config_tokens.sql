
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_tokens (
	guild_id TEXT NOT NULL,
	token_id INTEGER NOT NULL,
	active BOOLEAN,
	CONSTRAINT guild_config_tokens_pkey PRIMARY KEY (guild_id, token_id),
	CONSTRAINT guild_config_tokens_guild_id_fkey FOREIGN KEY(guild_id) REFERENCES discord_guilds(id),
	CONSTRAINT guild_config_tokens_token_id_fkey FOREIGN KEY(token_id) REFERENCES tokens(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS guild_config_tokens_guild_token_id_uidx ON guild_config_tokens (guild_id, token_id);

ALTER TABLE tokens
	ADD COLUMN IF NOT EXISTS coin_gecko_id TEXT UNIQUE,
	ADD COLUMN IF NOT EXISTS name TEXT;


-- +migrate Down
DROP TABLE IF EXISTS guild_config_tokens;

ALTER TABLE tokens
	DROP COLUMN IF EXISTS coin_market_cap_id,
	DROP COLUMN IF EXISTS name;

