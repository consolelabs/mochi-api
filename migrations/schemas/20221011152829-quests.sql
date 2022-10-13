
-- +migrate Up
CREATE TYPE quest_routine AS ENUM ('daily', 'weekly', 'monthly', 'yearly', 'once');

CREATE TYPE quest_action AS ENUM ('gm', 'vote', 'trade', 'gift', 'ticker', 'watchlist');

CREATE TABLE IF NOT EXISTS quests (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	title TEXT NOT NULL,
	action quest_action NOT NULL,
	frequency INT NOT NULL,
	routine quest_routine NOT NULL DEFAULT 'daily'::quest_routine
);

INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Say GM/GN', 'gm', 1, 'daily');
INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Vote for Mochi Bot on topgg/discordbotlist', 'vote', 2, 'daily');
INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Trading NFT', 'trade', 1, 'daily');
INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Gift XP for other people', 'gift', 1, 'daily');
INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Check any token price 3 times with $ticker', 'ticker', 3, 'daily');
INSERT INTO public.quests (title, action, frequency, routine) VALUES ('Check your watchlist 3 times', 'watchlist', 3, 'daily');

CREATE TABLE IF NOT EXISTS quests_reward_types (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL
);

INSERT INTO public.quests_reward_types (name) VALUES ('xp');
INSERT INTO public.quests_reward_types (name) VALUES ('coin');

CREATE TABLE IF NOT EXISTS quests_user_logs (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	guild_id TEXT,
	user_id TEXT NOT NULL,
	quest_id UUID NOT NULL,
	action quest_action NOT NULL,
	target INT NOT NULL DEFAULT 0,
	created_at timestamptz DEFAULT now(),
	FOREIGN KEY (guild_id) REFERENCES discord_guilds(id)
);

CREATE TABLE IF NOT EXISTS quests_user_list (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id TEXT NOT NULL,
	quest_id UUID NOT NULL,
	action quest_action NOT NULL,
	routine quest_routine NOT NULL DEFAULT 'daily'::quest_routine,
	current INT NOT NULL DEFAULT 0,
	target INT NOT NULL DEFAULT 0,
	is_completed BOOLEAN NOT NULL DEFAULT FALSE,
	is_claimed BOOLEAN NOT NULL DEFAULT FALSE,
	start_time timestamptz DEFAULT now(),
	end_time timestamptz DEFAULT now(),
	FOREIGN KEY (quest_id) REFERENCES quests(id),
	UNIQUE (user_id, quest_id, start_time)
);

CREATE TABLE IF NOT EXISTS quests_passes (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS quests_user_pass (
	user_id TEXT NOT NULL,
	pass_id UUID NOT NULL,
	created_at timestamptz DEFAULT now(),
	expired_at timestamptz DEFAULT now(),
	active BOOLEAN NOT NULL DEFAULT FALSE,
	FOREIGN KEY (pass_id) REFERENCES quests_passes(id)
);

CREATE TABLE IF NOT EXISTS quests_rewards (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	quest_id UUID NOT NULL,
	reward_type_id UUID NOT NULL,
	reward_amount INT NOT NULL DEFAULT 0,
	pass_id UUID DEFAULT NULL,
	FOREIGN KEY (quest_id) REFERENCES quests(id),
	FOREIGN KEY (reward_type_id) REFERENCES quests_reward_types(id),
	FOREIGN KEY (pass_id) REFERENCES quests_passes(id),
	UNIQUE (quest_id, reward_type_id, pass_id)
);

CREATE TABLE IF NOT EXISTS quests_user_rewards (
	user_id TEXT NOT NULL,
	quest_id UUID NOT NULL,
	reward_id UUID NOT NULL,
	reward_type_id UUID NOT NULL,
	reward_amount INT NOT NULL DEFAULT 0,
	pass_id UUID DEFAULT NULL,
	start_time timestamptz,
	claimed_at timestamptz DEFAULT now(),
	FOREIGN KEY (quest_id) REFERENCES quests(id),
	FOREIGN KEY (reward_id) REFERENCES quests_rewards(id),
	FOREIGN KEY (reward_type_id) REFERENCES quests_reward_types(id),
	FOREIGN KEY (pass_id) REFERENCES quests_passes(id),
	UNIQUE (user_id, quest_id, reward_id, pass_id, start_time)
);

-- +migrate Down
DROP TABLE IF EXISTS quests_user_rewards;
DROP TABLE IF EXISTS quests_rewards;
DROP TABLE IF EXISTS quests_user_pass;
DROP TABLE IF EXISTS quests_passes;
DROP TABLE IF EXISTS quests_user_list;
DROP TABLE IF EXISTS quests_user_logs;
DROP TABLE IF EXISTS quests_reward_types;
DROP TABLE IF EXISTS quests;
DROP TYPE IF EXISTS quest_action;
DROP TYPE IF EXISTS quest_routine;
