
-- +migrate Up
ALTER TABLE nft_collections ADD COLUMN created_at timestamptz DEFAULT now();
-- +migrate Down
ALTER TABLE nft_collections DROP COLUMN created_at;