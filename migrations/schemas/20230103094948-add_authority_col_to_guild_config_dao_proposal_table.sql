
-- +migrate Up
CREATE TYPE proposal_authority AS enum (
    'admin',
    'token_holder'
);

CREATE TABLE IF NOT EXISTS dao_guideline_messages  (
    id SERIAL NOT NULL PRIMARY KEY,
    authority proposal_authority NOT NULL,
    message TEXT
);
CREATE UNIQUE INDEX authority_unique_idx on dao_guideline_messages (authority);

ALTER TABLE guild_config_dao_proposal
	ADD COLUMN authority proposal_authority NOT NULL;

-- +migrate Down
ALTER TABLE guild_config_dao_proposal
	DROP COLUMN authority;

DROP TABLE IF EXISTS dao_guideline_messages;

DROP TYPE IF EXISTS proposal_authority;

