
-- +migrate Up
CREATE TYPE feedback_status AS ENUM ('none', 'confirmed', 'completed');

CREATE TABLE IF NOT EXISTS user_feedbacks (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	discord_id TEXT NOT NULL,
    command TEXT NOT NULL,
	feedback TEXT NOT NULL,
    status feedback_status DEFAULT 'none',
	created_at timestamptz NOT NULL DEFAULT NOW(),
	confirmed_at timestamptz,
	completed_at timestamptz
);

-- +migrate Down
DROP TABLE IF EXISTS user_feedbacks;
DROP TYPE feedback_status;