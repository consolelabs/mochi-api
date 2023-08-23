
-- +migrate Up
alter table product_bot_commands add column discord_alias text;
alter table product_bot_commands add column telegram_alias text;
-- +migrate Down
alter table product_bot_commands drop column discord_alias;
alter table product_bot_commands drop column telegram_alias;
