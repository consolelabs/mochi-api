/* Replace with your SQL commands */

-- +migrate Up
ALTER TABLE tokens ADD COLUMN IF NOT EXISTS guild_default bool DEFAULT FALSE;
