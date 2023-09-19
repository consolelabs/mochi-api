
-- +migrate Up
alter table profile_command_usages add column if not exists user_platform_id text not null default '';

-- +migrate Down
alter table profile_command_usages drop column if exists user_platform_id;
