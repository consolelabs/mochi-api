-- +migrate Up
create table discord_user_gm_streaks
(
    guild_id         text not null,
    discord_id       text not null,
    streak_count     integer   default 0,
    total_count      integer   default 0,
    last_streak_date timestamp default now(),
    created_at       timestamp default now(),
    updated_at       timestamp default now(),
    constraint discord_user_gm_streaks_pk
        primary key (discord_id, guild_id)
);

create unique index discord_user_gm_streaks_discord_id_uindex
    on discord_user_gm_streaks (discord_id);