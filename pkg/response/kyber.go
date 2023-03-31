package response

type KyberSwapRoutes struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    RouteSummaryData `json:"data"`
}

type RouteSummaryData struct {
	RouteSummary  RouteSummary `json:"routeSummary"`
	RouterAddress string       `json:"routerAddress"`
}

type RouteSummary struct {
	TokenIn                      string           `json:"tokenIn"`
	AmountIn                     string           `json:"amountIn"`
	AmountInUsd                  string           `json:"amountInUsd"`
	TokenInMarketPriceAvailable  bool             `json:"tokenInMarketPriceAvailable"`
	TokenOut                     string           `json:"tokenOut"`
	AmountOut                    string           `json:"amountOut"`
	AmountOutUsd                 string           `json:"amountOutUsd"`
	TokenOutMarketPriceAvailable bool             `json:"tokenOutMarketPriceAvailable"`
	Gas                          string           `json:"gas"`
	GasPrice                     string           `json:"gasPrice"`
	GasUsd                       string           `json:"gasUsd"`
	ExtraFee                     ExtraFee         `json:"extraFee"`
	Route                        [][]RouteElement `json:"route"`
}

type ExtraFee struct {
	FeeAmount   string `json:"feeAmount"`
	ChargeFeeBy string `json:"chargeFeeBy"`
	IsImBps     bool   `json:"isImBps"`
	FeeReceiver string `json:"feeReceiver"`
}

type RouteElement struct {
	Pool              string      `json:"pool"`
	TokenIn           string      `json:"tokenIn"`
	TokenOut          string      `json:"tokenOut"`
	LimitReturnAmount string      `json:"limitReturnAmount"`
	SwapAmount        string      `json:"swapAmount"`
	AmountOut         string      `json:"amountOut"`
	Exchange          string      `json:"exchange"`
	PoolLength        int64       `json:"poolLength"`
	PoolType          string      `json:"poolType"`
	PoolExtra         interface{} `json:"poolExtra"`
	Extra             interface{} `json:"extra"`
}
