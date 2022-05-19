
-- +migrate Up
alter table guild_config_wallet_verification_messages add column discord_message_id text;

-- +migrate Down
alter table guild_config_wallet_verification_messages drop column discord_message_id;