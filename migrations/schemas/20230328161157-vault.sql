
-- +migrate Up
create type vault_threshold as enum ('50', '66', '75', '100');
create table if not exists vaults (
  id serial primary key,
  guild_id text,
  name text,
  threshold vault_threshold,
  wallet_address text,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now()
);
-- +migrate Down
drop table if exists vaults;
drop type if exists vault_threshold;