/* Replace with your SQL commands */

-- +migrate Down
ALTER TABLE guild_config_reaction_roles
    RENAME TO reaction_role_configs;

CREATE TABLE "guild_config_reaction_roles" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "guild_id" text NOT NULL REFERENCES "discord_guilds" ("id"),
    "channel_id" text NOT NULL,
    "author" text,
    "author_avatar" text,
    "title" text,
    "title_url" text,
    "header_image" text,
    "message" text NOT NULL,
    "footer_image" text,
    "footer_message" text,
    "footer_image_small" text,
    "reaction_roles" JSONB DEFAULT '[]'::JSONB
);
