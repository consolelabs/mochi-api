-- +migrate Up
CREATE TABLE IF NOT EXISTS token_infos(
  source text NOT NULL,
  token text NOT NULL,
  data jsonb NOT NULL DEFAULT '{}' ::jsonb
);

CREATE UNIQUE INDEX IF NOT EXISTS token_infos_source_token_idx ON token_infos(source, token);

-- +migrate Down
DROP TABLE IF EXISTS token_infos;

