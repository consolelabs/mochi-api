-- +migrate Up
create table product_changelog_snapshots  (
     id SERIAL PRIMARY KEY,
     filename text,
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
);

-- +migrate Down
drop table product_changelog_snapshots;