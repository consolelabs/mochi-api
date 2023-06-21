
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;

-- +migrate Down
DROP EXTENSION fuzzystrmatch;