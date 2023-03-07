
-- +migrate Up
CREATE TABLE IF NOT EXISTS coingecko_token_aliases
(
    alias TEXT NOT NULL
);

CREATE UNIQUE INDEX coingecko_token_aliases_alias_uindex
    ON coingecko_token_aliases (alias);

-- +migrate Down
DROP TABLE IF EXISTS coingecko_token_aliases;