-- +migrate Up
CREATE TABLE IF NOT EXISTS chains (
	id SERIAL NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	rpc TEXT NOT NULL,
	api_base_url TEXT NOT NULL,
	api_key TEXT NOT NULL,
	tx_base_url TEXT NOT NULL,
	currency TEXT
);

INSERT INTO
	chains(
		id,
		name,
		rpc,
		api_base_url,
		api_key,
		tx_base_url,
		currency
	)
VALUES
	(
		1,
		'Ethereum Mainnet',
		'https://mainnet.infura.io/v3/5b389eb75c514cf6b1711d70084b0114',
		'https://api.etherscan.io/api?',
		'',
		'https://etherscan.io/tx',
		'ETH'
	);

INSERT INTO
	chains(
		id,
		name,
		rpc,
		api_base_url,
		api_key,
		tx_base_url,
		currency
	)
VALUES
	(
		56,
		'Binance Smart Chain Mainnet',
		'https://bsc-dataseed.binance.org',
		'https://api.bscscan.com/api?',
		'',
		'https://bscscan.com/tx',
		'BNB'
	);

INSERT INTO
	chains(
		id,
		name,
		rpc,
		api_base_url,
		api_key,
		tx_base_url,
		currency
	)
VALUES
	(
		250,
		'Fantom Opera',
		'https://rpc.ftm.tools',
		'https://api.ftmscan.com/api?',
		'',
		'https://ftmscan.com/tx',
		'FTM'
	);

ALTER TABLE
	tokens
ADD
	COLUMN is_native BOOLEAN DEFAULT FALSE,
ADD
	CONSTRAINT tokens_chain_id_fkey FOREIGN KEY(chain_id) REFERENCES chains(id);

-- +migrate Down
ALTER TABLE
	tokens DROP COLUMN IF EXISTS is_native,
	DROP CONSTRAINT tokens_chain_id_fkey;

DROP TABLE IF EXISTS chains;
