
-- +migrate Up
create table message_repost_histories
(
    id                uuid default uuid_generate_v4(),
    guild_id          text,
    repost_channel_id text,
    origin_message_id text,
    origin_channel_id text,
    created_at        timestamp with time zone default now()
);

create unique index message_repost_histories_origin_message_id_uindex
    on message_repost_histories (origin_message_id);

-- +migrate Down
DROP TABLE IF EXISTS message_repost_histories;