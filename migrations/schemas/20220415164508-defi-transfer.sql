
-- +migrate Up
create table if not exists tokens (
    id serial NOT NULL primary key,
	address text,
	symbol text,
	chain_id INTEGER,
	decimals INTEGER,
	discord_bot_supported bool
);

create table if not exists discord_bot_transactions (
    tx_hash text primary key,
    from_discord_id text,
    to_discord_id text,
    to_address text,
    amount numeric,
    reason text,
    type text,
    guild_id text,
    channel_id text,
    token_id INTEGER,
    created_at timestamptz DEFAULT now(),
    CONSTRAINT fk_token
      FOREIGN KEY(token_id) 
	  REFERENCES tokens(id)
);

-- +migrate Down
DROP TABLE IF EXISTS "discord_bot_transactions";
DROP TABLE IF EXISTS "tokens";
