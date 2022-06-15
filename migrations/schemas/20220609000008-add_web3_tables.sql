-- +migrate Up
CREATE TYPE defi_types AS ENUM ('deposit', 'withdraw', 'tip', 'airdrop');

CREATE TYPE defi_roles AS ENUM ('sender', 'recipient');

CREATE TABLE IF NOT EXISTS user_defi_activity_logs (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	type DEFI_TYPES NOT NULL,
	role DEFI_ROLES NOT NULL,
	token_id INTEGER NOT NULL REFERENCES tokens(id),
	user_id TEXT NOT NULL REFERENCES users(id),
	guild_id TEXT REFERENCES discord_guilds(id),
	amount NUMERIC NOT NULL,
	created_at timestamptz DEFAULT now()
);

CREATE VIEW user_balances(user_id, token_id, balance) AS
SELECT
	user_id,
	token_id,
	SUM(amount)
FROM
	user_defi_activity_logs
GROUP BY
	user_id,
	token_id;

-- +migrate Down
DROP VIEW user_balances;

DROP TABLE IF EXISTS user_defi_activity_logs;

DROP TYPE defi_roles;

DROP TYPE defi_types;