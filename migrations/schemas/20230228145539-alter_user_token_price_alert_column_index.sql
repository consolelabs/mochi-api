
-- +migrate Up
ALTER TABLE user_token_price_alerts ADD COLUMN IF NOT EXISTS id SERIAL;

-- +migrate Down
ALTER TABLE user_token_price_alerts DROP COLUMN IF EXISTS id;