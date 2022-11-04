
-- +migrate Up
ALTER TABLE offchain_tip_bot_tokens ADD COLUMN IF NOT EXISTS service_fee float8;
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS service_fee float8;
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS fee_amount float8;
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS service_fee float8;
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS fee_amount float8;

-- +migrate Down
ALTER TABLE offchain_tip_bot_tokens DROP COLUMN IF EXISTS service_fee;
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS service_fee;
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS fee_amount;
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS service_fee;
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS fee_amount;
