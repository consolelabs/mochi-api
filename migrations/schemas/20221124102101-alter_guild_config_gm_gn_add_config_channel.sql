
-- +migrate Up
drop index guild_config_gm_gn_guild_id_uidx;
-- +migrate Down
CREATE UNIQUE INDEX guild_config_gm_gn_guild_id_uidx ON guild_config_gm_gn (guild_id);