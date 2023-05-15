
-- +migrate Up
alter table vault_transactions add column sender text;
-- +migrate Down
alter table vault_transactions drop column sender;