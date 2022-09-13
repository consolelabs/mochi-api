
-- +migrate Up
CREATE TABLE IF NOT EXISTS coingecko_supported_tokens (
	id TEXT NOT NULL PRIMARY KEY,
	symbol TEXT NOT NULL,
	name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS coingecko_supported_tokens_symbol_idx ON coingecko_supported_tokens(symbol);

-- +migrate Down
DROP TABLE IF EXISTS coingecko_supported_tokens;