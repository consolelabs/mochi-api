-- +migrate Up
ALTER TABLE emojis RENAME TO product_metadata_emojis;

-- +migrate Down
DROP table if exists product_metadata_emojis;