
-- +migrate Up
ALTER TABLE twitter_posts ADD COLUMN IF NOT EXISTS content TEXT;
ALTER TABLE twitter_posts ADD COLUMN IF NOT EXISTS created_at timestamp DEFAULT now();

CREATE TABLE IF NOT EXISTS twitter_post_streaks (
	guild_id TEXT NOT NULL,
	twitter_id TEXT NOT NULL,
	twitter_handle TEXT NOT NULL,
	streak_count INT NOT NULL DEFAULT 1,
	total_count INT NOT NULL DEFAULT 1,
	last_streak_date TIMESTAMP NOT NULL,
  created_at       TIMESTAMP DEFAULT now(),
  updated_at       TIMESTAMP DEFAULT now(),
	FOREIGN KEY (guild_id) REFERENCES discord_guilds(id),
	UNIQUE (guild_id, twitter_id),
	UNIQUE (guild_id, twitter_handle)
);

-- +migrate Down
DROP TABLE IF EXISTS twitter_post_streaks;
ALTER TABLE twitter_posts DROP COLUMN IF EXISTS created_at;
ALTER TABLE twitter_posts DROP COLUMN IF EXISTS content;