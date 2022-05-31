/* Replace with your SQL commands */
-- +migrate Down


DROP TABLE IF EXISTS "guild_config_invite_trackers";
DROP TABLE IF EXISTS "guild_config_reaction_roles";

DROP TABLE IF EXISTS "guild_user_roles";
DROP TABLE IF EXISTS "guild_users";