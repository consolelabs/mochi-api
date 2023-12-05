
-- +migrate Up
alter table chains add column type varchar(255);
-- +migrate Down
alter table chains drop column type;
