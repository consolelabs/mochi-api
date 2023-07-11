
-- +migrate Up
create table if not exists guild_config_log_channels (
    id serial primary key,
    guild_id text,
    channel_id text,
    log_type text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create unique index guild_config_log_channels_guild_id_log_type_uindex on guild_config_log_channels (guild_id, log_type);
-- +migrate Down
drop table if exists guild_config_log_channels;
