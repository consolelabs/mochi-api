
-- +migrate Up
CREATE TABLE "discord_guild_stat_channels" (
    "id" uuid default uuid_generate_v4() not null constraint discord_guild_stat_channels_pk primary key,
    "guild_id" text NOT NULL REFERENCES "discord_guilds" ("id"),
    "channel_id" text,
    "count_type" text,
    "created_at" timestamp with time zone default now(),
    "updated_at" timestamp with time zone default now()
);
-- +migrate Down
DROP TABLE IF EXISTS "discord_guild_stat_channels";