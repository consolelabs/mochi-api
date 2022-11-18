
-- +migrate Up
ALTER TABLE user_watchlist_items ADD COLUMN IF NOT EXISTS is_fiat BOOLEAN DEFAULT false;

-- +migrate Down
ALTER TABLE user_watchlist_items DROP COLUMN IF EXISTS is_fiat;
