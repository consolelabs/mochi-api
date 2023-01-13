
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_token_balances (
	user_address TEXT NOT NULL,
	chain_type TEXT NOT NULL,
	token_id INTEGER NOT NULL REFERENCES tokens(id),
	balance FLOAT8 NOT NULL DEFAULT 0,
	PRIMARY KEY (user_address, chain_type, token_id)
);

-- +migrate Down
DROP TABLE IF EXISTS user_token_balances;