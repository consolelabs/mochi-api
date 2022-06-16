package response

import (
	"math/big"
)

type InDiscordWalletWithdrawResponse struct {
	FromDiscordID    string     `json:"fromDiscordId"`
	ToAddress        string     `json:"toAddress"`
	Amount           float64    `json:"amount"`
	Cryptocurrency   string     `json:"cryptocurrency"`
	TxHash           string     `json:"txHash"`
	TxURL            string     `json:"txURL"`
	WithdrawalAmount *big.Float `json:"withdrawalAmount"`
	TransactionFee   float64    `json:"transactionFee"`
}

type InDiscordWalletTransferResponse struct {
	FromDiscordID  string  `json:"fromDiscordID"`
	ToDiscordID    string  `json:"toDiscordID"`
	Amount         float64 `json:"amount"`
	Cryptocurrency string  `json:"cryptocurrency"`
	TxHash         string  `json:"txHash"`
	TxUrl          string  `json:"txUrl"`
	TransactionFee float64 `json:"transactionFee"`
}

type UserBalancesResponse struct {
	Balances      map[string]float64 `json:"balances"`
	BalancesInUSD map[string]float64 `json:"balances_in_usd"`
}

type HistoricalMarketChartResponse struct {
	Prices [][]float64 `json:"prices"`
}

type SearchedCoinsListResponse struct {
	Coins []SearchedCoin `json:"coins"`
}

type SearchedCoin struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	MarketCapRank int    `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}
type CoinPriceHistoryResponse struct {
	TokenTickers
	From string `json:"from"`
	To   string `json:"to"`
}

type TokenTickers struct {
	Timestamps []int64   `json:"timestamps"`
	Prices     []float64 `json:"prices"`
	Times      []string  `json:"times"`
}

type GetCoinResponse struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Symbol        string       `json:"symbol"`
	MarketCapRank int          `json:"market_cap_rank"`
	Image         CoinImage    `json:"image"`
	MarketData    MarketData   `json:"market_data"`
	Tickers       []TickerData `json:"tickers"`
}

type TickerData struct {
	Base         string  `json:"base"`
	Target       string  `json:"target"`
	Last         float32 `json:"last"`
	CoinID       string  `json:"coin_id"`
	TargetCoinID string  `json:"target_coin_id"`
}

type MarketData struct {
	CurrentPrice                       map[string]float64 `json:"current_price"`
	MarketCap                          map[string]float64 `json:"market_cap"`
	PriceChangePercentage1hInCurrency  map[string]float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency map[string]float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency  map[string]float64 `json:"price_change_percentage_7d_in_currency"`
}

type CoinImage struct {
	Thumb  string `json:"thumb"`
	Small  string `json:"small"`
	Larget string `json:"large"`
}

type CoinPriceResponse map[string]map[string]float64
