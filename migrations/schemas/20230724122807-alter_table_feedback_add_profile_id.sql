
-- +migrate Up
alter table user_feedbacks add column if not exists profile_id text;
-- +migrate Down
alter table user_feedbacks drop column if exists profile_id;
