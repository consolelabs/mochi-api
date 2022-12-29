
-- +migrate Up
CREATE VIEW view_dao_proposal AS (
    SELECT
        sum(
            v.point),
        v.choice,
        v.proposal_id,
        p.guild_id
    FROM
        dao_vote AS v
        JOIN dao_proposal AS p ON p.id = v.proposal_id
    GROUP BY
        v.choice,
        v.proposal_id,
        p.guild_id
);

alter table guild_config_dao_proposal add column if not exists type vote_option;
alter table guild_config_dao_proposal add column if not exists required_amount integer;
alter table guild_config_dao_proposal add column if not exists chain_id integer;
alter table guild_config_dao_proposal add column if not exists address text;
alter table guild_config_dao_proposal add column if not exists symbol text;

-- +migrate Down
drop view if exists view_dao_proposal;

alter table guild_config_dao_proposal drop column if exists type;
alter table guild_config_dao_proposaldrop column if exists required_amount;
alter table guild_config_dao_proposal drop column if exists chain_id;
alter table guild_config_dao_proposal drop column if exists address;
alter table guild_config_dao_proposal drop column if exists symbol;