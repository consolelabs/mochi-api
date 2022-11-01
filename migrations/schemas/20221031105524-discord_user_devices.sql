
-- +migrate Up
CREATE TABLE IF NOT EXISTS discord_user_devices (
	id TEXT NOT NULL,
	ios_noti_token TEXT NOT NULL, 
	updated_at timestamptz NOT NULL DEFAULT NOW(),
	created_at timestamptz NOT NULL DEFAULT NOW(),
    UNIQUE (id)
);

-- +migrate Down
DROP TABLE IF EXISTS discord_user_devices;
