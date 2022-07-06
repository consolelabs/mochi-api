
-- +migrate Up
ALTER TABLE nft_collections ADD COLUMN author TEXT;
-- +migrate Down
ALTER TABLE nft_collections DROP COLUMN author;
