
-- +migrate Up
ALTER TABLE moniker_configs DROP CONSTRAINT moniker_configs_guild_id_fkey;

-- +migrate Down
ALTER TABLE moniker_configs ADD CONSTRAINT guild_id FOREIGN KEY (guild_id) REFERENCES discord_guilds(id);