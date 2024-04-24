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
	Asset         string                         `json:"asset"`
	Free          string                         `json:"free"`
	Locked        string                         `json:"locked"`
	Freeze        string                         `json:"freeze"`
	Withdrawing   string                         `json:"withdrawing"`
	Ipoable       string                         `json:"ipoable"`
	BtcValuation  string                         `json:"btcValuation"`
	DetailStaking *BinanceStakingProductPosition `json:"detail_staking"`
	DetailLending *BinancePositionAmountVos      `json:"detail_lending"`
	DetailString  string                         `json:"detail_string"`
}

type BinanceStakingProductPosition struct {
	PositionId          int64  `json:"positionId"`
	ProjectId           string `json:"projectId"`
	Asset               string `json:"asset"`
	Amount              string `json:"amount"`
	PurchaseTime        int64  `json:"purchaseTime"`
	Duration            int64  `json:"duration"`
	AccrualDays         int64  `json:"accrualDays"`
	RewardAsset         string `json:"rewardAsset"`
	RewardAmt           string `json:"rewardAmt"`
	NexInterestPay      string `json:"nexInterestPay"`
	NextInterestPayDate int64  `json:"nextInterestPayDate"`
	PayInterestPeriod   int64  `json:"payInterestPeriod"`
	InterestEndDate     int64  `json:"interestEndDate"`
	DeliveryDate        int64  `json:"deliveryDate"`
	RedeemPeriod        int64  `json:"redeemPeriod"`
	CanRedeemEarly      bool   `json:"canRedeemEarly"`
	Type                string `json:"type"`
	Status              string `json:"status"`
	CanReStake          bool   `json:"canReStake"`
	Apy                 string `json:"apy"`
}

type BinanceLendingAccount struct {
	TotalAmountInBTC       string                     `json:"totalAmountInBTC"`
	TotalAmountInUSDT      string                     `json:"totalAmountInUSDT"`
	TotalFixedAmountInBTC  string                     `json:"totalFixedAmountInBTC"`
	TotalFixedAmountInUSDT string                     `json:"totalFixedAmountInUSDT"`
	TotalFlexibleInBTC     string                     `json:"totalFlexibleInBTC"`
	TotalFlexibleInUSDT    string                     `json:"totalFlexibleInUSDT"`
	PositionAmountVos      []BinancePositionAmountVos `json:"positionAmountVos"`
}

type BinancePositionAmountVos struct {
	Asset        string `json:"asset"`
	Amount       string `json:"amount"`
	AmountInBTC  string `json:"amountInBTC"`
	AmountInUSDT string `json:"amountInUSDT"`
}

type BinanceSimpleEarnAccount struct {
	TotalAmountInBTC          string `json:"totalAmountInBTC"`
	TotalAmountInUSDT         string `json:"totalAmountInUSDT"`
	TotalFlexibleAmountInBTC  string `json:"totalFlexibleAmountInBTC"`
	TotalFlexibleAmountInUSDT string `json:"totalFlexibleAmountInUSDT"`
	TotalLockedInBTC          string `json:"totalLockedInBTC"`
	TotalLockedInUSDT         string `json:"totalLockedInUSDT"`
}

type BinanceFutureAccountBalance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}

type BinanceFutureAccountBalanceResponse struct {
	Data []BinanceFutureAccountBalance `json:"data"`
}

type BinanceFutureAccount struct {
	Positions []BinanceFutureAccountPosition `json:"positions"`
	Assets    []BinanceFutureAccountAsset    `json:"assets"`
}

type BinanceFutureAccountPosition struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	Isolated               bool   `json:"isolated"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int64  `json:"updateTime"`
}

type BinanceFuturePositionInformation struct {
	Positions []BinanceFuturePositionInfo `json:"positions"`
	ApiKey    string                      `json:"apiKey"`
}

type BinanceFutureAccountPositionResponse struct {
	Data []BinanceFuturePositionInformation `json:"data"`
}

type BinanceFutureAccountAsset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

type BinanceFuturePositionInfo struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	UpdateTime       int64  `json:"updateTime"`
}

type BinanceApiTickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type BinanceSpotTransaction struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"order_id"`
	OrderListId             int64  `json:"order_list_id"`
	ClientOrderId           string `json:"client_order_id"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"orig_qty"`
	ExecutedQty             string `json:"executed_qty"`
	CummulativeQuoteQty     string `json:"cummulative_quote_qty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"time_in_force"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stop_price"`
	IcebergQty              string `json:"iceberg_qty"`
	Time                    int64  `json:"time"`
	UpdateTime              int64  `json:"update_time"`
	IsWorking               bool   `json:"is_working"`
	OrigQuoteOrderQty       string `json:"orig_quote_order_qty"`
	WorkingTime             int64  `json:"working_time"`
	SelfTradePreventionMode string `json:"self_trade_prevention_mode"`
}
