
-- +migrate Up
CREATE TABLE IF NOT EXISTS offchain_tip_bot_deposit_logs (
	chain_id UUID NOT NULL,
	tx_hash TEXT NOT NULL,
	token_id UUID NOT NULL,
	from_address TEXT NOT NULL,
	to_address TEXT NOT NULL,
	amount float8,
	amount_in_usd float8,
	user_id TEXT NOT NULL,
	block_number INTEGER,
	signed_at timestamptz,
	FOREIGN KEY (chain_id) REFERENCES offchain_tip_bot_chains(id),
	FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id),
	FOREIGN KEY (user_id) REFERENCES users(id),
	PRIMARY KEY (chain_id, tx_hash)
);

ALTER TABLE offchain_tip_bot_chains
	ADD COLUMN is_evm BOOLEAN NOT NULL DEFAULT FALSE,
	ADD COLUMN support_deposit BOOLEAN NOT NULL DEFAULT FALSE,
	ALTER COLUMN chain_id DROP NOT NULL,
	ALTER COLUMN chain_id TYPE INTEGER USING NULL;

-- +migrate Down
DROP TABLE IF EXISTS offchain_tip_bot_deposit_logs;

ALTER TABLE offchain_tip_bot_chains
	DROP COLUMN is_evm,
	DROP COLUMN support_deposit,
	ALTER COLUMN chain_id TYPE TEXT USING NULL;
