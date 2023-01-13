
-- +migrate Up
drop table if exists nft_sales_trackers;
alter table guild_config_sales_trackers add column if not exists contract_address text,
 add column if not exists chain text,
 add if not exists created_at timestamp default now(),
 add if not exists updated_at timestamp default now();
-- +migrate Down
CREATE TABLE IF NOT EXISTS nft_sales_trackers (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  contract_address TEXT NOT NULL,
  platform TEXT NOT NULL,
  sales_config_id UUID NOT NULL,
  FOREIGN KEY (sales_config_id) REFERENCES guild_config_sales_trackers(id) ON DELETE CASCADE
);
alter table guild_config_sales_trackers drop column contract_address, drop column chain, drop column created_at, drop column updated_at;