
-- +migrate Up
ALTER TABLE guild_user_activity_logs ALTER COLUMN guild_id DROP NOT NULL;

-- +migrate Down
ALTER TABLE guild_user_activity_logs ALTER COLUMN guild_id SET NOT NULL;