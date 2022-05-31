/* Replace with your SQL commands */

-- +migrate Up
ALTER TABLE guild_config_invite_trackers ADD CONSTRAINT unique_guild_id UNIQUE (guild_id);
