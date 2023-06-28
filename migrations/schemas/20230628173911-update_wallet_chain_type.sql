
-- +migrate Up
UPDATE user_wallet_watchlist_items SET chain_type = 'evm' where chain_type = 'eth';

-- +migrate Down
UPDATE user_wallet_watchlist_items SET chain_type = 'eth' where chain_type = 'evm';
