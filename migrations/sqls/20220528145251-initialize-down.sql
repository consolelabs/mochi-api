/* Replace with your SQL commands */
-- +migrate Down
DROP VIEW IF EXISTS guild_user_xps;
DROP TABLE IF EXISTS guild_config_activities, guild_user_activity_xps, activities, config_xp_levels;
