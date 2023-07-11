
-- +migrate Up
ALTER TABLE user_watchlist_items RENAME TO user_token_watchlist_items;

ALTER TABLE user_token_watchlist_items ADD COLUMN profile_id text not null default '';
ALTER TABLE user_nft_watchlist_items ADD COLUMN profile_id text not null default '';
ALTER TABLE user_wallet_watchlist_items ADD COLUMN profile_id text not null default '';

UPDATE user_token_watchlist_items SET profile_id = user_id WHERE user_id is not null;
UPDATE user_nft_watchlist_items SET profile_id = user_id WHERE user_id is not null;
UPDATE user_wallet_watchlist_items SET profile_id = user_id WHERE user_id is not null;

ALTER TABLE public.user_nft_watchlist_items
    add constraint user_nft_watchlist_items_pid_symbol_collection_address__key
        unique (profile_id, symbol, collection_address, chain_id);
ALTER TABLE public.user_wallet_watchlist_items
    add unique (profile_id, address);

ALTER TABLE user_token_watchlist_items ALTER COLUMN user_id drop not null;
ALTER TABLE user_nft_watchlist_items ALTER COLUMN user_id drop not null;
ALTER TABLE user_wallet_watchlist_items ALTER COLUMN user_id drop not null;

-- +migrate Down
ALTER TABLE user_wallet_watchlist_items DROP COLUMN profile_id;
ALTER TABLE user_nft_watchlist_items DROP COLUMN profile_id;
ALTER TABLE user_token_watchlist_items DROP COLUMN profile_id;

ALTER TABLE public.user_nft_watchlist_items drop constraint if exists user_nft_watchlist_items_user_id_symbol_collection_address__key;
ALTER TABLE public.user_nft_watchlist_items
    add constraint user_nft_watchlist_items_user_id_symbol_collection_address__key
        unique (user_id, symbol, collection_address, chain_id);
ALTER TABLE public.user_wallet_watchlist_items
    add unique (user_id, address);

ALTER TABLE user_token_watchlist_items RENAME TO user_watchlist_items;