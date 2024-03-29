package request

type BuildSwapRouteRequest struct {
	Recipient         string      `json:"recipient"`
	Sender            string      `json:"sender"`
	Source            string      `json:"source"`
	SlippageTolerance int64       `json:"slippageTolerance"`
	SkipSimulateTx    bool        `json:"skipSimulateTx"`
	RouteSummary      interface{} `json:"routeSummary"`
	SwapData          interface{} `json:"swapData"`
}

type MochiPaySwapRequest struct {
	ProfileId     string `json:"profile_id"`
	OriginId      string `json:"origin_id"`
	Platform      string `json:"platform"`
	Address       string `json:"address"`
	FromToken     string `json:"from_token"`
	ToToken       string `json:"to_token"`
	ChainId       int64  `json:"chain_id"`
	AmountIn      string `json:"amount_in"`
	AmountOut     string `json:"amount_out"`
	ChainName     string `json:"chain_name"`
	EncodedData   string `json:"encoded_data"`
	RouterAddress string `json:"router_address"`
	Gas           string `json:"gas"`
	Aggregator    string `json:"aggregator"`
	Decimal       int64  `json:"decimal"`
}
