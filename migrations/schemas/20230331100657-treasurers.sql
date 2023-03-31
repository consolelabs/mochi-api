
-- +migrate Up
create table if not exists treasurer_requests(
  id serial primary key,
  vault_id integer,
  guild_id text,
  user_discord_id text,
  message text,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  deleted_at timestamp
);

create table if not exists treasurers (
  id serial primary key,
  vault_id integer,
  guild_id text,
  user_discord_id text,
  request_id integer,
  message text,
  constraint treasurers_vault_id_fkey foreign key (vault_id) references vaults (id),
  unique (vault_id, guild_id, user_discord_id),
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
-- +migrate Down
drop table if exists treasurer_requests;
drop table if exists treasurers;
