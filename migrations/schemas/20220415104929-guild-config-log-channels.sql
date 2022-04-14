
-- +migrate Up
CREATE TABLE IF NOT EXISTS "guild_config_log_channels" (
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds" ("id"),
  "channel_id" bigint NOT NULL,
  
  PRIMARY KEY ("guild_id", "channel_id")
);
-- +migrate Down
DROP TABLE IF EXISTS "guild_config_log_channels";