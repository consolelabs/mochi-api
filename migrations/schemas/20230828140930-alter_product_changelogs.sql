-- +migrate Up
alter table product_changelogs add column file_name text;
-- +migrate Down
alter table product_changelogs drop column file_name;