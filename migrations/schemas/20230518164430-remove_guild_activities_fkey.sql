
-- +migrate Up
ALTER TABLE guild_config_activities DROP CONSTRAINT guild_config_activities_guild_id_fkey;
ALTER TABLE guild_user_activity_logs DROP CONSTRAINT guild_user_activity_xps_guild_id_fkey;

-- +migrate Down
ALTER TABLE guild_config_activities ADD CONSTRAINT guild_config_activities_guild_id_fkey FOREIGN KEY(guild_id) REFERENCES discord_guilds(id);
ALTER TABLE guild_user_activity_logs ADD CONSTRAINT guild_user_activity_xps_guild_id_fkey FOREIGN KEY(guild_id) REFERENCES discord_guilds(id);
