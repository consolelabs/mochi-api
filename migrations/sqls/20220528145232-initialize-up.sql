/* Replace with your SQL commands */

-- +migrate Up
create table whitelist_campaigns (
id serial NOT NULL primary key,
name text NOT NULL,
guild_id text not null REFERENCES discord_guilds (id),
created_at timestamp default now(),
constraint name_guild_id_unique UNIQUE (name, guild_id)
);

create table whitelist_campaign_users (
address text not null,
discord_id text not null,
notes text,
whitelist_campaign_id integer not null REFERENCES whitelist_campaigns (id),
created_at timestamp default now(),
constraint whitelist_campaign_users_unique UNIQUE (address, discord_id, whitelist_campaign_id)
);