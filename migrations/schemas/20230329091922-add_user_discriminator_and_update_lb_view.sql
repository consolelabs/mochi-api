
-- +migrate Up
ALTER TABLE "users" ADD COLUMN "discriminator" text;

DROP VIEW IF EXISTS guild_user_xps;
CREATE VIEW guild_user_xps AS
WITH tmp AS (
  SELECT
    guax.guild_id,
    guax.user_id,
    sum(guax.earned_xp) AS total_xp,
    count(*) AS nr_of_actions,
    rank() OVER (
      PARTITION BY guax.guild_id
      ORDER BY
        sum(guax.earned_xp) DESC
    ) AS guild_rank
  FROM
    guild_user_activity_logs guax
  GROUP BY
    guax.guild_id,
    guax.user_id
)
SELECT
  tmp.guild_id,
  tmp.user_id,
  users.username,
  guild_users.nickname,
  tmp.total_xp,
  tmp.nr_of_actions,
  tmp.guild_rank,
  (
    SELECT
      l.level
    FROM
      config_xp_levels l
    WHERE
      tmp.total_xp >= l.min_xp
    ORDER BY
      l.level DESC
    LIMIT
      1
  ) AS level
FROM
  tmp
  JOIN users ON tmp.user_id = users.id
  JOIN guild_users ON guild_users.user_id = tmp.user_id AND guild_users.guild_id = tmp.guild_id;

-- +migrate Down
ALTER TABLE "users" DROP COLUMN "discriminator";

DROP VIEW IF EXISTS guild_user_xps;
CREATE VIEW guild_user_xps AS WITH tmp(
  guild_id, user_id, total_xp, nr_of_actions,
  guild_rank
) AS (
  SELECT
    guax.guild_id,
    guax.user_id,
    sum(guax.earned_xp) AS sum,
    count(*) AS count,
    rank() OVER (
      PARTITION BY guax.guild_id
      ORDER BY
        (
          sum(guax.earned_xp)
        ) DESC
    ) AS guild_rank
  FROM
    guild_user_activity_logs guax
  GROUP BY
    guax.guild_id,
    guax.user_id
)
SELECT
  tmp.guild_id,
  tmp.user_id,
  users.username,
  tmp.total_xp,
  tmp.nr_of_actions,
  tmp.guild_rank,
  (
    SELECT
      l.level
    FROM
      config_xp_levels l
    WHERE
      tmp.total_xp >= l.min_xp
    ORDER BY
      l.level DESC
    LIMIT
      1
  ) AS level
FROM
  tmp
  JOIN users ON tmp.user_id = users.id;
