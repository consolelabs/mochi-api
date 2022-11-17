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
