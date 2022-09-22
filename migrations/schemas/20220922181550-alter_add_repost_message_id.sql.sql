
-- +migrate Up
ALTER TABLE message_repost_histories ADD COLUMN repost_message_id TEXT;

-- +migrate Down
ALTER TABLE message_repost_histories DROP COLUMN repost_message_id;