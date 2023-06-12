
-- +migrate Up
DROP table if exists earn_infos cascade;

DROP table if exists user_earns;

DROP type if exists earn_status cascade;

CREATE TYPE profile_airdrop_campaign_status AS ENUM ('new', 'skipped', 'done', 'success', 'failure');

CREATE TABLE airdrop_campaigns (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    detail TEXT,
    prev_airdrop_campaign_id INT REFERENCES airdrop_campaigns(id),
    created_at TIMESTAMP with time zone default now(),
    updated_at TIMESTAMP with time zone default now(),
    deadline_at TIMESTAMP with time zone
);

CREATE TABLE profile_airdrop_campaigns (
    id SERIAL PRIMARY KEY,
    profile_id TEXT,
    airdrop_campaign_id INT REFERENCES airdrop_campaigns(id) on delete cascade,
    status profile_airdrop_campaign_status,
    is_favorite BOOLEAN,
    created_at TIMESTAMP with time zone default now(),
    updated_at TIMESTAMP with time zone default now(),
    UNIQUE (profile_id, airdrop_campaign_id)
);
-- +migrate Down
DROP table if exists airdrop_campaigns cascade;

DROP table if exists profile_airdrop_campaigns;

DROP type if exists profile_airdrop_campaign_status cascade;
