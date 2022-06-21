
-- +migrate Up
ALTER TABLE guild_config_sales_tracker RENAME TO guild_config_sales_trackers;

-- +migrate Down
ALTER TABLE guild_config_sales_trackers RENAME TO guild_config_sales_tracker;