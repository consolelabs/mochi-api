/* Replace with your SQL commands */
-- +migrate Down
DROP VIEW IF EXISTS guild_user_xps;
CREATE VIEW guild_users_xps AS WITH tmp(guild_id, user_id, total_xp) AS (
	SELECT
		guax.guild_id,
		guax.user_id,
		SUM(a.xp)
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