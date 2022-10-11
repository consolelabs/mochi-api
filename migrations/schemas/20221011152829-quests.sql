
-- +migrate Up
CREATE TYPE quest_routine AS ENUM ('daily', 'weekly', 'monthly', 'yearly', 'once');

CREATE TYPE quest_reward_type AS ENUM ('xp', 'coin');

CREATE TYPE quest_action AS ENUM ('gm', 'vote', 'trade', 'gift', 'ticker', 'watchlist');

CREATE TABLE IF NOT EXISTS quests (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	title TEXT NOT NULL,
	action quest_action NOT NULL,
	frequency INT NOT NULL,
	bonus_frequency INT NOT NULL,
	routine quest_routine NOT NULL DEFAULT 'daily'::quest_routine
);

CREATE TABLE IF NOT EXISTS quests_reward_types (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name quest_reward_type NOT NULL DEFAULT 'xp'::quest_reward_type
);

CREATE TABLE IF NOT EXISTS quests_user_logs (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	guild_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	quest_id UUID NOT NULL,
	action quest_action NOT NULL,
	reward_type_id UUID NOT NULL,
	target INT NOT NULL DEFAULT 0,
	created_at timestamptz DEFAULT now(),
	FOREIGN KEY (guild_id) REFERENCES discord_guilds(id),
	FOREIGN KEY (reward_type_id) REFERENCES quests_reward_types(id)
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
	end_time timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS quests_passes (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS quests_rewards (
	quest_id UUID NOT NULL,
	reward_type_id UUID NOT NULL,
	amount INT NOT NULL DEFAULT 0,
	bonus_amount INT NOT NULL DEFAULT 0,
	pass_id UUID DEFAULT NULL,
	FOREIGN KEY (pass_id) REFERENCES quests_passes(id),
	UNIQUE (quest_id, reward_type_id, pass_id)
);

CREATE TABLE IF NOT EXISTS quests_user_pass (
	user_id TEXT NOT NULL,
	pass_id UUID NOT NULL,
	created_at timestamptz DEFAULT now(),
	expired_at timestamptz DEFAULT now(),
	active BOOLEAN NOT NULL DEFAULT FALSE
);

-- +migrate Down
DROP TABLE IF EXISTS quests_rewards;
DROP TABLE IF EXISTS quests_passes;
DROP TABLE IF EXISTS quests_user_list;
DROP TABLE IF EXISTS quests_user_logs;
DROP TABLE IF EXISTS quests_reward_types;
DROP TABLE IF EXISTS quests;
DROP TYPE IF EXISTS quest_action;
DROP TYPE IF EXISTS quest_reward_type;
DROP TYPE IF EXISTS quest_routine;
