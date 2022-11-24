
-- +migrate Up
ALTER TABLE user_nft_balances ADD COLUMN staking_nekos integer NOT NULL DEFAULT 0;
-- +migrate Down
ALTER TABLE user_nft_balances DROP COLUMN staking_nekos;