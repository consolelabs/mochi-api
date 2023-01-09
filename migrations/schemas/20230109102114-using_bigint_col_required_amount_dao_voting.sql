
-- +migrate Up
ALTER TABLE guild_config_dao_proposal
    ALTER COLUMN required_amount TYPE NUMERIC;

ALTER TABLE dao_proposal_vote_option
    ALTER COLUMN required_amount TYPE NUMERIC;

-- +migrate Down
ALTER TABLE guild_config_dao_proposal
    ALTER COLUMN required_amount TYPE INTEGER;

ALTER TABLE dao_proposal_vote_option
    ALTER COLUMN required_amount TYPE INTEGER;