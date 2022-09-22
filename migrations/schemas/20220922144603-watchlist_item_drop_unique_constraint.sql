
-- +migrate Up
ALTER TABLE user_watchlist_items ALTER COLUMN coin_gecko_id DROP NOT NULL;

ALTER TABLE user_watchlist_items DROP CONSTRAINT user_watchlist_items_user_id_coin_gecko_id_key;

-- +migrate Down
ALTER TABLE user_watchlist_items ALTER COLUMN coin_gecko_id SET NOT NULL;

ALTER TABLE user_watchlist_items ADD CONSTRAINT user_watchlist_items_user_id_coin_gecko_id_key UNIQUE(user_id, coin_gecko_id);
