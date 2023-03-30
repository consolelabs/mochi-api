
-- +migrate Up
create table if not exists vault_infos (
    id serial primary key,
    description text,
    mod_step text,
    normal_step text,
    instruction_link text,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

create table if not exists vault_configs (
    id serial primary key,
    guild_id text,
    channel_id text,
    unique (guild_id, channel_id),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);
create unique index if not exists idx_vault_guild_id_name on vaults (guild_id, name);
-- +migrate Down
drop table if exists vault_infos;
drop table if exists vault_configs;
drop index if exists idx_vault_guild_id_name;