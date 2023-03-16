
-- +migrate Up
ALTER TABLE user_token_balances ADD COLUMN user_discord_id TEXT;
ALTER TABLE user_token_balances DROP COLUMN user_address;
ALTER TABLE user_token_balances DROP COLUMN chain_type;
ALTER TABLE user_token_balances ALTER COLUMN balance TYPE NUMERIC;


-- +migrate Down
ALTER TABLE user_token_balances DROP COLUMN user_discord_id;
ALTER TABLE user_token_balances ADD COLUMN user_address TEXT;
ALTER TABLE user_token_balances ADD COLUMN chain_type TEXT;
ALTER TABLE user_token_balances ALTER COLUMN balance TYPE FLOAT8;