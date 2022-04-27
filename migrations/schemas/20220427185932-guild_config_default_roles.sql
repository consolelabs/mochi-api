
-- +migrate Up
create table guild_config_default_roles
(
    role_id  text,
    guild_id text
);

-- +migrate Down
DROP TABLE IF EXISTS "guild_config_default_roles";
