-- +migrate Up
DROP VIEW IF EXISTS guild_user_xps;

ALTER TABLE
	guild_user_activity_xps DROP COLUMN activity_id CASCADE,
ADD
	COLUMN activity_name TEXT,
ADD
	COLUMN earned_xp INTEGER;

ALTER TABLE
	guild_user_activity_xps RENAME TO guild_user_activity_logs;

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

-- +migrate Down
ALTER TABLE
	guild_user_activity_logs RENAME TO guild_user_activity_xps;

ALTER TABLE
	guild_user_activity_xps DROP COLUMN activity_name,
	DROP COLUMN earned_xp
ADD
	COLUMN activity_id INTEGER NOT NULL,
ADD
	CONSTRAINT guild_user_activity_xps_activity_id_fkey FOREIGN KEY(activity_id) REFERENCES activities(id);

DROP VIEW IF EXISTS guild_user_xps;

CREATE VIEW guild_user_xps AS WITH tmp(
	guild_id,
	user_id,
	total_xp,
	nr_of_actions,
	guild_rank
) AS (
	SELECT
		guax.guild_id,
		guax.user_id,
		SUM(a.xp),
		COUNT(guax),
		RANK() OVER (
			ORDER BY
				SUM(a.xp) DESC
		) guild_rank
	FROM
		guild_user_activity_xps guax
		JOIN activities a ON guax.activity_id = a.id
	GROUP BY
		guax.guild_id,
		guax.user_id
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