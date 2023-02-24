
-- +migrate Up
ALTER TABLE discord_wallet_verifications DROP CONSTRAINT discord_wallet_verifications_pkey;
ALTER TABLE discord_wallet_verifications ADD CONSTRAINT discord_wallet_verifications_pkey PRIMARY KEY (user_discord_id, guild_id, code);

ALTER TABLE discord_wallet_verifications ADD COLUMN IF NOT EXISTS channel_id TEXT DEFAULT '';
ALTER TABLE discord_wallet_verifications ADD COLUMN IF NOT EXISTS message_id TEXT DEFAULT '';

ALTER TABLE user_wallet_watchlist_items DROP CONSTRAINT IF EXISTS user_wallet_watchlist_items_user_id_alias_key;

ALTER TABLE user_wallet_watchlist_items ADD COLUMN IF NOT EXISTS is_owner BOOLEAN DEFAULT FALSE;

-- +migrate Down
ALTER TABLE discord_wallet_verifications DROP CONSTRAINT discord_wallet_verifications_pkey;
ALTER TABLE discord_wallet_verifications ADD CONSTRAINT discord_wallet_verifications_pkey PRIMARY KEY (user_discord_id, guild_id);

ALTER TABLE discord_wallet_verifications DROP COLUMN IF EXISTS channel_id;
ALTER TABLE discord_wallet_verifications DROP COLUMN IF EXISTS message_id;

ALTER TABLE user_wallet_watchlist_items ADD CONSTRAINT user_wallet_watchlist_items_user_id_alias_key UNIQUE(user_id, alias);

ALTER TABLE user_wallet_watchlist_items DROP COLUMN IF EXISTS is_owner;
