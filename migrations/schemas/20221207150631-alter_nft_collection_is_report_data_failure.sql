
-- +migrate Up
ALTER TABLE nft_collections ADD COLUMN IF NOT EXISTS is_report_data_failure BOOLEAN DEFAULT false;

-- +migrate Down
ALTER TABLE nft_collections DROP COLUMN IF EXISTS is_report_data_failure;
