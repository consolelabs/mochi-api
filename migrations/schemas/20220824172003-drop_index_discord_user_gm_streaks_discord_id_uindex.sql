
-- +migrate Up
DROP INDEX discord_user_gm_streaks_discord_id_uindex;
CREATE UNIQUE INDEX discord_user_gm_streaks_discord_id_uindex ON discord_user_gm_streaks (discord_id, guild_id);

-- +migrate Down
CREATE UNIQUE INDEX discord_user_gm_streaks_discord_id_uindex ON discord_user_gm_streaks (discord_id, guild_id);
DROP INDEX discord_user_gm_streaks_discord_id_uindex;