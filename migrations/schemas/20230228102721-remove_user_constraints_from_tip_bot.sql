
-- +migrate Up
ALTER TABLE offchain_tip_bot_transfer_histories DROP CONSTRAINT IF EXISTS offchain_tip_bot_transfer_histories_sender_id_fkey;
ALTER TABLE offchain_tip_bot_assign_contract DROP CONSTRAINT IF EXISTS offchain_tip_bot_assign_contract_user_id_fkey;
ALTER TABLE offchain_tip_bot_assign_contract_logs DROP CONSTRAINT IF EXISTS offchain_tip_bot_assign_contract_logs_user_id_fkey;

-- +migrate Down
ALTER TABLE offchain_tip_bot_transfer_histories ADD CONSTRAINT offchain_tip_bot_transfer_histories_sender_id_fkey FOREIGN KEY(sender_id) REFERENCES users(id);
ALTER TABLE offchain_tip_bot_assign_contract ADD CONSTRAINT offchain_tip_bot_assign_contract_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE offchain_tip_bot_assign_contract_logs ADD CONSTRAINT offchain_tip_bot_assign_contract_logs_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);