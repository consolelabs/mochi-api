
-- +migrate Up
create table wallet_snapshot (
    id SERIAL PRIMARY KEY,
    wallet_address text,
    is_evm boolean,
    total_usd_balance text not null,
    snapshot_time timestamp not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
-- +migrate Down
drop table wallet_snapshot;
