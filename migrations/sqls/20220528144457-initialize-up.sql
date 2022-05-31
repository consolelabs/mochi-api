/* Replace with your SQL commands */

-- +migrate Up
create table reaction_role_configs
(
    id             uuid  default uuid_generate_v4(),
    guild_id       text
        constraint reaction_roles_discord_guilds_id_fk
            references discord_guilds,
    message_id     text,
    reaction_roles jsonb default '[]'::jsonb
);