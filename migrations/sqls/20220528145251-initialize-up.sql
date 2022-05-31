/* Replace with your SQL commands */

-- +migrate Up
CREATE TABLE IF NOT EXISTS activities (
	id SERIAL NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	xp INTEGER NOT NULL DEFAULT 0,
	guild_default BOOLEAN
);

CREATE TABLE IF NOT EXISTS guild_config_activities (
	guild_id TEXT NOT NULL,
	activity_id INTEGER NOT NULL,
	active BOOLEAN,
	CONSTRAINT guild_config_activities_pkey PRIMARY KEY (guild_id, activity_id),
	CONSTRAINT guild_config_activities_guild_id_fkey FOREIGN KEY(guild_id) REFERENCES discord_guilds(id),
	CONSTRAINT guild_config_activities_activity_id_fkey FOREIGN KEY(activity_id) REFERENCES activities(id)
);

CREATE TABLE IF NOT EXISTS config_xp_levels (
	level SERIAL NOT NULL PRIMARY KEY,
	min_xp INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS guild_user_activity_xps (
	id SERIAL NOT NULL PRIMARY KEY,
	guild_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	activity_id INTEGER NOT NULL,
	created_at timestamptz DEFAULT now(),
	CONSTRAINT guild_user_activity_xps_guild_id_fkey FOREIGN KEY(guild_id) REFERENCES discord_guilds(id),
	CONSTRAINT guild_user_activity_xps_user_id_fkey FOREIGN KEY(user_id) REFERENCES users(id),
	CONSTRAINT guild_user_activity_xps_activity_id_fkey FOREIGN KEY(activity_id) REFERENCES activities(id)
);

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