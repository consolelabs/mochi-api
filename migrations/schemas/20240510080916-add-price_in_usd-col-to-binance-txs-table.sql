-- +migrate Up
ALTER TABLE
    IF EXISTS binance_spot_transactions
ADD
    COLUMN price_in_usd TEXT NOT NULL default '0.0';

UPDATE
    binance_spot_transactions
SET
    price_in_usd = price
Where
    price_in_usd = '0.0';

-- +migrate Down
ALTER TABLE
    IF EXISTS binance_spot_transactions DROP COLUMN price_in_usd;