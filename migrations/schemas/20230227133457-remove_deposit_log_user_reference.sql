
-- +migrate Up
ALTER TABLE offchain_tip_bot_deposit_logs DROP CONSTRAINT offchain_tip_bot_deposit_logs_user_id_fkey;

-- +migrate Down
ALTER TABLE offchain_tip_bot_deposit_logs ADD CONSTRAINT offchain_tip_bot_deposit_logs_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);
