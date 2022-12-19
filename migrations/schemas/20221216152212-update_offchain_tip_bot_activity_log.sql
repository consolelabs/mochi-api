
-- +migrate Up
ALTER TABLE offchain_tip_bot_activity_logs
	ALTER COLUMN guild_id DROP NOT NULL,
	ALTER COLUMN user_id DROP NOT NULL,
	ALTER COLUMN channel_id DROP NOT NULL,
	ADD COLUMN sender_address TEXT;

ALTER TABLE offchain_tip_bot_transfer_histories
	ALTER COLUMN guild_id DROP NOT NULL,
	ALTER COLUMN sender_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE offchain_tip_bot_activity_logs
	ALTER COLUMN guild_id SET NOT NULL,
	ALTER COLUMN user_id SET NOT NULL,
	ALTER COLUMN channel_id SET NOT NULL,
	DROP COLUMN sender_address;

ALTER TABLE offchain_tip_bot_transfer_histories
	ALTER COLUMN guild_id SET NOT NULL,
	ALTER COLUMN sender_id SET NOT NULL;
