/* Replace with your SQL commands */
-- +migrate Up
alter table guild_config_wallet_verification_messages add column discord_message_id text;
