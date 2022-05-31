/* Replace with your SQL commands */

-- +migrate Down
alter table guild_config_wallet_verification_messages drop column discord_message_id;