
-- +migrate Up
alter table vaults add column solana_wallet_address text;
-- +migrate Down
alter table vaults drop column solana_wallet_address;
