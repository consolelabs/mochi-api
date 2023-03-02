
-- +migrate Up
CREATE UNIQUE INDEX user_token_price_alerts_id_uindex ON user_token_price_alerts (id);

-- +migrate Down
DROP INDEX user_token_price_alerts_id_uindex;
