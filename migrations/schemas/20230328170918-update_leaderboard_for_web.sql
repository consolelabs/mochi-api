
-- +migrate Up
ALTER TABLE "public"."guild_users" ADD COLUMN "avatar" TEXT;
ALTER TABLE "public"."guild_users" ADD COLUMN "joined_at" TIMESTAMP;
ALTER TABLE "public"."guild_users" ADD COLUMN "roles" JSONB default '[]'::JSONB;

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

-- +migrate Down
ALTER TABLE "public"."guild_users" DROP COLUMN "avatar";
ALTER TABLE "public"."guild_users" DROP COLUMN "joined_at";
ALTER TABLE "public"."guild_users" DROP COLUMN "roles";

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
  tmp;
