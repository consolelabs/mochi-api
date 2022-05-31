/* Replace with your SQL commands */

-- +migrate Up
DROP TABLE IF EXISTS "guild_config_reaction_roles";

ALTER TABLE reaction_role_configs
    RENAME TO guild_config_reaction_roles;