
-- +migrate Up

DROP TABLE IF EXISTS temp_associated_accounts;

CREATE TYPE "public"."profile_type" AS ENUM ('user', 'application', 'vault', 'application_vault');

CREATE TABLE "public"."profiles" (
    "id" text NOT NULL PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "profile_name" text,
    "avatar" text,
    "type" "public"."profile_type" DEFAULT 'user'::profile_type,
    "application_id" int4,
    "vault_id" int4,
    "active_score" int4 NOT NULL DEFAULT 0,
    "did_onboarding_telegram" bool NOT NULL DEFAULT false
);

CREATE TABLE "public"."associated_accounts" (
    "id" int8 NOT NULL,
    "profile_id" text NOT NULL,
    "platform" varchar(255) NOT NULL,
    "platform_identifier" varchar(255) NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "platform_metadata" json,
    CONSTRAINT "fk_profile_id" FOREIGN KEY ("profile_id") REFERENCES "public"."profiles"("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX aa_platform_platform_identifier_unique ON public.associated_accounts USING btree (platform, platform_identifier);
CREATE INDEX idx_associated_accounts ON public.associated_accounts USING btree (profile_id);
CREATE INDEX idx_associated_accounts_metadata_username ON public.associated_accounts USING gin (to_tsvector('english'::regconfig, (platform_metadata ->> 'username'::text)));
CREATE INDEX idx_onlyy_associated_accounts_profile_id ON public.associated_accounts USING btree (profile_id) INCLUDE (id, platform, platform_identifier, platform_metadata, created_at, updated_at);

-- +migrate Down
DROP TABLE IF EXISTS associated_accounts;
DROP TABLE IF EXISTS profiles;

DROP TYPE IF EXISTS profile_type;