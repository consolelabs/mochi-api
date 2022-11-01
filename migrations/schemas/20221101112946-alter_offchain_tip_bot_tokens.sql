
-- +migrate Up
ALTER TABLE offchain_tip_bot_tokens ADD COLUMN IF NOT EXISTS coin_gecko_id text;
-- +migrate Down
ALTER TABLE offchain_tip_bot_tokens DROP COLUMN IF EXISTS coin_gecko_id;
