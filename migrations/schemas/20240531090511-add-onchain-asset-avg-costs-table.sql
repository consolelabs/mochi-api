-- +migrate Up
CREATE TABLE IF NOT EXISTS onchain_asset_avg_costs (
    wallet_address TEXT NOT NULL,
    token_address TEXT NOT NULL,
    symbol TEXT NOT NULL,
    blockchain TEXT NOT NULL,
    average_cost FLOAT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX "uidx_wallet_address_token_address_chain" ON onchain_asset_avg_costs (wallet_address, token_address, blockchain);

-- +migrate Down
DELETE TABLE IF EXISTS onchain_asset_avg_costs;