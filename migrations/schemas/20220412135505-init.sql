
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "discord_guilds" (
  "id" bigint PRIMARY KEY,
  "name" text,
  "bot_scopes" JSONB NOT NULL DEFAULT '[]'::JSONB,
  "alias" text
);


CREATE TABLE "guild_roles" (
  "id" bigint PRIMARY KEY,
  "guild_id" bigint NOT NULL REFERENCES "discord_guilds"("id"),
  "name" text NOT NULL
);

CREATE TABLE "users" (
    "id" bigint PRIMARY KEY,
    "referral_code" text NOT NULL DEFAULT "substring"(md5((random())::text), 0, 9),
    "invited_by" bigint,
    "username" text,
    "nickname" text,
    "in_discord_wallet_address" text,
    "in_discord_wallet_number" int8,
    "join_date" timestamptz
);

-- +migrate Down

DROP TABLE IF EXISTS "guild_roles";
DROP TABLE IF EXISTS "discord_guilds";

DROP TABLE IF EXISTS "users";

DROP EXTENSION IF EXISTS "uuid-ossp";
