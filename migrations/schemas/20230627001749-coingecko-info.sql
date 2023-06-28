
-- +migrate Up
-- update default value
alter table coingecko_supported_tokens drop column coingecko_info;

create table coingecko_infos (
  id text primary key,
  info jsonb not null default '{}'::jsonb
);

-- +migrate Down
alter table coingecko_supported_tokens add column coingecko_info jsonb not null default '{}'::jsonb;
drop table coingecko_info;
