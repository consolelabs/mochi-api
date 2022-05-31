/* Replace with your SQL commands */
-- +migrate Down
ALTER TABLE tokens DROP COLUMN IF EXISTS guild_default;
