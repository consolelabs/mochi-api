
-- +migrate Up
alter table coingecko_supported_tokens add column coingecko_info jsonb;

-- +migrate Down
alter table coingecko_supported_tokens drop column coingecko_info;

