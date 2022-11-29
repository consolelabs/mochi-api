
-- +migrate Up
create table if not exists offchain_tip_bot_config_notify (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    guild_id TEXT NOT NULL,
    channel_id TEXT NOT NULL,
    token TEXT NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
-- +migrate Down
drop table if exists offchain_tip_bot_config_notify;