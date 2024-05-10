package model

import "time"

type BinanceSpotTransaction struct {
	ID                      int64     `json:"id"`
	ProfileId               string    `json:"profile_id"`
	Pair                    string    `json:"pair"`
	Symbol                  string    `json:"symbol"`
	OrderId                 int64     `json:"order_id"`
	OrderListId             int64     `json:"order_list_id"`
	ClientOrderId           string    `json:"client_order_id"`
	Price                   string    `json:"price"`
	PriceInUsd              string    `json:"price_in_usd"`
	OrigQty                 string    `json:"orig_qty"`
	ExecutedQty             string    `json:"executed_qty"`
	CumulativeQuoteQty      string    `json:"cumulative_quote_qty"`
	Status                  string    `json:"status"`
	TimeInForce             string    `json:"time_in_force"`
	Type                    string    `json:"type"`
	Side                    string    `json:"side"`
	StopPrice               string    `json:"stop_price"`
	IcebergQty              string    `json:"iceberg_qty"`
	Time                    int64     `json:"time"`
	UpdateTime              int64     `json:"update_time"`
	IsWorking               bool      `json:"is_working"`
	OrigQuoteOrderQty       string    `json:"orig_quote_order_qty"`
	WorkingTime             time.Time `json:"working_time"`
	SelfTradePreventionMode string    `json:"self_trade_prevention_mode"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
