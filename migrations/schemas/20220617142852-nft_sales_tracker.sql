
-- +migrate Up
CREATE TABLE IF NOT EXISTS nft_sales_trackers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  contract_address TEXT NOT NULL,
  platform TEXT NOT NULL,
  sales_config_id UUID NOT NULL,
  FOREIGN KEY (sales_config_id) REFERENCES guild_config_sales_trackers(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS nft_sales_trackers;

