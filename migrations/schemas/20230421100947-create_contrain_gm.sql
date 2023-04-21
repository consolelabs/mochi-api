
-- +migrate Up
alter table guild_config_gm_gn add column if not exists created_at timestamp not null default now();
alter table guild_config_gm_gn add column if not exists updated_at timestamp not null default now();
-- +migrate Down
alter table guild_config_gm_gn drop column if exists created_at;
alter table guild_config_gm_gn drop column if exists updated_at;
