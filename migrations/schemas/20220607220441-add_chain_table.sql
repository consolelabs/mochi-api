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
		'SM5BHYSNIRZ1HEWJ1JPHVTMJS95HRA6DQF',
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
		'VTKF4RG4HP6WXQ5QTAJ8MHDDIUFYD6VZHC',
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
		'XEKSVDF5VWQDY5VY6ZNT6AK9QPQRH483EF',
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
DROP TABLE IF EXISTS chains;

ALTER TABLE
	tokens DROP COLUMN IF EXISTS is_native,
	DROP CONSTRAINT tokens_chain_id_fkey;