
-- +migrate Up
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS destination_wallet;

create type privacy_target_group as enum ('all', 'receivers', 'friends');
ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS tx_target_group privacy_target_group not null DEFAULT 'all';
ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS show_destination_wallet bool not null DEFAULT true;

-- +migrate Down
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS tx_target_group;
ALTER TABLE user_privacy_settings DROP COLUMN IF EXISTS show_destination_wallet;

ALTER TABLE user_privacy_settings ADD COLUMN IF NOT EXISTS destination_wallet jsonb not null DEFAULT '{"enable": true,"target_group":"all"}';
