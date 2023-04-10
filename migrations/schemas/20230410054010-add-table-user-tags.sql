
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_tags
(
    id serial primary key,
    guild_id text,
    user_id text NOT NULL,
    mention_username bool DEFAULT FALSE,
    mention_role bool DEFAULT FALSE,
    updated_at timestamptz DEFAULT now(),
    created_at timestamptz DEFAULT now()
    );
CREATE UNIQUE INDEX user_guild_user_id_user_tag_uidx ON user_tags (user_id, guild_id);
-- +migrate Down
DROP TABLE IF EXISTS user_tags;
