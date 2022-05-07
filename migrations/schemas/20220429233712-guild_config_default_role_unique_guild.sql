
-- +migrate Up
alter table guild_config_default_roles
drop constraint guild_config_default_roles_pk;

alter table guild_config_default_roles
    add constraint guild_config_default_roles_pk
        unique (guild_id);

-- +migrate Down
alter table guild_config_default_roles
drop constraint guild_config_default_roles_pk;

alter table guild_config_default_roles
    add constraint guild_config_default_roles_pk
        unique (role_id);