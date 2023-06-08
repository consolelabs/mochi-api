-- +migrate Up
create table emojis  (
     id SERIAL PRIMARY KEY,
     code text,
     discord_id text,
     telegram_id text,
     twitter_id text,
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
);

-- +migrate Down
drop table emojis;