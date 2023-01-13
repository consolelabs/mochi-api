
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_token_roles (
    id SERIAL NOT NULL PRIMARY KEY,
	token_id INTEGER NOT NULL REFERENCES tokens(id),
    guild_id TEXT NOT NULL,
    role_id TEXT NOT NULL,
    required_amount FLOAT8 NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS guild_id_idx on guild_config_token_roles(guild_id);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_token_roles;