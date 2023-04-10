
-- +migrate Up
alter table user_tags drop column if exists mention_username;
alter table user_tags drop column if exists mention_role;
alter table user_tags add column if not exists is_active boolean default false;
-- +migrate Down
alter table user_tags drop column if exists is_active;
alter table user_tags add column if not exists mention_username boolean default false;
alter table user_tags add column if not exists mention_role boolean default false;

