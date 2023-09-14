
-- +migrate Up
alter table product_hashtags add column telegram_description text;
alter table product_hashtags add column email_title text;
alter table product_hashtags add column email_subject text;
-- +migrate Down
alter table product_hashtags drop column telegram_description;
alter table product_hashtags drop column email_title;
alter table product_hashtags drop column email_subject;
