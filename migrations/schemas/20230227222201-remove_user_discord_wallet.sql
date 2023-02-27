
-- +migrate Up
ALTER TABLE users DROP COLUMN in_discord_wallet_address;
ALTER TABLE users DROP COLUMN in_discord_wallet_number;

-- +migrate Down
ALTER TABLE users ADD COLUMN in_discord_wallet_address TEXT;
ALTER TABLE users ADD COLUMN in_discord_wallet_number BIGINT;
