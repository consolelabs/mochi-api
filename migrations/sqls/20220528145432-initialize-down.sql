/* Replace with your SQL commands */
-- +migrate Down
alter table nft_collections add constraint  nft_collections_pkey  primary key (address);

alter table nft_collections alter column chain_id type integer;

alter table nft_collections drop column is_verified;

drop index address_chain_id_unique_idx;
