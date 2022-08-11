
-- +migrate Up
ALTER TABLE mochi_nft_sales 
ADD COLUMN hodl TEXT,
ADD COLUMN pnl TEXT,
ADD COLUMN sub_pnl TEXT;

-- +migrate Down
ALTER TABLE mochi_nft_sales 
DROP COLUMN hodl,
DROP COLUMN pnl,
DROP COLUMN sub_pnl;