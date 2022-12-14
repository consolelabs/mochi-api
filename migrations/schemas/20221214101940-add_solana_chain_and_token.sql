
-- +migrate Up
INSERT INTO chains (id, name, rpc, api_base_url, api_key, tx_base_url, currency, short_name, coin_gecko_id) VALUES (999,'Solana','','https://explorer.solana.com','','https://explorer.solana.com/tx','SOL','sol','solana');
INSERT INTO tokens (symbol, chain_id, decimals, discord_bot_supported, coin_gecko_id, name, guild_default, is_native) VALUES ('SOL', 999, 9, true, 'solana', 'Solana', false, true);
INSERT INTO offchain_tip_bot_chains (id, chain_id, chain_name, currency, rpc_url, explorer_url, status) VALUES ('f26c41dd-7625-4049-b886-8fa23424a37b', '0x12345', 'Solana', 'SOL', 'https://try-rpc.mainnet.solana.blockdaemon.tech', 'https://explorer.solana.com', 1);
INSERT INTO offchain_tip_bot_tokens (id, token_id, token_name, token_symbol, icon, status, coin_gecko_id) VALUES ('27153e8b-ce7c-4cee-aadd-f7291bb88b0a', '30', 'Solana', 'SOL', 'https://assets.coingecko.com/coins/images/4128/large/solana.png?1640133422', 1, 'solana');

-- +migrate Down
DELETE FROM chains WHERE id = 999;
DELETE FROM tokens WHERE chain_id = 999;
DELETE FROM offchain_tip_bot_chains WHERE id = 'f26c41dd-7625-4049-b886-8fa23424a37b';
DELETE FROM offchain_tip_bot_tokens WHERE id = '27153e8b-ce7c-4cee-aadd-f7291bb88b0a';
