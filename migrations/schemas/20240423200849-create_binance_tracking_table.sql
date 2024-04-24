
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
    order_id integer,
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
    time timestamp,
    update_time timestamp,
    is_working boolean,
    orig_quote_order_qty text,
    working_time timestamp,
    self_trade_prevention_mode text
);

create unique index binance_spot_transactions_profile_id_order_id_index on binance_spot_transactions (profile_id, order_id);
-- +migrate Down
drop table if exists binance_trackings;
drop index if exists binance_spot_transactions_profile_id_order_id_index;
drop table if exists binance_spot_transactions;