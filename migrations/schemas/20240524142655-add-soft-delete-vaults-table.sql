-- +migrate Up
ALTER TABLE
    vaults
ADD
    COLUMN deleted_at timestamptz;

-- +migrate Down
ALTER TABLE
    vaults DROP COLUMN deleted_at;