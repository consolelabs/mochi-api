
-- +migrate Up
create table reaction_role_configs
(
    id             uuid  default uuid_generate_v4(),
    guild_id       text
        constraint reaction_roles_discord_guilds_id_fk
            references discord_guilds,
    channel_id     text,
    title          text,
    title_url      text,
    thumbnail_url  text,
    description    text,
    footer_image   text,
    footer_message text,
    message_id     text,
    reaction_roles jsonb default '[]'::jsonb
);

-- +migrate Down
DROP TABLE IF EXISTS "reaction_role_configs";