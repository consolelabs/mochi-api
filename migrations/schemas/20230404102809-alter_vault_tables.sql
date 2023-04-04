
-- +migrate Up
create type treasurer_request_type as enum ('add', 'remove', 'transfer');
alter table treasurer_requests add column if not exists "type" treasurer_request_type;
alter table treasurer_requests add column if not exists "requester" text;

create type treasurer_submit_status as enum ('pending', 'approved', 'rejected');
create table if not exists treasurer_submissions (   
    id serial primary key,
    vault_id integer references vaults(id) on delete cascade,
    guild_id text,
    request_id integer references treasurer_requests(id) on delete cascade,
    submitter text,
    status treasurer_submit_status,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);

alter table treasurers drop column if exists "request_id";
alter table treasurers drop column if exists "message";
alter table treasurers add column if not exists "role" text;
-- +migrate Down
alter table treasurer_requests drop column if exists "type";
alter table treasurer_requests drop column if exists "requester";
drop type if exists treasurer_request_type;

drop table if exists treasurer_submissions;
drop type if exists treasurer_submit_status;

alter table treasurers add column if not exists "request_id" integer;
alter table treasurers add column if exists "message" text;
alter table treasurers drop column if exists "role";