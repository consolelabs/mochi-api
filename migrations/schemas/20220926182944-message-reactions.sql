
-- +migrate Up
CREATE TABLE IF NOT EXISTS message_reactions (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	message_id TEXT NOT NULL,
	guild_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
    reaction TEXT NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS message_reactions;
