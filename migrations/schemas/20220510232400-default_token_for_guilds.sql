
-- +migrate Up
ALTER TABLE tokens ADD COLUMN IF NOT EXISTS guild_default bool DEFAULT FALSE;

-- +migrate Down
ALTER TABLE tokens DROP COLUMN IF EXISTS guild_default;
