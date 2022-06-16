
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_sales_tracker (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  guild_id TEXT NOT NULL,
  channel_id TEXT NOT NULL
);

CREATE UNIQUE INDEX guild_config_sales_tracker_guild_id_uidx ON guild_config_sales_tracker (guild_id);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_sales_tracker;
