
-- +migrate Up
CREATE TABLE IF NOT EXISTS nft_add_request_histories (
	id SERIAL NOT NULL PRIMARY KEY,
	address TEXT NOT NULL,
	chain_id INTEGER NOT NULL,
	guild_id TEXT NOT NULL,
	channel_id TEXT NOT NULL,
	message_id TEXT NOT NULL UNIQUE,
	created_at timestamp DEFAULT now(),
	UNIQUE (address, chain_id)
);

-- +migrate Down
DROP TABLE IF EXISTS nft_add_request_histories;
