
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_watchlist_items (
	user_id TEXT NOT NULL,
	coin_gecko_id TEXT NOT NULL,
	symbol TEXT NOT NULL,
	UNIQUE(user_id, coin_gecko_id)
);

CREATE INDEX IF NOT EXISTS user_watchlist_items_user_id_idx ON user_watchlist_items(user_id);

-- +migrate Down
DROP TABLE IF EXISTS user_watchlist_items;