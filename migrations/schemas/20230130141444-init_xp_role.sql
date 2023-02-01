
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_xp_roles (
    id SERIAL PRIMARY KEY,
    guild_id TEXT NOT NULL,
    role_id TEXT NOT NULL,
    required_xp INTEGER NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS guild_id_role_id_idx on guild_config_xp_roles(guild_id, role_id);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_xp_roles;