/* Replace with your SQL commands */

-- +migrate Up
ALTER TABLE guild_users ALTER COLUMN invited_by TYPE
TEXT;

ALTER TABLE invite_histories ADD COLUMN type TEXT DEFAULT 'normal';
