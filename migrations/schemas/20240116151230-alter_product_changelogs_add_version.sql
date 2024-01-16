-- +migrate Up
ALTER TABLE product_changelogs ADD COLUMN IF NOT EXISTS version text;

-- +migrate Down
ALTER TABLE product_changelogs DROP COLUMN IF EXISTS version;