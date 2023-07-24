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
	Prices     [][]float64 `json:"prices"`
	MarketCaps [][]float64 `json:"market_caps"`
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
	ID                           string                            `json:"id"`
	Name                         string                            `json:"name"`
	Symbol                       string                            `json:"symbol"`
	AssetPlatformID              string                            `json:"asset_platform_id"`
	AssetPlatform                *AssetPlatformResponseData        `json:"asset_platform"`
	Platforms                    interface{}                       `json:"platforms"`
	BlockTimeInMinutes           int64                             `json:"block_time_in_minutes"`
	HashingAlgorithm             interface{}                       `json:"hashing_algorithm"`
	Categories                   []string                          `json:"categories"`
	Localization                 map[string]string                 `json:"localization"`
	Description                  map[string]string                 `json:"description"`
	Links                        interface{}                       `json:"links"`
	Image                        CoinImage                         `json:"image"`
	GenesisDate                  interface{}                       `json:"genesis_date"`
	SentimentVotesUpPercentage   float64                           `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                           `json:"sentiment_votes_down_percentage"`
	WatchlistUsers               int64                             `json:"watchlist_users"`
	MarketCapRank                int64                             `json:"market_cap_rank"`
	CoingeckoId                  string                            `json:"coingecko_id"`
	CoingeckoRank                int64                             `json:"coingecko_rank"`
	CoingeckoScore               float64                           `json:"coingecko_score"`
	MarketData                   MarketData                        `json:"market_data"`
	CommunityData                interface{}                       `json:"community_data"`
	DeveloperData                interface{}                       `json:"developer_data"`
	Tickers                      []TickerData                      `json:"tickers"`
	ContractAddress              string                            `json:"contract_address"`
	DetailPlatforms              map[string]CoinPlatformDetailData `json:"detail_platforms"`
	// CoingeckoInfo                *TokenInfoResponse                `json:"coingecko_info"`
}

type TokenInfoResponse struct {
	ID               string                     `json:"id"`
	CoingeckoId      string                     `json:"coingecko_id"`
	Name             string                     `json:"name"`
	Image            CoinImage                  `json:"image"`
	MarketData       MarketData                 `json:"market_data"`
	AssetPlatform    *AssetPlatformResponseData `json:"asset_platform"`
	Contracts        []TokenInfoKeyValue        `json:"contracts"`
	Websites         []TokenInfoKeyValue        `json:"websites"`
	Explorers        []TokenInfoKeyValue        `json:"explorers"`
	Wallets          []TokenInfoKeyValue        `json:"wallets"`
	Communities      []TokenInfoKeyValue        `json:"communities"`
	Tags             []TokenInfoKeyValue        `json:"tags"`
	DescriptionLines []string                   `json:"description_lines"`
}

type TokenInfoKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TickerData struct {
	Base         string  `json:"base"`
	Target       string  `json:"target"`
	Last         float32 `json:"last"`
	CoinID       string  `json:"coin_id"`
	TargetCoinID string  `json:"target_coin_id"`
}

type MarketData struct {
	CurrentPrice                           map[string]float64 `json:"current_price"`
	TotalValueLocked                       interface{}        `json:"total_value_locked"`
	McapToTvlRatio                         interface{}        `json:"mcap_to_tvl_ratio"`
	FdvToTvlRatio                          interface{}        `json:"fdv_to_tvl_ratio"`
	Roi                                    interface{}        `json:"roi"`
	Ath                                    map[string]float64 `json:"ath"`
	AthChangePercentage                    map[string]float64 `json:"ath_change_percentage"`
	AthDate                                interface{}        `json:"ath_date"`
	Atl                                    map[string]float64 `json:"atl"`
	MarketCap                              map[string]float64 `json:"market_cap"`
	MarketCapRank                          int64              `json:"market_cap_rank"`
	TotalMarketCap                         map[string]float64 `json:"total_market_cap"`
	TotalVolume                            map[string]float64 `json:"total_volume"`
	FullyDilutedValuation                  map[string]float64 `json:"fully_diluted_valuation"`
	High24h                                map[string]float64 `json:"high_24h"`
	Low24h                                 map[string]float64 `json:"low_24h"`
	PriceChange24h                         float64            `json:"price_change_24h"`
	PriceChangePercentage1h                float64            `json:"price_change_percentage_1h"`
	PriceChangePercentage24h               float64            `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64            `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64            `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64            `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64            `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64            `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64            `json:"price_change_percentage_1y"`
	PriceChange24hInCurrency               map[string]float64 `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      map[string]float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     map[string]float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      map[string]float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     map[string]float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     map[string]float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     map[string]float64 `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    map[string]float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      map[string]float64 `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           map[string]float64 `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency map[string]float64 `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            float64            `json:"total_supply"`
	MaxSupply                              float64            `json:"max_supply"`
	CirculatingSupply                      float64            `json:"circulating_supply"`
}

type CoinImage struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
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
	BaseCoin          *GetCoinResponse                 `json:"base_coin"`
	TargetCoin        *GetCoinResponse                 `json:"target_coin"`
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
	Name            string `json:"name"`
	ShortName       string `json:"shortname"`
}

type GetCoinByContractResponseData struct {
	ID              string                            `json:"id"`
	Name            string                            `json:"name"`
	Symbol          string                            `json:"symbol"`
	DetailPlatforms map[string]CoinPlatformDetailData `json:"detail_platforms"`
	Image           CoinImage                         `json:"image"`
	Decimal         int                               `json:"-"`
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

type GetTopGainerLoser struct {
	TopGainers []GetTopGainerLoserCoin `json:"top_gainers"`
	TopLosers  []GetTopGainerLoserCoin `json:"top_losers"`
}

type GetTopGainerLoserCoin struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Image         string  `json:"image"`
	MarketCapRank int64   `json:"market_cap_rank"`
	Usd           float64 `json:"usd"`
	Usd24hVol     float64 `json:"usd_24h_vol"`
	Usd24hChange  float64 `json:"usd_24h_change"`
	Usd7dChange   float64 `json:"usd_7d_change"`
	Usd1hChange   float64 `json:"usd_1h_change"`
	Usd14dChange  float64 `json:"usd_14d_change"`
	Usdh30dChange float64 `json:"usd_30d_change"`
	Usd60dChange  float64 `json:"usd_60d_change"`
	Usd1yChange   float64 `json:"usd_1y_change"`
}

type GetHistoricalGlobalMarketResponse struct {
	MarketCapChart MarketCapChartData `json:"market_cap_chart"`
}

type MarketCapChartData struct {
	MarketCap [][]float64 `json:"market_cap"`
	Volume    [][]float64 `json:"volume"`
}

type GetGlobalDataResponse struct {
	Data GetGlobalData `json:"data"`
}
type GetGlobalData struct {
	TotalMarketCap map[string]float64 `json:"total_market_cap"`
}
