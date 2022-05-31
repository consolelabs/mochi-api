/* Replace with your SQL commands */

-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_gm_gn (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  guild_id TEXT NOT NULL,
  channel_id TEXT NOT NULL
);

CREATE UNIQUE INDEX guild_config_gm_gn_guild_id_uidx ON guild_config_gm_gn (guild_id);
