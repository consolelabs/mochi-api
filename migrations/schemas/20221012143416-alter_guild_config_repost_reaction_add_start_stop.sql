
-- +migrate Up
ALTER TABLE guild_config_repost_reactions ADD COLUMN IF NOT EXISTS emoji_start text;
ALTER TABLE guild_config_repost_reactions ADD COLUMN IF NOT EXISTS emoji_stop text;

ALTER TABLE message_repost_histories ADD COLUMN IF NOT EXISTS is_start boolean;
ALTER TABLE message_repost_histories ADD COLUMN IF NOT EXISTS is_stop boolean;
ALTER TABLE message_repost_histories ADD COLUMN IF NOT EXISTS reaction_count integer;

-- +migrate Down
ALTER TABLE guild_config_repost_reactions DROP COLUMN IF EXISTS emoji_start;
ALTER TABLE guild_config_repost_reactions DROP COLUMN IF EXISTS emoji_stop;

ALTER TABLE message_repost_histories DROP COLUMN IF EXISTS is_start;
ALTER TABLE message_repost_histories DROP COLUMN IF EXISTS is_stop;
ALTER TABLE message_repost_histories DROP COLUMN IF EXISTS reaction_count;
