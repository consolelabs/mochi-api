
-- +migrate Up
create table if not exists product_bot_commands (
    id serial primary key,
    code text,
    discord_command text,
    telegram_command text,
    scope integer,
    description text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

insert into product_bot_commands (code, discord_command, telegram_command, scope, description) values 
    ('TIP', '</tip:1062577077708136499>', '/tip', 0, 'send coin to user or group of users'),
    ('PROFILE', '</profile:1062577078173696110>', '/profile', 0, 'check profile of user');
-- +migrate Down
drop table if exists product_bot_commands;