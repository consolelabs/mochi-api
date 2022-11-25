
-- +migrate Up
ALTER TABLE discord_user_token_alerts ADD COLUMN IF NOT EXISTS symbol TEXT;

-- +migrate Down
ALTER TABLE discord_user_token_alerts DROP COLUMN IF EXISTS symbol;
