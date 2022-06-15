-- +migrate Up
ALTER TABLE
	discord_guilds
ADD
	COLUMN log_channel TEXT;

-- +migrate Down
ALTER TABLE
	discord_guilds DROP COLUMN log_channel;