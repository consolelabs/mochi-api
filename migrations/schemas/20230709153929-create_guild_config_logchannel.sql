
-- +migrate Up
create table if not exists guild_config_log_channel (
    id serial primary key,
    guild_id text,
    channel_id text,
    log_type text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

create unique index guild_config_log_channel_guild_id_log_type_channel_id_uindex on guild_config_log_channel (guild_id, log_type, channel_id);
-- +migrate Down
drop table if exists guild_config_log_channel;
