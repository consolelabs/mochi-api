-- +migrate Up
ALTER TABLE contents RENAME TO product_metadata_copies;

-- +migrate Down
DROP table if exists product_metadata_copies;