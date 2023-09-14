
-- +migrate Up
update "friend_tech_key_watchlist_items" set "increase_alert_at" = 0 where "increase_alert_at" is null;
update "friend_tech_key_watchlist_items" set "decrease_alert_at" = 0 where "decrease_alert_at" is null;
update "friend_tech_key_watchlist_items" set "created_at" = now() where "created_at" is null;
update "friend_tech_key_watchlist_items" set "updated_at" = now() where "updated_at" is null;

alter table "friend_tech_key_watchlist_items" alter column "increase_alert_at" set not null;
alter table "friend_tech_key_watchlist_items" alter column "increase_alert_at" set  default 0;
alter table "friend_tech_key_watchlist_items" alter column "decrease_alert_at" set not null;
alter table "friend_tech_key_watchlist_items" alter column "decrease_alert_at" set default 0;
alter table "friend_tech_key_watchlist_items" alter column "created_at" set not null;
alter table "friend_tech_key_watchlist_items" alter column "created_at" set default now();
alter table "friend_tech_key_watchlist_items" alter column "updated_at" set not null;
alter table "friend_tech_key_watchlist_items" alter column "updated_at" set default now();
-- +migrate Down
