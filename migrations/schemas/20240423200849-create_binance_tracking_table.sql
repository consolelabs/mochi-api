
-- +migrate Up
create table if not exists binance_trackings (
    id serial primary key,
    profile_id text,
    spot_last_time timestamp
);

create table binance_spot_transactions (
    id serial,
    profile_id text,
    symbol text,
    pair text,
    order_id bigint not null,
    order_list_id integer,
    client_order_id text,
    price text,
    orig_qty text,
    executed_qty text,
    cumulative_quote_qty text,
    status text,
    time_in_force text,
    type text,
    side text,
    stop_price text,
    iceberg_qty text,
    time bigint,
    update_time bigint,
    is_working boolean,
    orig_quote_order_qty text,
    working_time timestamp,
    self_trade_prevention_mode text,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
drop table if exists binance_trackings;
drop table if exists binance_spot_transactions;
