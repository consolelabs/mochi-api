
-- +migrate Up
ALTER TABLE offchain_tip_bot_tokens drop constraint if exists offchain_tip_bot_tokens_token_id_key;
ALTER TABLE user_token_support_requests add column coin_gecko_id text not null default '';
ALTER TABLE user_token_support_requests add column decimal int not null default 0;
ALTER TABLE user_token_support_requests add column icon text;

-- +migrate Down
ALTER TABLE offchain_tip_bot_tokens add constraint offchain_tip_bot_tokens_token_id_key unique(token_id);
ALTER TABLE user_token_support_requests drop column coin_gecko_id;
ALTER TABLE user_token_support_requests drop column decimal;
ALTER TABLE user_token_support_requests drop column icon;
