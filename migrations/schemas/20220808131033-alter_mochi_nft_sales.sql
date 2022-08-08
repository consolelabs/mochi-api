
-- +migrate Up
ALTER TABLE mochi_nft_sales 
ADD COLUMN collection_address TEXT,
ADD COLUMN token_id TEXT,
ADD COLUMN tx_url TEXT;

-- +migrate Down
ALTER TABLE mochi_nft_sales 
DROP COLUMN collection_address,
DROP COLUMN token_id,
DROP COLUMN tx_url;