
-- +migrate Up
alter table vaults add column wallet_number integer;
-- +migrate Down
alter table vaults drop column wallet_number;
