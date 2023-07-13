
-- +migrate Up
alter table coingecko_supported_tokens add column is_not_supported boolean default false;
-- +migrate Down
alter table coingecko_supported_tokens drop column is_not_supported;
