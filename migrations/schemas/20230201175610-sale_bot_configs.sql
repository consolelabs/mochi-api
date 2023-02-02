
-- +migrate Up
CREATE TABLE IF NOT EXISTS sale_bot_marketplaces (
	id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS sale_bot_twitter_configs (
	id SERIAL PRIMARY KEY,
	address TEXT UNIQUE NOT NULL,
	slug TEXT NOT NULL DEFAULT '',
	collection_name TEXT NOT NULL,
	chain_id INT NOT NULL,
	marketplace_id INT NOT NULL REFERENCES sale_bot_marketplaces(id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS sale_bot_twitter_configs;
DROP TABLE IF EXISTS sale_bot_marketplaces;
