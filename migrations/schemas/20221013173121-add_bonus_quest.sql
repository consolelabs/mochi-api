
-- +migrate Up
ALTER TYPE quest_action ADD VALUE 'bonus';

-- +migrate Down
DELETE FROM public.quests WHERE action::TEXT = 'bonus';
ALTER TABLE quests ALTER COLUMN action TYPE TEXT;
ALTER TABLE quests_user_logs ALTER COLUMN action TYPE TEXT;
ALTER TABLE quests_user_list ALTER COLUMN action TYPE TEXT;
DROP TYPE quest_action;
CREATE TYPE quest_action AS ENUM ('gm', 'vote', 'trade', 'gift', 'ticker', 'watchlist');
ALTER TABLE quests ALTER COLUMN action TYPE quest_action USING action::quest_action;
ALTER TABLE quests_user_logs ALTER COLUMN action TYPE quest_action USING action::quest_action;
ALTER TABLE quests_user_list ALTER COLUMN action TYPE quest_action USING action::quest_action;