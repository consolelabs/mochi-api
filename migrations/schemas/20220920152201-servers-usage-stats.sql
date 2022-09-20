
-- +migrate Up
CREATE TABLE IF NOT EXISTS servers_usage_stats (
	id SERIAL NOT NULL PRIMARY KEY,
	guild_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	command TEXT NOT NULL,
    args TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS servers_usage_stats;
