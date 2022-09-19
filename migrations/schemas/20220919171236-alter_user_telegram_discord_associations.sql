
-- +migrate Up
DROP TABLE IF EXISTS user_telegram_discord_associations;

CREATE TABLE IF NOT EXISTS user_telegram_discord_associations (
	telegram_username TEXT NOT NULL PRIMARY KEY,
	discord_id TEXT NOT NULL
);

CREATE UNIQUE INDEX user_telegram_discord_associations_discord_id_uidx ON user_telegram_discord_associations(discord_id);

-- +migrate Down
DROP TABLE IF EXISTS user_telegram_discord_associations;

CREATE TABLE IF NOT EXISTS user_telegram_discord_associations (
	telegram_id INTEGER NOT NULL PRIMARY KEY,
	discord_id TEXT NOT NULL
);

CREATE UNIQUE INDEX user_telegram_discord_associations_discord_id_uidx ON user_telegram_discord_associations(discord_id);