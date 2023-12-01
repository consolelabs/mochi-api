
-- +migrate Up
alter table offchain_tip_bot_tokens add column decimal integer;
-- +migrate Down
alter table offchain_tip_bot_tokens drop column decimal;
