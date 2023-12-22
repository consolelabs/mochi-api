
-- +migrate Up
ALTER TABLE user_payment_settings DROP COLUMN default_token_id;

-- +migrate Down
ALTER TABLE user_payment_settings ADD COLUMN default_token_id text;