
-- +migrate Up
ALTER TABLE user_feedbacks ADD COLUMN IF NOT EXISTS message_id TEXT;

-- +migrate Down
ALTER TABLE user_feedbacks DROP COLUMN IF EXISTS message_id;
