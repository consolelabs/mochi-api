
-- +migrate Up
alter table coingecko_supported_tokens add column detail_platforms json;
-- +migrate Down
alter table coingecko_supported_tokens drop column detail_platforms;
