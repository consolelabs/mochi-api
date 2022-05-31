/* Replace with your SQL commands */

-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "discord_guilds" (
  "id" text PRIMARY KEY,
  "name" text,
  "bot_scopes" JSONB NOT NULL DEFAULT '[]'::JSONB,
  "alias" text,
  "created_at" timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE "guild_roles" (
  "id" text PRIMARY KEY,
  "guild_id" text NOT NULL REFERENCES "discord_guilds"("id"),
  "name" text NOT NULL
);

CREATE TABLE "users" (
    "id" text PRIMARY KEY,
    "referral_code" text NOT NULL DEFAULT "substring"(md5((random())::text), 0, 9),
    "invited_by" text,
    "username" text,
    "in_discord_wallet_address" text,
    "in_discord_wallet_number" int8
);
