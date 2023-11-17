
-- +migrate Up
alter table offchain_tip_bot_tokens add column chain_id text;
-- +migrate Down
alter table offchain_tip_bot_tokens drop column chain_id;
