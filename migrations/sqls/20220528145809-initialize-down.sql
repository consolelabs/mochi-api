/* Replace with your SQL commands */
-- +migrate Down
ALTER TABLE discord_guilds DROP COLUMN global_xp;
