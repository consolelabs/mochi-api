
-- +migrate Up
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS channel_id text;
ALTER TABLE offchain_tip_bot_activity_logs ADD COLUMN IF NOT EXISTS fail_reason text;
ALTER TABLE offchain_tip_bot_activity_logs ALTER COLUMN receiver TYPE varchar(255)[] using array[receiver];
CREATE UNIQUE INDEX offchain_tip_bot_user_balances_user_id_token_id_uidx ON offchain_tip_bot_user_balances (user_id, token_id);
-- +migrate Down
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS channel_id;
ALTER TABLE offchain_tip_bot_activity_logs DROP COLUMN IF EXISTS fail_reason;
ALTER TABLE offchain_tip_bot_activity_logs ALTER COLUMN receiver TYPE varchar using receiver::varchar;
DROP INDEX IF EXISTS offchain_tip_bot_user_balances_user_id_token_id_uidx;