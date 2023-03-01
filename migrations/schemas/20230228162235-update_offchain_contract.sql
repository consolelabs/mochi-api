
-- +migrate Up
ALTER TABLE offchain_tip_bot_contracts DROP COLUMN status;
ALTER TABLE offchain_tip_bot_contracts DROP COLUMN assign_status;
ALTER TABLE offchain_tip_bot_contracts DROP COLUMN centralize_wallet;

ALTER TABLE offchain_tip_bot_assign_contract DROP COLUMN status;
ALTER TABLE offchain_tip_bot_assign_contract ADD COLUMN created_at timestamptz DEFAULT NOW();

ALTER TABLE offchain_tip_bot_assign_contract_logs DROP COLUMN status;
ALTER TABLE offchain_tip_bot_assign_contract_logs ADD COLUMN created_at timestamptz DEFAULT NOW();

-- +migrate Down
ALTER TABLE offchain_tip_bot_contracts ADD COLUMN status SMALLINT NOT NULL DEFAULT 0;
ALTER TABLE offchain_tip_bot_contracts ADD COLUMN assign_status smallint NOT NULL DEFAULT 0;
ALTER TABLE offchain_tip_bot_contracts ADD COLUMN centralize_wallet TEXT NOT NULL DEFAULT '';

ALTER TABLE offchain_tip_bot_assign_contract ADD COLUMN status int;
ALTER TABLE offchain_tip_bot_assign_contract DROP COLUMN created_at;

ALTER TABLE offchain_tip_bot_assign_contract_logs ADD COLUMN status int;
ALTER TABLE offchain_tip_bot_assign_contract_logs DROP COLUMN created_at;