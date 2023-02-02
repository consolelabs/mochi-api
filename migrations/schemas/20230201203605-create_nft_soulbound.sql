
-- +migrate Up
create table if not exists nft_soulbounds (
    id serial primary key,
    collection_address text,
    trait_type text,
    value text,
    total_soulbound integer,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
-- +migrate Down
drop table if exists nft_soulbounds;