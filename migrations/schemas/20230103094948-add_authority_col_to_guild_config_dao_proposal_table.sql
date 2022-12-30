
-- +migrate Up
CREATE TYPE proposal_authority AS enum (
    'admin',
    'token_holder'
);

ALTER TABLE guild_config_dao_proposal
	ADD COLUMN authority proposal_authority NOT NULL;

-- +migrate Down
DROP enum IF EXISTS proposal_authority;
ALTER TABLE guild_config_dao_proposal
	DROP COLUMN authority;

