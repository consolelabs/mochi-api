
-- +migrate Up
CREATE TABLE IF NOT EXISTS onchain_tip_bot_transactions (
	id SERIAL NOT NULL PRIMARY KEY,
	sender_discord_id TEXT NOT NULL REFERENCES users(id),
	recipient_discord_id TEXT NOT NULL REFERENCES users(id),
	recipient_address TEXT,
	guild_id TEXT NOT NULL REFERENCES discord_guilds(id),
	channel_id TEXT NOT NULL,
	amount FLOAT8 NOT NULL,
	token_symbol TEXT NOT NULL,
	each BOOLEAN NOT NULL DEFAULT FALSE,
	"all" BOOLEAN NOT NULL DEFAULT FALSE,
	transfer_type TEXT NOT NULL,
	full_command TEXT NOT NULL,
	duration INTEGER NOT NULL DEFAULT 0,
	message TEXT NOT NULL,
	image TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT 'pending',
	tx_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT now(),
	updated_at TIMESTAMPTZ DEFAULT now(),
	claimed_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS onchain_tip_bot_transactions;