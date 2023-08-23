-- +migrate Up
create table product_changelogs  (
     id SERIAL PRIMARY KEY,
     product int,
     title text,
     content text,
     github_url text,
     thumbnail_url text,
     created_at timestamp not null default now(),
     updated_at timestamp not null default now()
);

-- +migrate Down
drop table product_changelogs;

