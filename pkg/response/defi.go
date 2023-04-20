package response

import (
	"math/big"
	"time"

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
	Coins []CoingeckoSearchedCoin `json:"coins"`
}

type CoingeckoSearchedCoin struct {
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
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Symbol          string          `json:"symbol"`
	MarketCapRank   int             `json:"market_cap_rank"`
	AssetPlatformID string          `json:"asset_platform_id"`
	Image           CoinImage       `json:"image"`
	MarketData      MarketData      `json:"market_data"`
	Tickers         []TickerData    `json:"tickers"`
	Description     CoinDescription `json:"description"`
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
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type CoinDescription struct {
	EngDescription string `json:"en"`
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

type GetSupportedTokenResponse struct {
	Data *model.Token `json:"data"`
}

type GetCoinResponseWrapper struct {
	Data *GetCoinResponse `json:"data"`
}

type CompareTokenReponseData struct {
	BaseCoin              *GetCoinResponse                 `json:"base_coin"`
	TargetCoin            *GetCoinResponse                 `json:"target_coin"`
	Ratios                []float64                        `json:"ratios"`
	Times                 []string                         `json:"times"`
	BaseCoinSuggestions   []model.CoingeckoSupportedTokens `json:"base_coin_suggestions"`
	TargetCoinSuggestions []model.CoingeckoSupportedTokens `json:"target_coin_suggestions"`
	From                  string                           `json:"from"`
	To                    string                           `json:"to"`
}

type CompareTokenResponse struct {
	Data *CompareTokenReponseData `json:"data"`
}

type AddToWatchlistResponseData struct {
	BaseSuggestions   []model.CoingeckoSupportedTokens `json:"base_suggestions"`
	TargetSuggestions []model.CoingeckoSupportedTokens `json:"target_suggestions"`
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
	MarketCap     float64 `json:"market_cap"`
	MarketCapRank int64   `json:"market_cap_rank"`
	SparkLineIn7d struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage24h           float64 `json:"price_change_percentage_24h"`
	PriceChangePercentage7dInCurrency  float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage1hInCurrency  float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency float64 `json:"price_change_percentage_24h_in_currency"`
	IsPair                             bool    `json:"is_pair"`
	IsDefault                          bool    `json:"is_default"`
}

type GetWatchlistResponse struct {
	Pagination *PaginationResponse  `json:"metadata"`
	Data       []CoinMarketItemData `json:"data"`
}

type CoingeckoSupportedTokenResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type SearchCoinResponse struct {
	Data []model.CoingeckoSupportedTokens `json:"data"`
}

type GetFiatHistoricalExchangeRatesResponse struct {
	LatestRate float64   `json:"latest_rate"`
	Rates      []float64 `json:"rates"`
	Times      []string  `json:"times"`
	From       string    `json:"from"`
	To         string    `json:"to"`
}

type TokenPriceAlertResponseData struct {
	UserDiscordID  string    `json:"user_discord_id"`
	Symbol         string    `json:"symbol"`
	Currency       string    `json:"currency"`
	AlertType      string    `json:"alert_type"`
	Frequency      string    `json:"frequency"`
	Value          float64   `json:"value"`
	PriceByPercent float64   `json:"price_by_percent"`
	SnoozedTo      time.Time `json:"snoozed_to"`
}

type AddTokenPriceAlertResponse struct {
	Data *TokenPriceAlertResponseData `json:"data"`
}

type ListTokenPriceAlertResponse struct {
	UserDiscordID string               `json:"user_discord_id"`
	Symbol        string               `json:"symbol"`
	Currency      string               `json:"currency"`
	AlertType     model.AlertType      `json:"alert_type"`
	Frequency     model.AlertFrequency `json:"frequency"`
	Price         float64              `json:"price"`
	SnoozedTo     time.Time            `json:"snoozed_to"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

type CreateUserTokenSupportRequest struct {
	Data *model.UserTokenSupportRequest `json:"data"`
}

type GasTrackerResponseData struct {
	Data []GasTrackerResponse `json:"data"`
}

type ChainGasTrackerResponseData struct {
	Data *GasTrackerResponse `json:"data"`
}

type GasTrackerResponse struct {
	Chain           string `json:"chain"`
	SafeGasPrice    string `json:"safe_gas_price"`
	ProposeGasPrice string `json:"propose_gas_price"`
	FastGasPrice    string `json:"fast_gas_price"`
	EstSafeTime     string `json:"est_safe_time"`
	EstProposeTime  string `json:"est_propose_time"`
	EstFastTime     string `json:"est_fast_time"`
}

type GetCoinsMarketDataResponse struct {
	Data []CoinMarketItemData `json:"data"`
}

type AssetPlatformResponseData struct {
	ChainIdentifier *int64 `json:"chain_identifier"`
	ID              string `json:"id"`
}

type GetCoinByContractResponseData struct {
	ID              string                            `json:"id"`
	Name            string                            `json:"name"`
	Symbol          string                            `json:"symbol"`
	DetailPlatforms map[string]CoinPlatformDetailData `json:"detail_platforms"`
	Image           CoinImage                         `json:"image"`
}

type CoinPlatformDetailData struct {
	DecimalPlace    int    `json:"decimal_place"`
	ContractAddress string `json:"contract_address"`
}

type GetTrendingSearch struct {
	Coins     []GetTrendingSearchCoin `json:"coins"`
	Exchanges interface{}             `json:"exchanges"` // this field coingecko return empty
}

type GetTrendingSearchCoin struct {
	Item Coin `json:"item"`
}

type Coin struct {
	Id            string  `json:"id"`
	CoinId        int64   `json:"coin_id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	MarketCapRank int64   `json:"market_cap_rank"`
	Thumb         string  `json:"thumb"`
	Small         string  `json:"small"`
	Large         string  `json:"large"`
	Slug          string  `json:"slug"`
	PriceBtc      float64 `json:"price_btc"`
	Score         int64   `json:"score"`
}
