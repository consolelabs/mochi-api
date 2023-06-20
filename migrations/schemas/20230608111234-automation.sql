
-- +migrate Up
CREATE TABLE IF NOT EXISTS auto_triggers
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    guild_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_conditions
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    trigger_id TEXT NOT NULL,
    type_id TEXT NOT NULL,
    channel_id TEXT NULL,
    index INTEGER NOT NULL,
    platform TEXT NOT NULL,
    child_id TEXT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_condition_values
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    condition_id TEXT NOT NULL,
    child_id TEXT NULL,
    type TEXT NOT NULL,
    index INTEGER NOT NULL,
    operator TEXT NOT NULL,
    matches TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_condition_types
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_condition_type_presets
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    type_id TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_actions
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    user_ids TEXT NULL,
    trigger_id TEXT NOT NULL,
    type_id TEXT NOT NULL,
    channel_ids TEXT NULL,
    index INTEGER NOT NULL,
    action_data TEXT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NULL,
    embed_id TEXT NULL,
    then_action_id TEXT NULL,
    limit_per_user INTEGER default 1,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_action_types
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embeds
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    action_id TEXT NOT NULL,
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
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    proxy_url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embed_videos
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    embed_id TEXT NOT NULL,
    url TEXT NOT NULL,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auto_embed_footers
(
<<<<<<< HEAD
<<<<<<< HEAD
    id uuid PRIMARY KEY,
=======
    id uuid DEFAULT uuid_generate_v4() primary key,
>>>>>>> 95675554... chore: seed data for demo auto trigger
=======
    id SERIAL PRIMARY KEY,
>>>>>>> 929f0a93... fix: migrate script error
    embed_id TEXT NOT NULL,
    text TEXT NOT NULL,
    icon_url TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
<<<<<<< HEAD
);

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS auto_action_histories
(
    id uuid PRIMARY KEY,
    user_id TEXT NOT NULL,
    trigger_id TEXT NOT NULL,
    action_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    total INTEGER default 1,
    created_at TIMESTAMP NOT NULL DEFAULT now()
=======
>>>>>>> 929f0a93... fix: migrate script error
);

=======
>>>>>>> 95675554... chore: seed data for demo auto trigger
CREATE TABLE IF NOT EXISTS auto_action_histories
(
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    trigger_id TEXT NOT NULL,
    action_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    total INTEGER default 1,
    created_at TIMESTAMP NOT NULL DEFAULT now()
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
