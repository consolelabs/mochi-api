
-- +migrate Up
ALTER TABLE guild_users ALTER COLUMN invited_by TYPE
TEXT;

ALTER TABLE invite_histories ADD COLUMN type TEXT DEFAULT 'normal';

-- +migrate Down
ALTER TABLE guild_users ALTER COLUMN invited_by TYPE
BIGINT USING coalesce
(cast
(nullif
(invited_by,'') as BIGINT),0);

ALTER TABLE invite_histories DROP COLUMN type;