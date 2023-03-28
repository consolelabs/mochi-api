
-- +migrate Up
ALTER TABLE user_token_support_requests ADD COLUMN guild_id TEXT NOT NULL;
ALTER TABLE user_token_support_requests DROP COLUMN token_name;

-- +migrate Down
ALTER TABLE user_token_support_requests DROP COLUMN guild_id;
ALTER TABLE user_token_support_requests ADD COLUMN token_name TEXT NOT NULL;

