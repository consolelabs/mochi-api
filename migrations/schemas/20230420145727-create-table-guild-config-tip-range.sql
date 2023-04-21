
-- +migrate Up
CREATE TABLE IF NOT EXISTS guild_config_tip_ranges
(
    id serial primary key,
    guild_id text not null REFERENCES discord_guilds (id),
    min float,
    max float,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

CREATE UNIQUE INDEX guild_config_tip_range_uidx
    on guild_config_tip_ranges (guild_id);

-- +migrate Down
DROP TABLE IF EXISTS "guild_config_tip_ranges";

DROP INDEX guild_config_tip_range_uidx;
