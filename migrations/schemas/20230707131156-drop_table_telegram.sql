
-- +migrate Up
drop table if exists user_telegrams;
drop table if exists user_telegram_discord_associations;
-- +migrate Down
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

CREATE TABLE IF NOT EXISTS user_telegram_discord_associations (
	telegram_id INTEGER NOT NULL PRIMARY KEY,
	discord_id TEXT NOT NULL
);

CREATE UNIQUE INDEX user_telegram_discord_associations_discord_id_uidx ON user_telegram_discord_associations(discord_id);