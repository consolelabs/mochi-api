package response

type ExchangeSymbolResponse struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"base_asset"`
	QuoteAsset string `json:"quote_asset"`
}

type GetExchangeInfoResponse struct {
	Timezone string                   `json:"timezone"`
	Symbols  []ExchangeSymbolResponse `json:"symbols"`
}

type GetTickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type GetKlinesDataResponse struct {
	OpenTime         int64
	OPrice           string
	HPrice           string
	LPrice           string
	CPrice           string
	Volume           string
	CloseTime        int64
	QuoteAssetVolume string
	NumOfTrades      int64
}

type WebsocketKlinesDataResponse struct {
	Symbol string              `json:"s"`
	Data   WebsocketKlinesData `json:"k"`
}

type WebsocketKlinesData struct {
	Symbol           string `json:"s"`
	OPrice           string `json:"o"`
	HPrice           string `json:"h"`
	LPrice           string `json:"l"`
	CPrice           string `json:"c"`
	Volume           string `json:"v"`
	QuoteAssetVolume string `json:"q"`
	NumOfTrades      int64  `json:"n"`
}

type WebsocketAggTradeDataResponse struct {
	Symbol string `json:"s"`
	Price  string `json:"p"`
}

type GetAvgPriceBySymbolResponse struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

type BinanceApiKeyPermissionResponse struct {
	IpRestrict                     bool  `json:"ipRestrict"`
	CreateTime                     int64 `json:"createTime"`
	EnableWithdrawals              bool  `json:"enableWithdrawals"`
	EnableInternalTransfer         bool  `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool  `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool  `json:"enableVanillaOptions"`
	EnableReading                  bool  `json:"enableReading"`
	EnableFutures                  bool  `json:"enableFutures"`
	EnableMargin                   bool  `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool  `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime int64 `json:"tradingAuthorityExpirationTime"`
}

type BinanceUserAssetResponse struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	Ipoable      string `json:"ipoable"`
	BtcValuation string `json:"btcValuation"`
}
