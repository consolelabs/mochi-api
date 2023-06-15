
-- +migrate Up
CREATE TYPE airdrop_campaign_status AS ENUM ('','live','ended','claimable','cancelled');

alter table airdrop_campaigns add column status airdrop_campaign_status;

alter table airdrop_campaigns add column reward_amount integer;

alter table airdrop_campaigns add column reward_token_symbol text;

ALTER TYPE profile_airdrop_campaign_status
    RENAME TO old_profile_airdrop_campaign_status;

CREATE TYPE profile_airdrop_campaign_status AS ENUM ('','ignored', 'joined', 'claimed', 'not_eligible');

ALTER TABLE profile_airdrop_campaigns
ALTER COLUMN status TYPE profile_airdrop_campaign_status
USING status::text::profile_airdrop_campaign_status;

ALTER TABLE profile_airdrop_campaigns
    ALTER COLUMN status SET DEFAULT '';

DROP TYPE old_profile_airdrop_campaign_status;

CREATE TABLE raw_airdrop_campaign_datas (
    id int,
    title varchar,
    description text,
    source varchar,
    link text,
    created_at TIMESTAMP with time zone default now(),
    updated_at TIMESTAMP with time zone default now()
);
-- +migrate Down

