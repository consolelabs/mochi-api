
-- +migrate Up
CREATE TABLE IF NOT EXISTS offchain_tip_bot_user_balance_snapshots (
	user_id TEXT NOT NULL,
	token_id UUID NOT NULL,
	token_symbol TEXT,
	action TEXT,
	changed_amount FLOAT8,
	amount FLOAT8,
	created_at timestamptz DEFAULT NOW(),
	PRIMARY KEY (user_id, token_id, created_at)
);

-- +migrate Down
DROP TABLE IF EXISTS offchain_tip_bot_user_balance_snapshots;
