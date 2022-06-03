
-- +migrate Up
ALTER TABLE nft_collections
ADD COLUMN id UUID PRIMARY KEY DEFAULT uuid_generate_v4();

CREATE TABLE IF NOT EXISTS user_nft_balances (
	user_address TEXT NOT NULL,
	chain_type TEXT NOT NULL,
	nft_collection_id UUID NOT NULL REFERENCES nft_collections(id),
	token_id TEXT NOT NULL DEFAULT '',
	balance INTEGER NOT NULL DEFAULT 0,
	PRIMARY KEY (user_address, chain_type ,nft_collection_id, token_id)
);

CREATE TABLE IF NOT EXISTS guild_config_nft_roles (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	nft_collection_id UUID NOT NULL REFERENCES nft_collections(id),
	guild_id TEXT NOT NULL,
	role_id TEXT NOT NULL,
	number_of_tokens INTEGER NOT NULL,
	token_id TEXT,
	UNIQUE (guild_id, role_id)
);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_nft_roles;

DROP TABLE IF EXISTS user_nft_balances;

ALTER TABLE nft_collections DROP COLUMN id;