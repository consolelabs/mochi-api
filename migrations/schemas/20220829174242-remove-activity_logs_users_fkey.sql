
-- +migrate Up
ALTER TABLE guild_user_activity_logs DROP CONSTRAINT IF EXISTS guild_user_activity_xps_user_id_fkey;

-- +migrate Down
ALTER TABLE guild_user_activity_logs ADD CONSTRAINT guild_user_activity_xps_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id);