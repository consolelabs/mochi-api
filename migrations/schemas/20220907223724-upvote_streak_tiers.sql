
-- +migrate Up
CREATE TABLE IF NOT EXISTS upvote_streak_tiers
(
    id INT NOT NULL,
    streak_required INT NOT NULL,
    xp_per_interval INT NOT NULL,
    vote_interval INT NOT NULL,
    CONSTRAINT upvote_streak_tiers_pkey PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE IF EXISTS upvote_streak_tiers;