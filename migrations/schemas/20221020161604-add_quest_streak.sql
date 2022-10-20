
-- +migrate Up
CREATE TABLE IF NOT EXISTS quests_streak (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	title TEXT NOT NULL,
	action quest_action NOT NULL,
	streak_from INT NOT NULL,
	streak_to INT,
	multiplier DOUBLE PRECISION NOT NULL,
	UNIQUE (action, streak_from)
);

ALTER TABLE quests_user_list ADD COLUMN IF NOT EXISTS multiplier DOUBLE PRECISION NOT NULL DEFAULT 1;

INSERT INTO public.quests_streak (title, action, streak_from, streak_to, multiplier) VALUES ('Upvoting streak for Mochi Bot', 'vote', 3, 10, 1.2);
INSERT INTO public.quests_streak (title, action, streak_from, streak_to, multiplier) VALUES ('Upvoting streak for Mochi Bot', 'vote', 11, 30, 1.5);
INSERT INTO public.quests_streak (title, action, streak_from, streak_to, multiplier) VALUES ('Upvoting streak for Mochi Bot', 'vote', 31, 60, 2);
INSERT INTO public.quests_streak (title, action, streak_from, streak_to, multiplier) VALUES ('Upvoting streak for Mochi Bot', 'vote', 61, 100, 2.5);

-- +migrate Down
ALTER TABLE quests_user_list DROP COLUMN IF EXISTS multiplier;
DROP TABLE IF EXISTS quests_streak;