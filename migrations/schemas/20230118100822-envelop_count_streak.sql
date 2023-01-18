
-- +migrate Up
CREATE TABLE IF NOT EXISTS envelops (
	id SERIAL NOT NULL PRIMARY KEY,
	user_id TEXT NOT NULL,
	command TEXT NOT NULL ,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE VIEW user_envelops AS (
	SELECT
		envelops.user_id,
		count(envelops.user_id) AS total_envelop
	FROM
		envelops
	GROUP BY
		envelops.user_id
	ORDER BY
		total_envelop DESC
);


-- +migrate Down
DROP VIEW IF EXISTS user_envelops;
DROP TABLE IF EXISTS envelops;