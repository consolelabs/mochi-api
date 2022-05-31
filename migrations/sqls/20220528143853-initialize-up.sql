/* Replace with your SQL commands */

-- +migrate Up
ALTER TABLE "users"
  DROP COLUMN IF EXISTS "referral_code",
  DROP COLUMN IF EXISTS "invited_by";

ALTER TABLE "guild_users"
  ADD COLUMN "invited_by" bigint,
  ADD COLUMN "nickname" text;

CREATE TABLE "invite_histories" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "created_at" timestamptz DEFAULT now(),
  "guild_id" text NOT NULL REFERENCES "discord_guilds" ("id"),
  "user_id" text NOT NULL REFERENCES "users" ("id"),
  "invited_by" text NOT NULL REFERENCES "users" ("id"),
  "metadata" JSONB NOT NULL DEFAULT '{}'::JSONB
);