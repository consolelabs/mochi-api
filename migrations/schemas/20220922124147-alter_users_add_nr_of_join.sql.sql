
-- +migrate Up
ALTER TABLE guild_config_wallet_verification_messages ADD COLUMN verify_role_id TEXT;

-- +migrate Down
ALTER TABLE guild_config_wallet_verification_messages DROP COLUMN verify_role_id;