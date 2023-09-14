
-- +migrate Up
alter table product_hashtags add column telegram_title varchar(255);
-- +migrate Down
alter table product_hashtags drop column telegram_title;
