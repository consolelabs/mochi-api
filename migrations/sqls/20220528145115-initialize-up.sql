/* Replace with your SQL commands */

-- +migrate Up
CREATE TABLE "discord_guild_stats" (
    "id" uuid default uuid_generate_v4() not null constraint discord_guild_stats_pk primary key,
    "guild_id" text NOT NULL REFERENCES "discord_guilds" ("id"),

    "nr_of_members" INTEGER,
    "nr_of_users" INTEGER,
    "nr_of_bots" INTEGER,

    "nr_of_channels" INTEGER,
    "nr_of_text_channels" INTEGER,
    "nr_of_voice_channels" INTEGER,
    "nr_of_stage_channels" INTEGER,
    "nr_of_categories" INTEGER,
    "nr_of_announcement_channels" INTEGER,

    "nr_of_emojis" INTEGER,
    "nr_of_static_emojis" INTEGER,
    "nr_of_animated_emojis" INTEGER,

    "nr_of_stickers" INTEGER,
    "nr_of_standard_stickers" INTEGER,
    "nr_of_guild_stickers" INTEGER,

    "nr_of_roles" INTEGER,
    "created_at" timestamp with time zone default now()
);

CREATE UNIQUE INDEX discord_guild_stats_guild_id_uidx ON discord_guild_stats (guild_id);
