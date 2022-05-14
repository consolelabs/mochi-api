
-- +migrate Up
DROP VIEW IF EXISTS guild_user_xps;
CREATE VIEW guild_user_xps AS WITH tmp(
	guild_id,
	user_id,
	total_xp,
	nr_of_actions,
	guild_rank
) AS (
	SELECT
		guild_id,
		user_id,
		SUM(earned_xp),
		COUNT(*),
		RANK() OVER (
			PARTITION BY guild_id
			ORDER BY
				SUM(earned_xp) DESC
		) guild_rank
	FROM
		guild_user_activity_logs guax
	GROUP BY
		guild_id,
		user_id
)
SELECT
	guild_id,
	user_id,
	total_xp,
	nr_of_actions,
	guild_rank,
	(
		SELECT
			l.level
		from
			config_xp_levels l
		WHERE
			total_xp >= l.min_xp
		ORDER BY
			l.level DESC
		LIMIT
			1
	)
FROM
	tmp;

-- +migrate Down
DROP VIEW IF EXISTS guild_user_xps;
CREATE VIEW guild_user_xps AS WITH tmp(
	guild_id,
	user_id,
	total_xp,
	nr_of_actions,
	guild_rank
) AS (
	SELECT
		guild_id,
		user_id,
		SUM(earned_xp),
		COUNT(*),
		RANK() OVER (
			ORDER BY
				SUM(earned_xp) DESC
		) guild_rank
	FROM
		guild_user_activity_logs guax
	GROUP BY
		guild_id,
		user_id
)
SELECT
	guild_id,
	user_id,
	total_xp,
	nr_of_actions,
	guild_rank,
	(
		SELECT
			l.level
		from
			config_xp_levels l
		WHERE
			total_xp >= l.min_xp
		ORDER BY
			l.level DESC
		LIMIT
			1
	)
FROM
	tmp;
