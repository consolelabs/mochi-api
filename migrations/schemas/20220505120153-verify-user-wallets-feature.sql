
-- +migrate Up
create table guild_config_wallet_verification_messages (
	guild_id text primary key,
	verify_channel_id text not null,
	content text,
	embedded_message json,
	created_at timestamp default now()
);

create table discord_wallet_verifications (
	user_discord_id text not null,
	guild_id text not null,
	code text,
	created_at timestamp default now(),
	primary key(user_discord_id, guild_id)
);

create table user_wallets (
	user_discord_id text not null,
	guild_id text not null,
	address text not null,
	chain_type varchar default 'evm',
	created_at timestamp default now(),
	primary key (user_discord_id, guild_id),
	foreign key (user_discord_id) references users (id) on delete cascade,
	foreign key (guild_id) references discord_guilds (id) on delete cascade
);

-- +migrate Down
drop table user_wallets;
drop table discord_wallet_verifications;
drop table guild_config_wallet_verification_messages;