
-- +migrate Up
ALTER TABLE offchain_tip_bot_contracts ADD COLUMN private_key TEXT;

-- +migrate Down
ALTER TABLE offchain_tip_bot_contracts DROP COLUMN private_key;
