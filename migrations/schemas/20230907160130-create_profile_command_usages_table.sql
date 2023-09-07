-- +migrate Up
create table profile_command_usages  (
     id SERIAL PRIMARY KEY,
     profile_id text,
     command text,
     params text,
     platform text,
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
);

-- +migrate Down
drop table profile_command_usages;
