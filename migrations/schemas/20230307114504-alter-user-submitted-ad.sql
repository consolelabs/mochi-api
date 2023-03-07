
-- +migrate Up
ALTER TABLE user_submitted_ads ADD COLUMN is_podtown_ad BOOLEAN DEFAULT TRUE;

-- +migrate Down
ALTER TABLE user_submitted_ads DROP COLUMN is_podtown_ad;