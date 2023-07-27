
-- +migrate Up
alter table vault_treasurers add column if not exists user_profile_id text;
alter table vault_requests add column if not exists user_profile_id text;
alter table vault_requests add column if not exists requester_profile_id text;
alter table vault_submissions add column if not exists submitter_profile_id text;

alter table vault_treasurers drop constraint if exists treasurers_vault_id_guild_id_user_discord_id_key;
alter table vault_treasurers add constraint treasurers_vault_id_guild_id_user_profile_id_key unique (vault_id, guild_id, user_profile_id);
-- +migrate Down
alter table vault_treasurers drop column if exists user_profile_id;
alter table vault_requests drop column if exists user_profile_id;
alter table vault_requests drop column if exists requester_profile_id;
alter table vault_submissions drop column if exists submitter_profile_id;
alter table vault_treasurers drop constraint if exists treasurers_vault_id_guild_id_user_profile_id_key;
alter table vault_treasurers add constraint treasurers_vault_id_guild_id_user_discord_id_key unique (vault_id, guild_id, user_discord_id);