-- +migrate Up
CREATE TABLE IF NOT EXISTS "tono_command_permissions" (
    id SERIAL PRIMARY KEY,
    code TEXT,
    discord_permission_flag DECIMAL,
    description text,
    need_dm BOOL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS "tono_command_permissions";