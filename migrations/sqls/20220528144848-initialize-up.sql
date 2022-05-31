/* Replace with your SQL commands */
-- +migrate Up
create table guild_config_default_roles
(
    id         uuid                     default uuid_generate_v4() not null
        constraint guild_config_default_roles_pk
            primary key,
    role_id    text,
    guild_id   text,
    created_at timestamp with time zone default now()
);

create unique index guild_config_default_roles_role_id_uindex
    on guild_config_default_roles (role_id);
