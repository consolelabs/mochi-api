
-- +migrate Up
drop index address_chain_id_unique_idx;
create unique index address_chain_id_unique_idx on nft_collections (address, chain_id);
-- +migrate Down
create unique index address_chain_id_unique_idx on nft_collections (lower(address), chain_id);
drop index address_chain_id_unique_idx;