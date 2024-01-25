-- +migrate Up
ALTER TABLE
    product_changelogs
ADD
    COLUMN seo_description TEXT;

-- +migrate Down
ALTER TABLE
    product_changelogs DROP COLUMN seo_description;