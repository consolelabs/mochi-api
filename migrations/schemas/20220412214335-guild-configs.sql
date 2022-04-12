
-- +migrate Up

CREATE TABLE "guild_users" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds" ("id"),
  "user_id" bigint NOT NULL REFERENCES "users" ("id")
);
CREATE UNIQUE INDEX "uidx_guild_users_guild_user_id" ON "guild_users"("guild_id", "user_id");

CREATE TABLE "guild_user_roles" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds" ("id"),
  "user_id" bigint NOT NULL REFERENCES "users" ("id"),
  "role_id" bigint NOT NULL REFERENCES "guild_roles" ("id")
);

CREATE UNIQUE INDEX "uidx_guild_user_roles_user_role_id" ON "guild_user_roles"("guild_id", "user_id", "role_id");

CREATE TABLE "guild_config_invite_trackers" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds" ("id"),
  "channel_id" bigint NOT NULL,
  "webhook_url" text
);

CREATE TABLE "guild_config_reaction_roles" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds" ("id"),
  "channel_id" bigint NOT NULL,
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


-- +migrate Down


DROP TABLE IF EXISTS "guild_config_invite_trackers";
DROP TABLE IF EXISTS "guild_config_reaction_roles";

DROP TABLE IF EXISTS "guild_user_roles";
DROP TABLE IF EXISTS "guild_users";
