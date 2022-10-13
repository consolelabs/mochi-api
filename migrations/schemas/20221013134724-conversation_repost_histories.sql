
-- +migrate Up
CREATE TABLE IF NOT EXISTS conversation_repost_histories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    guild_id TEXT,
    origin_channel_id TEXT,
    origin_start_message_id TEXT,
    origin_stop_message_id TEXT,
    repost_channel_id TEXT,
    created_at timestamp default now()
);

create unique index conversation_repost_histories_guild_id_origin_channel_id_origin_start_message_id_uindex
    on conversation_repost_histories (guild_id, origin_channel_id, origin_start_message_id);
ALTER TABLE guild_config_repost_reactions ADD COLUMN IF NOT EXISTS reaction_type text default 'message';
-- +migrate Down
DROP TABLE IF EXISTS conversation_repost_histories;
ALTER TABLE guild_config_repost_reactions DROP COLUMN IF EXISTS reaction_type;