
-- +migrate Up
create table if not exists user_telegrams (
    id serial primary key,
    chat_id integer not null,
    username text not null,
    first_name text,
    last_name text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
create unique index user_telegrams_chat_id_username_uidx ON user_telegrams (chat_id, username);
-- +migrate Down
drop table if exists user_telegrams;