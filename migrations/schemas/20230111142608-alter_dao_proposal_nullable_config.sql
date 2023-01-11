
-- +migrate Up
ALTER TABLE dao_proposal DROP CONSTRAINT dao_proposal_guild_config_dao_proposal_id_fkey;

-- +migrate Down
ALTER TABLE dao_proposal ADD CONSTRAINT dao_proposal_guild_config_dao_proposal_id_fkey FOREIGN KEY (guild_config_dao_proposal_id) REFERENCES guild_config_dao_proposal(id);