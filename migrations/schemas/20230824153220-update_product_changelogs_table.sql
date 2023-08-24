-- +migrate Up
ALTER TABLE product_changelogs ALTER COLUMN product TYPE text;

-- +migrate Down
ALTER TABLE product_changelogs ALTER COLUMN product TYPE int;