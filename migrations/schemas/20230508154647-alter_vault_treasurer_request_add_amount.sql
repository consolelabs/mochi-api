
-- +migrate Up
alter table treasurer_requests add column amount varchar;
alter table treasurer_requests add column chain varchar;
alter table treasurer_requests add column token varchar;
alter table treasurer_requests add column address text;
-- +migrate Down
alter table treasurer_requests drop column amount;
alter table treasurer_requests drop column chain;
alter table treasurer_requests drop column token;
alter table treasurer_requests drop column address;