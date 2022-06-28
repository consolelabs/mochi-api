
-- +migrate Up
alter table nft_collections drop constraint nft_collections_pkey;

alter table nft_collections alter column chain_id type text;

alter table nft_collections add column is_verified bool default false;

create unique index address_chain_id_unique_idx on nft_collections (lower(address), chain_id);
-- +migrate Down
alter table nft_collections add constraint  nft_collections_pkey  primary key (address);

alter table nft_collections alter column chain_id type integer;

alter table nft_collections drop column is_verified;

drop index address_chain_id_unique_idx;
