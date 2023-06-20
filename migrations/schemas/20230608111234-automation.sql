
-- +migrate Up
CREATE TABLE IF NOT EXISTS auto_triggers
(
    id uuid DEFAULT uuid_generate_v4(),
    guild_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at timestamp with time zone not null default now(),
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_conditions
(
    id uuid DEFAULT uuid_generate_v4(),
    trigger_id TEXT NOT NULL,
    type_id TEXT NOT NULL,
    channel_id TEXT NULL,
    index INTEGER NOT NULL,
    platform TEXT NOT NULL,
    child_id TEXT NULL,
    updated_at timestamp with time zone not null default now(),
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_condition_values
(
    id uuid DEFAULT uuid_generate_v4(),
    condition_id TEXT NOT NULL,
    child_id TEXT NULL,
    type TEXT NOT NULL,
    index INTEGER NOT NULL,
    operator TEXT NOT NULL,
    matches TEXT NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_condition_types
(
    id uuid DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_condition_type_presets
(
    id uuid DEFAULT uuid_generate_v4(),
    type_id TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_actions
(
    id uuid DEFAULT uuid_generate_v4(),
    user_ids TEXT NULL,
    trigger_id TEXT NOT NULL,
    type_id TEXT NOT NULL,
    channel_ids NULL,
    index INTEGER NOT NULL,
    action_data TEXT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NULL,
    embed_id TEXT NULL,
    then_action_id TEXT NULL,
    limit_per_user INTEGER default 1,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_action_types
(
    id uuid DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_embeds
(
    id uuid DEFAULT uuid_generate_v4(),
    action_id TEXT NOT NULL,
    author_id TEXT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    color TEXT NOT NULL,
    image_id TEXT NULL,
    thumbnail_id TEXT NULL,
    footer_id TEXT NULL,
    video_id TEXT NULL,
    url TEXT NULL,
    type TEXT NULL,
    fields TEXT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_embed_images
(
    id uuid DEFAULT uuid_generate_v4(),
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    proxy_url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_embed_videos
(
    id uuid DEFAULT uuid_generate_v4(),
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at timestamp with time zone default now()
);

CREATE TABLE IF NOT EXISTS auto_embed_footers
(
    id uuid DEFAULT uuid_generate_v4(),
    embed_id TEXT NOT NULL,
    text TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at timestamp with time zone default now()
);


CREATE TABLE IF NOT EXISTS auto_action_histories
(
    id uuid DEFAULT uuid_generate_v4(),
    user_id TEXT NOT NULL,
    trigger_id TEXT NOT NULL,
    action_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    total INTEGER default 1,
    created_at timestamp with time zone default now()
);

-- +migrate Down
DROP TABLE IF EXISTS auto_triggers;
DROP TABLE IF EXISTS auto_conditions;
DROP TABLE IF EXISTS auto_condition_values;
DROP TABLE IF EXISTS auto_condition_types;
DROP TABLE IF EXISTS auto_condition_type_presets;
DROP TABLE IF EXISTS auto_actions;
DROP TABLE IF EXISTS auto_action_types;
DROP TABLE IF EXISTS auto_embeds;
DROP TABLE IF EXISTS auto_embed_images;
DROP TABLE IF EXISTS auto_embed_videos;
DROP TABLE IF EXISTS auto_embed_footers;
DROP TABLE IF EXISTS auto_action_histories;
