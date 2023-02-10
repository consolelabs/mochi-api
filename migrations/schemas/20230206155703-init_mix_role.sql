-- +migrate Up
CREATE TABLE IF NOT EXISTS mix_role_token_requirements (
    id SERIAL PRIMARY KEY,
	token_id INTEGER NOT NULL REFERENCES tokens(id),
    required_amount FLOAT8 NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS mix_role_nft_requirements (
    id SERIAL PRIMARY KEY,
	nft_collection_id UUID NOT NULL REFERENCES nft_collections(id),
	required_amount INTEGER NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS guild_config_mix_roles (
    id SERIAL PRIMARY KEY,
    guild_id TEXT NOT NULL,
    role_id TEXT NOT NULL,
    nft_requirement_id INTEGER REFERENCES mix_role_nft_requirements(id),
    token_requirement_id INTEGER REFERENCES mix_role_token_requirements(id),
    required_level INTEGER NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS guild_id_role_id_idx on guild_config_mix_roles(guild_id, role_id);

-- +migrate Down
DROP TABLE IF EXISTS guild_config_mix_roles;
DROP TABLE IF EXISTS mix_role_token_requirements;
DROP TABLE IF EXISTS mix_role_nft_requirements;