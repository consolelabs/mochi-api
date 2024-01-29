
-- +migrate Up
create table token_price_snapshot (
    id SERIAL PRIMARY KEY,
    symbol text unique ,
    chain text,
    price text not null,
    snapshot_time timestamp not null default now(),
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
-- +migrate Down
drop table token_price_snapshot;
