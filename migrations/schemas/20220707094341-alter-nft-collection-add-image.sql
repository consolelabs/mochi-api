
-- +migrate Up
ALTER TABLE nft_collections ADD COLUMN image TEXT;
-- +migrate Down
ALTER TABLE nft_collections DROP COLUMN image;
