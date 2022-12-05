
-- +migrate Up
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS image TEXT;
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS message TEXT;

-- +migrate Down
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS image;
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS message;