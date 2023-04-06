package response

import "github.com/defipod/mochi/pkg/model"

type SwapRouteResponseData struct {
	Data SwapRouteResponse `json:"data"`
}

type SwapRouteResponse struct {
	Code    int64     `json:"code"`
	Message string    `json:"message"`
	Data    SwapRoute `json:"data"`
}

type SwapRoute struct {
	TokenIn       model.KyberswapSupportedToken `json:"tokenIn"`
	TokenOut      model.KyberswapSupportedToken `json:"tokenOut"`
	RouterAddress string                        `json:"routerAddress"`
	RouteSummary  RouteSummary                  `json:"routeSummary"`
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
	ExtraFee                     model.ExtraFee   `json:"extraFee"`
	Route                        [][]RouteElement `json:"route"`
}

type RouteElement struct {
	Pool              string      `json:"pool"`
	TokenIn           string      `json:"tokenIn"`
	TokenOut          string      `json:"tokenOut"`
	TokenOutSymbol    string      `json:"tokenOutSymbol"`
	LimitReturnAmount string      `json:"limitReturnAmount"`
	SwapAmount        string      `json:"swapAmount"`
	AmountOut         string      `json:"amountOut"`
	Exchange          string      `json:"exchange"`
	PoolLength        int64       `json:"poolLength"`
	PoolType          string      `json:"poolType"`
	PoolExtra         interface{} `json:"poolExtra"`
	Extra             interface{} `json:"extra"`
}
