
-- +migrate Up
alter table treasurer_requests add column if not exists is_approved boolean default false;
create table if not exists vault_transactions (
    id serial primary key,
    guild_id text,
    vault_id integer,
    action text,
    from_address text,
    to_address text,
    target text,
    amount text,
    token text,
    threshold text,
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now()
);
-- +migrate Down
alter table treasurer_requests drop column if exists is_approved;
drop table if exists vault_transactions;