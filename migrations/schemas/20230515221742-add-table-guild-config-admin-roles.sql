
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_admin_roles (
    id SERIAL PRIMARY KEY,
    guild_id TEXT NOT NULL,
    role_id TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS guild_config_admin_roles_idx on guild_config_admin_roles(guild_id, role_id);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_admin_roles;
