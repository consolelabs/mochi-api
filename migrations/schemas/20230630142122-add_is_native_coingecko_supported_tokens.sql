
-- +migrate Up
alter table coingecko_supported_tokens add column is_native boolean default false;
-- +migrate Down
alter table coingecko_supported_tokens drop column is_native;
