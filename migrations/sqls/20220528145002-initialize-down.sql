/* Replace with your SQL commands */

-- +migrate Down
DROP TABLE IF EXISTS guild_config_tokens;

ALTER TABLE tokens
	DROP COLUMN IF EXISTS coin_market_cap_id,
	DROP COLUMN IF EXISTS name;

