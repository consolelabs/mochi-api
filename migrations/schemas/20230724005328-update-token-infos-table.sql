-- +migrate Up
DROP INDEX IF EXISTS token_infos_source_token_idx;

ALTER TABLE token_infos
  ADD PRIMARY KEY (token);

ALTER TABLE token_infos
  DROP COLUMN source;

-- +migrate Down
ALTER TABLE token_infos
  DROP CONSTRAINT token_infos_pkey;

CREATE UNIQUE INDEX IF NOT EXISTS token_infos_source_token_idx ON token_infos(source, token);

ALTER TABLE token_infos
  ADD COLUMN source text NOT NULL DEFAULT '' ::text;

