
-- +migrate Up
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS amount float8;
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS token text;
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS action text;
-- +migrate Down
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS amount;
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS token;
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS action;