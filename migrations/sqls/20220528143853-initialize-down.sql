/* Replace with your SQL commands */
ALTER TABLE "guild_users"
  DROP COLUMN "invited_by",
  DROP COLUMN "nickname";

ALTER TABLE "users"
  ADD COLUMN "referal_code" text NOT NULL DEFAULT "substring"(md5((random())::text), 0, 9),
  ADD COLUMN "invited_by" bigint;

DROP TABLE "invite_histories";
