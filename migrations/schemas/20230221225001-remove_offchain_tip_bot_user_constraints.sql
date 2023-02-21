
-- +migrate Up
ALTER TABLE offchain_tip_bot_user_balances DROP CONSTRAINT IF EXISTS offchain_tip_bot_user_balances_user_id_fkey;
ALTER TABLE offchain_tip_bot_activity_logs DROP CONSTRAINT IF EXISTS offchain_tip_bot_activity_logs_user_id_fkey;
ALTER TABLE offchain_tip_bot_transfer_histories DROP CONSTRAINT IF EXISTS offchain_tip_bot_transfer_histories_receiver_id_fkey;
ALTER TABLE offchain_tip_bot_transfer_histories ADD COLUMN IF NOT EXISTS tx_hash TEXT NOT NULL DEFAULT '';

-- +migrate Down
ALTER TABLE offchain_tip_bot_user_balances ADD CONSTRAINT offchain_tip_bot_user_balances_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE offchain_tip_bot_activity_logs ADD CONSTRAINT offchain_tip_bot_activity_logs_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE offchain_tip_bot_transfer_histories ADD CONSTRAINT offchain_tip_bot_transfer_histories_receiver_id_fkey FOREIGN KEY(receiver_id) REFERENCES users(id);
ALTER TABLE offchain_tip_bot_transfer_histories DROP COLUMN IF EXISTS tx_hash;
