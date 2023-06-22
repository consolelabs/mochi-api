
-- +migrate Up
alter table treasurer_requests add column message_url text;
alter table treasurer_submissions add column message_url text;
-- +migrate Down
alter table treasurer_requests drop column message_url;
alter table treasurer_submissions drop column message_url;
