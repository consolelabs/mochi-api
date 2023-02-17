
-- +migrate Up
ALTER TABLE user_wallet_watchlist_items ADD COLUMN IF NOT EXISTS type TEXT;

-- +migrate Down
ALTER TABLE user_wallet_watchlist_items DROP COLUMN IF EXISTS type;
