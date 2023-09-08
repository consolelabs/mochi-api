-- +migrate Up
alter table product_changelogs add column is_expired bool default false;
-- +migrate Down
alter table product_changelogs drop column is_expired;