
-- +migrate Up
CREATE TABLE IF NOT EXISTS discord_user_token_alerts (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	token_id TEXT NOT NULL,
	discord_id TEXT NOT NULL,
	price_set FLOAT NOT NULL,
    trend TEXT NOT NULL,
    device_id TEXT NOT NULL,
	is_enable BOOLEAN DEFAULT true,
	updated_at timestamptz NOT NULL DEFAULT NOW(),
	created_at timestamptz NOT NULL DEFAULT NOW(),
	FOREIGN KEY (device_id) REFERENCES discord_user_devices(id)
);

-- +migrate Down
DROP TABLE IF EXISTS discord_user_token_alerts;
