
-- +migrate Up
ALTER TABLE user_token_support_requests ADD COLUMN token_name TEXT;
ALTER TABLE user_token_support_requests ADD COLUMN symbol TEXT;

-- +migrate Down
ALTER TABLE user_token_support_requests DROP COLUMN token_name;
ALTER TABLE user_token_support_requests DROP COLUMN symbol;

