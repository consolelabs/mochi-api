
-- +migrate Up
ALTER TABLE trade_offers DROP COLUMN to_address;
ALTER TABLE trade_offers RENAME COLUMN from_address TO owner_address;

-- +migrate Down
ALTER TABLE trade_offers ADD COLUMN IF NOT EXISTS to_address text;
ALTER TABLE trade_offers RENAME COLUMN owner_address to from_address;