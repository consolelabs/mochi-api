
-- +migrate Up
alter table offchain_tip_bot_tokens add column address text;
alter table offchain_tip_bot_tokens add column is_native boolean default false;
-- +migrate Down
alter table offchain_tip_bot_tokens drop column address;
alter table offchain_tip_bot_tokens drop column is_native;
