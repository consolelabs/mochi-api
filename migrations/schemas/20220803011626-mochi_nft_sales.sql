
-- +migrate Up
CREATE TABLE IF NOT EXISTS mochi_nft_sales
(
    id uuid DEFAULT uuid_generate_v4(),
    is_notified_twitter BOOLEAN DEFAULT FALSE,
    token_name TEXT ,
    collection_name TEXT,
    price TEXT,
    seller_address TEXT,
    buyer_address TEXT,
    marketplace TEXT,
    marketplace_url TEXT,
    image TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS mochi_nft_sales;