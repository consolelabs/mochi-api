
-- +migrate Up
ALTER TABLE user_nft_balances ADD COLUMN IF NOT EXISTS profile_id text not null default '';
ALTER TABLE user_nft_balances ADD COLUMN IF NOT EXISTS updated_at timestamptz not null default (now() at time zone 'utc');
ALTER TABLE user_nft_balances ADD COLUMN IF NOT EXISTS metadata jsonb not null default '{}';
ALTER TABLE user_nft_balances DROP COLUMN IF EXISTS token_id;

ALTER TABLE user_nft_balances ADD CONSTRAINT user_nft_balances_collection_id_address UNIQUE (nft_collection_id, user_address);

-- +migrate Down
ALTER TABLE user_nft_balances ADD COLUMN IF NOT EXISTS token_id TEXT not null default '';
ALTER TABLE user_nft_balances DROP COLUMN IF EXISTS updated_at;
ALTER TABLE user_nft_balances DROP COLUMN IF EXISTS profile_id;
ALTER TABLE user_nft_balances DROP COLUMN IF EXISTS metadata;

ALTER TABLE user_nft_balances DROP CONSTRAINT IF EXISTS user_nft_balances_collection_id_address;
