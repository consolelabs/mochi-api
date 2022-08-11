
-- +migrate Up
ALTER TABLE chains ADD COLUMN coin_gecko_id text unique;
UPDATE chains SET coin_gecko_id = CASE
    WHEN id = 1 then 'ethereum'
    WHEN id = 56 then 'binance-smart-chain'
    WHEN id = 250 then 'fantom'
end;

-- +migrate Down
ALTER TABLE chains DROP COLUMN coin_gecko_id;