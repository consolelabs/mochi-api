
-- +migrate Up
drop index guild_config_sales_tracker_guild_id_uidx;
-- +migrate Down
CREATE UNIQUE INDEX guild_config_sales_tracker_guild_id_uidx ON guild_config_sales_trackers (guild_id);