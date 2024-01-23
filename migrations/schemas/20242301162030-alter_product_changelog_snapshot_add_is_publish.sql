-- +migrate Up
ALTER TABLE product_changelog_snapshots ADD COLUMN IF NOT EXISTS is_public bool Default true;

-- +migrate Down
ALTER TABLE product_changelog_snapshots DROP COLUMN IF EXISTS is_public;