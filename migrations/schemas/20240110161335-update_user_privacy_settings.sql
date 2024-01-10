
-- +migrate Up
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS tx;
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS social_accounts;
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS wallets;

ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS destination_wallet jsonb not null DEFAULT '{"enable": true,"target_group":"all"}';

-- +migrate Down
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS destination_wallet;

ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS tx jsonb not null DEFAULT '{}';
ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS social_accounts jsonb not null DEFAULT '{}';
ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS wallets jsonb not null DEFAULT '{}';