
-- +migrate Up
ALTER TABLE token_infos RENAME COLUMN token TO id;

-- +migrate Down
ALTER TABLE token_infos RENAME COLUMN id TO token;
