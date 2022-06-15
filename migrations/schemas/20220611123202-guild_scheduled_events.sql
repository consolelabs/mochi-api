
-- +migrate Up
create table if not exists guild_scheduled_events (
	guild_id text not null references discord_guilds(id),
	event_id text not null,
	status int8 not null,
	primary key(guild_id, event_id)
);

-- +migrate Down
drop table if exists guild_scheduled_events;