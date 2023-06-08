
-- +migrate Up
CREATE TYPE earn_status AS ENUM ('new', 'skipped', 'done', 'success', 'failure');

CREATE TABLE earn_infos (
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    detail TEXT,
    prev_earn_id INT REFERENCES earn_infos(id),
    created_at TIMESTAMP with time zone default now(),
    updated_at TIMESTAMP with time zone default now(),
    deadline_at TIMESTAMP with time zone
);

CREATE TABLE user_earns (
    id SERIAL PRIMARY KEY,
    user_id TEXT,
    earn_id INT REFERENCES earn_infos(id) on delete cascade,
    status earn_status,
    is_favorite BOOLEAN,
    UNIQUE (user_id, earn_id)
);
-- +migrate Down
DROP table if exists earn_infos cascade;

DROP table if exists user_earns;

DROP type earn_status cascade;
