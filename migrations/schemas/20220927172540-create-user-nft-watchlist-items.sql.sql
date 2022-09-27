
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_nft_watchlist_items (
	user_id TEXT NOT NULL,
	symbol TEXT NOT NULL,
    collection_address TEXT NOT NULL,
    chain_id integer NOT NULL,
	created_at timestamptz DEFAULT now(),
	UNIQUE(user_id, symbol, collection_address, chain_id)
);

-- +migrate Down
DROP TABLE IF EXISTS user_nft_watchlist_items;