
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_wallet_watchlist_items (
	user_id TEXT NOT NULL REFERENCES users(id),
	address TEXT NOT NULL,
	alias TEXT,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	UNIQUE (user_id, address),
	UNIQUE (user_id, alias)
);

-- +migrate Down
DROP TABLE IF EXISTS user_wallet_watchlist_items;
