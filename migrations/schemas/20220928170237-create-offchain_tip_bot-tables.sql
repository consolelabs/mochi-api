-- +migrate Up
CREATE TABLE IF NOT EXISTS offchain_tip_bot_tokens (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  token_id TEXT NOT NULL,
  token_name varchar NOT NULL,
  token_symbol varchar NOT NULL,
  icon varchar,
  status smallint NOT NULL DEFAULT 0,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  UNIQUE(token_id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_chains (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  chain_id varchar NOT NULL,
  chain_name varchar NOT NULL,
  currency varchar NOT NULL,
  rpc_url varchar NOT NULL,
  explorer_url varchar NOT NULL,
  status smallint NOT NULL DEFAULT 0,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  UNIQUE(chain_id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_tokens_chains (
  token_id UUID NOT NULL,
  chain_id UUID NOT NULL,
  FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id),
  FOREIGN KEY (chain_id) REFERENCES offchain_tip_bot_chains(id),
  UNIQUE(token_id, chain_id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_contracts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  chain_id UUID NOT NULL,
  contract_address text NOT NULL,
  status smallint NOT NULL DEFAULT 0,
  assign_status smallint NOT NULL DEFAULT 0,
  centralize_wallet text NOT NULL,
  sweeped_time timestamptz DEFAULT now(),
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  FOREIGN KEY (chain_id) REFERENCES offchain_tip_bot_chains(id),
  UNIQUE(chain_id, contract_address)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_assign_contract (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  token_id UUID NOT NULL,
  chain_id UUID NOT NULL,
  user_id text NOT NULL,
  contract_id UUID NOT NULL,
  status int,
  expired_time timestamptz,
  FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id),
  FOREIGN KEY (chain_id) REFERENCES offchain_tip_bot_chains(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (contract_id) REFERENCES offchain_tip_bot_contracts(id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_user_balances (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id text NOT NULL,
  token_id UUID NOT NULL,
  amount float8,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_activity_logs (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id text NOT NULL,
  guild_id text NOT NULL,
  action varchar,
  receiver varchar,
  number_receivers int,
  duration int,
  token_id UUID NOT NULL,
  amount float8,
  full_command varchar,
  status varchar,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (token_id) REFERENCES offchain_tip_bot_tokens(id)
);

CREATE TABLE IF NOT EXISTS offchain_tip_bot_transfer_histories (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  sender_id text NOT NULL,
  receiver_id text NOT NULL,
  guild_id text NOT NULL,
  log_id UUID NOT NULL,
  status varchar,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL,
  FOREIGN KEY (sender_id) REFERENCES users(id),
  FOREIGN KEY (receiver_id) REFERENCES users(id),
  FOREIGN KEY (log_id) REFERENCES offchain_tip_bot_activity_logs(id)
);

-- +migrate Down
DROP TABLE IF EXISTS offchain_tip_bot_transfer_histories;
DROP TABLE IF EXISTS offchain_tip_bot_activity_logs;
DROP TABLE IF EXISTS offchain_tip_bot_user_balances;
DROP TABLE IF EXISTS offchain_tip_bot_assign_contract;
DROP TABLE IF EXISTS offchain_tip_bot_contracts;
DROP TABLE IF EXISTS offchain_tip_bot_tokens_chains;
DROP TABLE IF EXISTS offchain_tip_bot_chains;
DROP TABLE IF EXISTS offchain_tip_bot_tokens;