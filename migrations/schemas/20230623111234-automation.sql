
-- +migrate Up
CREATE TABLE IF NOT EXISTS auto_triggers
(
    id serial PRIMARY KEY,
    discord_guild_id TEXT NOT NULL,
    user_discord_id TEXT NOT NULL,
    name TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_types
(
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_conditions
(
    id serial PRIMARY KEY,
    trigger_id INTEGER NOT NULL REFERENCES auto_triggers(id),
    type_id INTEGER NOT NULL REFERENCES auto_types(id),
    channel_id TEXT NULL,
    index INTEGER NOT NULL,
    platform TEXT NOT NULL,
    child_id TEXT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_condition_values
(
    id serial PRIMARY KEY,
    condition_id INTEGER NOT NULL REFERENCES auto_conditions(id),
    child_id TEXT NULL,
    type_id INTEGER NOT NULL REFERENCES auto_types(id),
    index INTEGER NOT NULL,
    operator TEXT NOT NULL,
    matches TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS auto_condition_type_presets
(
    id serial PRIMARY KEY,
    type_id INTEGER NOT NULL REFERENCES auto_types(id),
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_actions
(
    id serial PRIMARY KEY,
    user_ids TEXT NULL,
    trigger_id INTEGER NOT NULL REFERENCES auto_triggers(id),
    type_id INTEGER NOT NULL REFERENCES auto_types(id),
    channel_ids TEXT NULL,
    index INTEGER NOT NULL,
    action_data TEXT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NULL,
    then_action_id INTEGER NULL,
    limit_per_user INTEGER default 1,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embeds
(
    id serial PRIMARY KEY,
    action_id INTEGER NOT NULL,
    author_id TEXT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    color TEXT NULL,
    thumbnail TEXT NULL,
    url TEXT NULL,
    type TEXT NULL,
    fields TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embed_images
(
    id serial PRIMARY KEY,
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    proxy_url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embed_videos
(
    id serial PRIMARY KEY,
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embed_footers
(
    id serial PRIMARY KEY,
    embed_id TEXT NOT NULL,
    text TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_action_histories
(
    id serial PRIMARY KEY,
    user_discord_id TEXT NOT NULL,
    trigger_id INTEGER NOT NULL REFERENCES auto_triggers(id),
    action_id INTEGER NOT NULL REFERENCES auto_actions(id),
    message_id TEXT NOT NULL,
    total INTEGER default 1,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS auto_embeds;
DROP TABLE IF EXISTS auto_embed_images;
DROP TABLE IF EXISTS auto_embed_videos;
DROP TABLE IF EXISTS auto_embed_footers;
DROP TABLE IF EXISTS auto_action_histories;
DROP TABLE IF EXISTS auto_actions;
DROP TABLE IF EXISTS auto_condition_type_presets;
DROP TABLE IF EXISTS auto_condition_values;
DROP TABLE IF EXISTS auto_conditions;
DROP TABLE IF EXISTS auto_types;
DROP TABLE IF EXISTS auto_triggers;
