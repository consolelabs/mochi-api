package geckoterminal

type GeckoTerminalSearch struct {
	Data GeckoTerminalSearchData `json:"data,omitempty"`
}

type GeckoTerminalSearchData struct {
	ID         string                        `json:"id,omitempty"`
	Type       string                        `json:"type,omitempty"`
	Attributes GeckoTerminalSearchAttributes `json:"attributes,omitempty"`
}

type GeckoTerminalSearchAttributes struct {
	Networks []interface{}                    `json:"networks,omitempty"`
	Dexes    []interface{}                    `json:"dexes,omitempty"`
	Pools    []GeckoTerminalSearchPoolElement `json:"pools,omitempty"`
	Pairs    []interface{}                    `json:"pairs,omitempty"`
}

type GeckoTerminalSearchPoolElement struct {
	Type               GeckoTerminalSearchType    `json:"type,omitempty"`
	Address            string                     `json:"address,omitempty"`
	APIAddress         string                     `json:"api_address,omitempty"`
	PriceInUsd         string                     `json:"price_in_usd,omitempty"`
	ReserveInUsd       string                     `json:"reserve_in_usd,omitempty"`
	FromVolumeInUsd    string                     `json:"from_volume_in_usd,omitempty"`
	ToVolumeInUsd      string                     `json:"to_volume_in_usd,omitempty"`
	PricePercentChange string                     `json:"price_percent_change,omitempty"`
	Network            *GeckoTerminalSearchDex    `json:"network,omitempty"`
	Dex                *GeckoTerminalSearchDex    `json:"dex,omitempty"`
	Tokens             []GeckoTerminalSearchToken `json:"tokens,omitempty"`
}

