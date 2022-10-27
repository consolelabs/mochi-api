
-- +migrate Up
CREATE TABLE IF NOT EXISTS migrate_balances (
	id SERIAL NOT NULL PRIMARY KEY,
	symbol text,
	created_at timestamptz NOT NULL DEFAULT NOW(),
    username text,
    user_discord_id text,
    txHash text,
    txURL text,
    transferredAmount float8
);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_twitter_blacklist;
