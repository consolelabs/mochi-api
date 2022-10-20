
-- +migrate Up
CREATE TABLE IF NOT EXISTS trade_offers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    from_address TEXT NOT NULL,
    to_address TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS trade_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    is_from BOOLEAN NOT NULL,
    token_address TEXT NOT NULL,
    token_ids JSONB NOT NULL DEFAULT '[]'::JSONB,
    trade_offer_id UUID NOT NULL,
    FOREIGN KEY (trade_offer_id) REFERENCES trade_offers(id) ON DELETE CASCADE   
);


-- +migrate Down
DROP TABLE IF EXISTS trade_items, trade_offers;