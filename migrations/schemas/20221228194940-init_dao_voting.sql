
-- +migrate Up
create type vote_option as enum (
    'nft_collection',
    'crypto_token'
);

create type vote_choice as enum (
    'Yes',
    'No',
    'Abstain'
);

create table if not exists guild_config_dao_proposal (
    id serial not null primary KEY,
    guild_id text,
    proposal_channel_id text,
    guideline_channel_id text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists dao_proposal (
    id serial not null primary key,
    guild_id text,
    guild_config_dao_proposal_id integer not null references guild_config_dao_proposal(id),
    voting_channel_id text,
    discussion_channel_id text,
    creator_id text not null references users(id),
    title text,
    description text,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    closed_at timestamptz
);

create table if not exists dao_vote_option (
    id serial not null primary key,
    type vote_option,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists dao_proposal_vote_option (
    id serial,
    proposal_id integer not null references dao_proposal(id),
    vote_option_id integer references dao_vote_option(id),
    address text,
    chain_id integer,
    symbol text,
    required_amount integer,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists dao_vote (
    id serial,
    proposal_id integer not null references dao_proposal(id),
    user_id text not null references users(id),
    choice vote_choice,
    point float8,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +migrate Down
drop enum if exists vote_choice;
drop enum if exists vote_option;
drop table if exists dao_vote;
drop table if exists dao_proposal_vote_option;
drop table if exists dao_vote_option;
drop table if exists dao_proposal;
drop table if exists guild_config_dao_proposal;