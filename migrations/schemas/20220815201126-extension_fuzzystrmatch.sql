
-- +migrate Up
CREATE EXTENSION fuzzystrmatch;

-- +migrate Down
DROP EXTENSION fuzzystrmatch;