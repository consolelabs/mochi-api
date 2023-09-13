
-- +migrate Up
create table if not exists "friend_tech_key_watchlist_items" (
  "id" serial not null primary key,
  "key_address" text,
  "profile_id" text,
  "increase_alert_at" integer,
  "decrease_alert_at" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);
alter table "friend_tech_key_watchlist_items" add constraint "unique_key_address_profile_id" unique ("key_address", "profile_id");

-- +migrate Down
alter table "friend_tech_key_watchlist_items" drop constraint "unique_key_address_profile_id";
drop table if exists "friend_tech_key_watchlist_items";
