/* Replace with your SQL commands */
-- +migrate Up
ALTER TABLE discord_guilds ADD COLUMN global_xp BOOLEAN DEFAULT FALSE;
