/* Replace with your SQL commands */

-- +migrate Up
create table guild_custom_commands (
	id text not null,
	guild_id text not null REFERENCES discord_guilds (id),
	description text,
	actions json,
	cooldown int8,
	cooldown_duration int8,
	enabled boolean,
	roles_permissions json,
	channels_permissions json,
	primary key(id, guild_id)
);