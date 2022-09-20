
-- +migrate Up
ALTER TABLE users ADD COLUMN nr_of_join INTEGER DEFAULT 0;
-- +migrate Down
ALTER TABLE users DROP COLUMN nr_of_join;