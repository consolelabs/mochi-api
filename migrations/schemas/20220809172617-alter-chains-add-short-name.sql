
-- +migrate Up
ALTER TABLE chains ADD COLUMN short_name text unique;
UPDATE chains SET short_name = CASE
    WHEN id = 1 then 'eth'
    WHEN id = 56 then 'bsc'
    WHEN id = 250 then 'ftm'
end;

-- +migrate Down
ALTER TABLE chains DROP COLUMN short_name;