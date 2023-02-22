
-- +migrate Up
ALTER TABLE user_token_price_alerts ADD COLUMN IF NOT EXISTS price_by_percent FLOAT8 DEFAULT 0;
ALTER TABLE user_token_price_alerts RENAME COLUMN price TO value;

-- +migrate Down
ALTER TABLE user_token_price_alerts DROP COLUMN IF EXISTS price_by_percent;
ALTER TABLE user_token_price_alerts RENAME COLUMN value TO price;