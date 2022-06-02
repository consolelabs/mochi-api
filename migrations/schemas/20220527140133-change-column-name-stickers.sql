
-- +migrate Up
ALTER TABLE "discord_guild_stats"
    RENAME COLUMN nr_of_guild_stickers TO nr_of_server_stickers;
    
ALTER TABLE "discord_guild_stats"
    RENAME COLUMN nr_of_standard_stickers TO nr_of_custom_stickers;

-- +migrate Down

ALTER TABLE "discord_guild_stats" 
    RENAME COLUMN nr_of_server_stickers TO nr_of_guild_stickers;

ALTER TABLE "discord_guild_stats" 
    RENAME COLUMN nr_of_custom_stickers TO nr_of_standard_stickers;