
-- +migrate Up
ALTER TABLE coingecko_supported_tokens ALTER COLUMN detail_platforms set default '[]';
UPDATE coingecko_supported_tokens SET detail_platforms = '[]' WHERE detail_platforms is null;
ALTER TABLE coingecko_supported_tokens ALTER COLUMN detail_platforms set not null;

-- +migrate Down
ALTER TABLE coingecko_supported_tokens ALTER COLUMN detail_platforms drop default;
ALTER TABLE coingecko_supported_tokens ALTER COLUMN detail_platforms drop not null;
