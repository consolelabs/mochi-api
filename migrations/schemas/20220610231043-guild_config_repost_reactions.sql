-- +migrate Up
create table guild_config_repost_reactions
(
    id                uuid default uuid_generate_v4() not null
        constraint guild_config_repost_reactions_pk
            primary key,
    guild_id          text
        constraint guild_config_repost_reactions_discord_guilds_id_fk
            references discord_guilds,
    quantity          integer,
    emoji             text,
    repost_channel_id text
);

create unique index guild_config_repost_reactions_emoji_uindex
    on guild_config_repost_reactions (emoji);
-- +migrate Down
DROP TABLE IF EXISTS guild_config_repost_reactions;
