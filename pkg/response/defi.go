package response

import (
	"math/big"

	"github.com/defipod/mochi/pkg/model"
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
	ID              string       `json:"id"`
	Name            string       `json:"name"`
	Symbol          string       `json:"symbol"`
	MarketCapRank   int          `json:"market_cap_rank"`
	AssetPlatformID string       `json:"asset_platform_id"`
	Image           CoinImage    `json:"image"`
	MarketData      MarketData   `json:"market_data"`
	Tickers         []TickerData `json:"tickers"`
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

type HistoricalTokenPricesResponse struct {
	Data []HistoricalTokenPrice `json:"data"`
}

type HistoricalTokenPrice struct {
	Name     string `json:"contract_name"`
	Decimals int    `json:"contract_decimals"`
	Symbol   string `json:"contract_ticker_symbol"`
	Address  string `json:"contract_address"`
}

type GetHistoricalMarketChartResponse struct {
	Data *CoinPriceHistoryResponse `json:"data"`
}

type InDiscordWalletTransferResponseWrapper struct {
	Data   []InDiscordWalletTransferResponse `json:"data"`
	Errors []string                          `json:"errors"`
}

type InDiscordWalletBalancesResponse struct {
	Status string                `json:"status"`
	Data   *UserBalancesResponse `json:"data"`
}

type GetSupportedTokensResponse struct {
	Data []model.Token `json:"data"`
}

type GetCoinResponseWrapper struct {
	Data *GetCoinResponse `json:"data"`
}

type SearchCoinsResponse struct {
	Data []SearchedCoin `json:"data"`
}

type CompareTokenReponseData struct {
	BaseCoin              *GetCoinResponse `json:"base_coin"`
	TargetCoin            *GetCoinResponse `json:"target_coin"`
	Ratios                []float32        `json:"ratios"`
	Times                 []string         `json:"times"`
	BaseCoinSuggestions   []SearchedCoin   `json:"base_coin_suggestions"`
	TargetCoinSuggestions []SearchedCoin   `json:"target_coin_suggestions"`
	From                  string           `json:"from"`
	To                    string           `json:"to"`
}

type CompareTokenResponse struct {
	Data *CompareTokenReponseData `json:"data"`
}

type AddToWatchlistResponseData struct {
	Suggestions []SearchedCoin `json:"suggestions"`
}

type AddToWatchlistResponse struct {
	Data *AddToWatchlistResponseData `json:"data"`
}

type CoinMarketItemData struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	CurrentPrice  float64 `json:"current_price"`
	Image         string  `json:"image"`
	SparkLineIn7d struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage24h          float64 `json:"price_change_percentage_24h"`
	PriceChangePercentage7dInCurrency float64 `json:"price_change_percentage_7d_in_currency"`
}

type GetWatchlistResponse struct {
	Data []CoinMarketItemData `json:"data"`
}
