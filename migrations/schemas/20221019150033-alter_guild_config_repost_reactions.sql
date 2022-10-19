
-- +migrate Up
DROP INDEX guild_config_repost_reactions_emoji_uindex;
CREATE UNIQUE INDEX guild_config_repost_reactions_emoji_guild_id_uindex on guild_config_repost_reactions (guild_id, emoji, emoji_start, emoji_stop);
-- +migrate Down
DROP INDEX guild_config_repost_reactions_emoji_guild_id_uindex;
CREATE UNIQUE INDEX guild_config_repost_reactions_emoji_uindex ON guild_config_repost_reactions (emoji);
