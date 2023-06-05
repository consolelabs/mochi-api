
-- +migrate Up
ALTER TABLE "user_wallet_watchlist_items" RENAME COLUMN "type"  TO "chain_type";
ALTER TABLE "user_wallet_watchlist_items" ADD COLUMN "type"  TEXT NOT NULL DEFAULT 'follow';
-- +migrate Down
ALTER TABLE "user_wallet_watchlist_items" DROP COLUMN "type";
ALTER TABLE "user_wallet_watchlist_items" RENAME COLUMN "chain_type"  TO "type";