type GeckoTerminalSearchDex struct {
	Name       *string `json:"name,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	ImageURL   *string `json:"image_url,omitempty"`
}

type GeckoTerminalSearchToken struct {
	Name        *string `json:"name,omitempty"`
	Symbol      *string `json:"symbol,omitempty"`
	ImageURL    *string `json:"image_url,omitempty"`
	IsBaseToken *bool   `json:"is_base_token,omitempty"`
}

type GeckoTerminalSearchType string

const (
	GeckoTerminalSearchTypePool GeckoTerminalSearchType = "pool"
)

// GetPool
type GeckoTerminalGetPool struct {
	Data     Data       `json:"data"`
	Included []Included `json:"included"`
}

type Data struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    DataAttributes    `json:"attributes"`
	Relationships DataRelationships `json:"relationships"`
}

type DataAttributes struct {
	Address               string `json:"address"`
	Name                  string `json:"name"`
	FullyDilutedValuation string `json:"fully_diluted_valuation"`
	BaseTokenID           string `json:"base_token_id"`
	PriceInUsd            string `json:"price_in_usd"`
	PriceInTargetToken    string `json:"price_in_target_token"`
	// ReserveInUsd              string                               `json:"reserve_in_usd"`
	// ReserveThresholdMet       bool                                 `json:"reserve_threshold_met"`
	FromVolumeInUsd string `json:"from_volume_in_usd"`
	ToVolumeInUsd   string `json:"to_volume_in_usd"`
	APIAddress      string `json:"api_address"`
	// PoolFee                   interface{}                          `json:"pool_fee"`
	// TokenWeightages           interface{}                          `json:"token_weightages"`
	// BalancerPoolID            interface{}                          `json:"balancer_pool_id"`
	// SwapCount24H              int64                                `json:"swap_count_24h"`
	// SwapURL                   string                               `json:"swap_url"`
	// SentimentVotes            SentimentVotes                       `json:"sentiment_votes"`
	PricePercentChange string `json:"price_percent_change"`
	// PricePercentChanges       PricePercentChanges                  `json:"price_percent_changes"`
	HistoricalData HistoricalData `json:"historical_data"`
	// LockedLiquidity           LockedLiquidity                      `json:"locked_liquidity"`
	// SecurityIndicators        []interface{}                        `json:"security_indicators"`
	// PoolReportsCount          int64                                `json:"pool_reports_count"`
	// PoolCreatedAt             interface{}                          `json:"pool_created_at"`
	// LatestSwapTimestamp       string                               `json:"latest_swap_timestamp"`
	// HighLowPriceDataByTokenID map[string]HighLowPriceDataByTokenID `json:"high_low_price_data_by_token_id"`
	// IsNsfw                    bool                                 `json:"is_nsfw"`
}

type HighLowPriceDataByTokenID struct {
	HighPriceInUsd24H     string `json:"high_price_in_usd_24h"`
	HighPriceTimestamp24H string `json:"high_price_timestamp_24h"`
	LowPriceInUsd24H      string `json:"low_price_in_usd_24h"`
	LowPriceTimestamp24H  string `json:"low_price_timestamp_24h"`
}

type HistoricalData struct {
	Last5M  Last `json:"last_5m"`
	Last15M Last `json:"last_15m"`
	Last30M Last `json:"last_30m"`
	Last1H  Last `json:"last_1h"`
	Last6H  Last `json:"last_6h"`
	Last24H Last `json:"last_24h"`
}

type Last struct {
	SwapsCount     int64  `json:"swaps_count"`
	PriceInUsd     string `json:"price_in_usd"`
	VolumeInUsd    string `json:"volume_in_usd"`
	BuySwapsCount  int64  `json:"buy_swaps_count"`
	SellSwapsCount int64  `json:"sell_swaps_count"`
}

type LockedLiquidity struct {
	Source               interface{} `json:"source"`
	UpdatedAt            int64       `json:"updated_at"`
	LockedPercent        interface{} `json:"locked_percent"`
	NextUnlockPercent    interface{} `json:"next_unlock_percent"`
	NextUnlockTimestamp  interface{} `json:"next_unlock_timestamp"`
	FinalUnlockTimestamp interface{} `json:"final_unlock_timestamp"`
}

type PricePercentChanges struct {
	Last5M  string `json:"last_5m"`
	Last15M string `json:"last_15m"`
	Last30M string `json:"last_30m"`
	Last1H  string `json:"last_1h"`
	Last6H  string `json:"last_6h"`
	Last24H string `json:"last_24h"`
}

type SentimentVotes struct {
	Total          int64 `json:"total"`
	UpPercentage   int64 `json:"up_percentage"`
	DownPercentage int64 `json:"down_percentage"`
}

type DataRelationships struct {
	Dex        Dex   `json:"dex"`
	Tokens     Pairs `json:"tokens"`
	PoolMetric Dex   `json:"pool_metric"`
	Pairs      Pairs `json:"pairs"`
}

type Dex struct {
	Data DAT `json:"data"`
}

type DAT struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Pairs struct {
	Data []DAT `json:"data"`
}

type Included struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    IncludedAttributes     `json:"attributes"`
	Relationships *IncludedRelationships `json:"relationships,omitempty"`
}

type IncludedAttributes struct {
	Name                         string                      `json:"name"`
	URL                          *string                     `json:"url,omitempty"`
	Rank                         interface{}                 `json:"rank"`
	CoinURL                      *string                     `json:"coin_url,omitempty"`
	TxURL                        *string                     `json:"tx_url,omitempty"`
	LogoURL                      *string                     `json:"logo_url,omitempty"`
	Identifier                   *string                     `json:"identifier,omitempty"`
	ChainID                      *int64                      `json:"chain_id,omitempty"`
	CGNetworkID                  *string                     `json:"cg_network_id,omitempty"`
	NativeCurrencySymbol         *string                     `json:"native_currency_symbol,omitempty"`
	NativeCurrencyAddress        *string                     `json:"native_currency_address,omitempty"`
	PoolReserveThreshold         *string                     `json:"pool_reserve_threshold,omitempty"`
	ImageURL                     *string                     `json:"image_url,omitempty"`
	IsNew                        *bool                       `json:"is_new,omitempty"`
	ExplorerURL                  *string                     `json:"explorer_url,omitempty"`
	ExplorerCoinURL              *string                     `json:"explorer_coin_url,omitempty"`
	ExplorerTxURL                *string                     `json:"explorer_tx_url,omitempty"`
	ExplorerLogoURL              *string                     `json:"explorer_logo_url,omitempty"`
	AnalyticsPoolPageURL         interface{}                 `json:"analytics_pool_page_url"`
	AnalyticsTokenPageURL        interface{}                 `json:"analytics_token_page_url"`
	Category                     *string                     `json:"category,omitempty"`
	Priority                     *int64                      `json:"priority"`
	BaseName                     *string                     `json:"base_name,omitempty"`
	BaseSymbol                   *string                     `json:"base_symbol,omitempty"`
	BaseAddress                  *string                     `json:"base_address,omitempty"`
	BasePriceInCurrency          *string                     `json:"base_price_in_currency,omitempty"`
	BasePriceInUsdPercentChange  *float64                    `json:"base_price_in_usd_percent_change,omitempty"`
	QuoteName                    *string                     `json:"quote_name,omitempty"`
	QuoteSymbol                  *string                     `json:"quote_symbol,omitempty"`
	QuoteAddress                 *string                     `json:"quote_address,omitempty"`
	QuotePriceInCurrency         *string                     `json:"quote_price_in_currency,omitempty"`
	QuotePriceInUsdPercentChange *float64                    `json:"quote_price_in_usd_percent_change,omitempty"`
	BasePriceInQuote             *string                     `json:"base_price_in_quote,omitempty"`
	QuotePriceInBase             *string                     `json:"quote_price_in_base,omitempty"`
	VolumeInCurrency             *string                     `json:"volume_in_currency,omitempty"`
	VolumeInUsd                  *string                     `json:"volume_in_usd,omitempty"`
	TransactionData              map[string]TransactionDatum `json:"transaction_data,omitempty"`
	VolumeData                   *VolumeData                 `json:"volume_data,omitempty"`
	PriceChangeData              map[string]PriceChangeDatum `json:"price_change_data,omitempty"`
	SwapCount                    *int64                      `json:"swap_count,omitempty"`
	BaseTokenID                  *string                     `json:"base_token_id,omitempty"`
	QuoteTokenID                 *string                     `json:"quote_token_id,omitempty"`
	BasePriceInUsd               *string                     `json:"base_price_in_usd,omitempty"`
	QuotePriceInUsd              *string                     `json:"quote_price_in_usd,omitempty"`
}

type PriceChangeDatum struct {
	Prices               Prices  `json:"prices"`
	BaseTokenUsd         float64 `json:"base_token_usd"`
	QuoteTokenUsd        float64 `json:"quote_token_usd"`
	BaseTokenPercentage  float64 `json:"base_token_percentage"`
	QuoteTokenPercentage float64 `json:"quote_token_percentage"`
}

type Prices struct {
	BaseTokenLowPriceInUsd       string `json:"base_token_low_price_in_usd"`
	BaseTokenHighPriceInUsd      string `json:"base_token_high_price_in_usd"`
	BaseTokenLastPriceInUsd      string `json:"base_token_last_price_in_usd"`
	QuoteTokenLowPriceInUsd      string `json:"quote_token_low_price_in_usd"`
	BaseTokenStartPriceInUsd     string `json:"base_token_start_price_in_usd"`
	QuoteTokenHighPriceInUsd     string `json:"quote_token_high_price_in_usd"`
	QuoteTokenLastPriceInUsd     string `json:"quote_token_last_price_in_usd"`
	BaseTokenLowPriceTimestamp   string `json:"base_token_low_price_timestamp"`
	QuoteTokenStartPriceInUsd    string `json:"quote_token_start_price_in_usd"`
	BaseTokenHighPriceTimestamp  string `json:"base_token_high_price_timestamp"`
	QuoteTokenLowPriceTimestamp  string `json:"quote_token_low_price_timestamp"`
	QuoteTokenHighPriceTimestamp string `json:"quote_token_high_price_timestamp"`
}

type TransactionDatum struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

type VolumeData struct {
	Last300_S   Last300_SClass  `json:"last_300_s"`
	Last900_S   Last300_SClass  `json:"last_900_s"`
	Last1800_S  Last1800_SClass `json:"last_1800_s"`
	Last3600_S  Last1800_SClass `json:"last_3600_s"`
	Last21600_S Last1800_SClass `json:"last_21600_s"`
	Last86400_S Last1800_SClass `json:"last_86400_s"`
}

type Last1800_SClass struct {
	BuysInUsd            string `json:"buys_in_usd"`
	SellsInUsd           string `json:"sells_in_usd"`
	TotalInUsd           string `json:"total_in_usd"`
	TotalInBaseToken     string `json:"total_in_base_token"`
	TotalInQuoteToken    string `json:"total_in_quote_token"`
	BuysInCurrencyToken  string `json:"buys_in_currency_token"`
	SellsInCurrencyToken string `json:"sells_in_currency_token"`
	TotalInCurrencyToken string `json:"total_in_currency_token"`
}

type Last300_SClass struct {
	BuysInUsd            string `json:"buys_in_usd"`
	SellsInUsd           int64  `json:"sells_in_usd"`
	TotalInUsd           string `json:"total_in_usd"`
	TotalInBaseToken     string `json:"total_in_base_token"`
	TotalInQuoteToken    string `json:"total_in_quote_token"`
	BuysInCurrencyToken  string `json:"buys_in_currency_token"`
	SellsInCurrencyToken int64  `json:"sells_in_currency_token"`
	TotalInCurrencyToken string `json:"total_in_currency_token"`
}

type IncludedRelationships struct {
	Explorers     *Pairs `json:"explorers,omitempty"`
	NetworkMetric *Dex   `json:"network_metric,omitempty"`
	Network       *Dex   `json:"network,omitempty"`
	DexMetric     *Dex   `json:"dex_metric,omitempty"`
	Pool          *Dex   `json:"pool,omitempty"`
}

type ScrapePool struct {
	PoolName              string `json:"pool_name"`
	Volume24h             string `json:"volume24h"`
	Liquidity             string `json:"liquidity"`
	FullyDilutedValuation string `json:"fully_diluted_valuation"`
	MarketCap             string `json:"market_cap"`
	PriceInUSD            string `json:"price_in_usd"`
	PriceInTargetToken    string `json:"price_in_target_token"`
}
