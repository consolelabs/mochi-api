
-- +migrate Up
alter table kyberswap_supported_tokens add constraint unique_address_symbol_chain unique (symbol, address, chain_name);
-- +migrate Down
alter table kyberswap_supported_tokens drop constraint unique_address_symbol_chain;
