
-- +migrate Up
alter table user_watchlist_items add column created_at timestamp default now();
alter table user_watchlist_items add column updated_at timestamp default now();
-- +migrate Down
alter table user_watchlist_items drop column created_at;
alter table user_watchlist_items drop column updated_at;
