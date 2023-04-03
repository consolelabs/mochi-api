
-- +migrate Up
create table if not exists kyberswap_supported_tokens (
    id serial primary key,
    address text,
    chain_id integer,
    chain_name text,
    decimals integer,
    symbol text,
    name text,
    logo_uri text,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);
-- +migrate Down
drop table if exists kyberswap_supported_tokens;