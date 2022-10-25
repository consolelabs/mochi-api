
-- +migrate Up
ALTER TABLE offchain_tip_bot_activity_logs DROP CONSTRAINT offchain_tip_bot_activity_logs_token_id_fkey;
-- +migrate Down
ALTER TABLE offchain_tip_bot_activity_logs ADD CONSTRAINT token_id FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id);