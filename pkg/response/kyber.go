package response

import "github.com/defipod/mochi/pkg/model"

type KyberSwapRoutes struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    RouteSummaryData `json:"data"`
}

type RouteSummaryData struct {
	RouteSummary  model.RouteSummary            `json:"routeSummary"`
	RouterAddress string                        `json:"routerAddress"`
	TokenIn       model.KyberswapSupportedToken `json:"tokenIn"`
	TokenOut      model.KyberswapSupportedToken `json:"tokenOut"`
}

type BuildRoute struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    BuildRouteData `json:"data"`
}

type BuildRouteData struct {
	AmountIn     string       `json:"amountIn"`
	AmountInUsd  string       `json:"amountInUsd"`
	AmountOut    string       `json:"amountOut"`
	AmountOutUsd string       `json:"amountOutUsd"`
	Gas          string       `json:"gas"`
	GasUsd       string       `json:"gasUsd"`
	OutputChange OutputChange `json:"outputChange"`
	Data         string       `json:"data"`
}

type OutputChange struct {
	Amount  string `json:"amount"`
	Percent int64  `json:"percent"`
	Level   int64  `json:"level"`
}

type KyberSwapRoutesSol struct {
	InputAmount     string                 `json:"inputAmount"`
	OutputAmount    string                 `json:"outputAmount"`
	MinOutputAmount string                 `json:"minOutputAmount"`
	AmountInUsd     float64                `json:"amountInUsd"`
	AmountOutUsd    float64                `json:"amountOutUsd"`
	ReceivedUsd     float64                `json:"receivedUsd"`
	Swaps           [][]model.RouteElement `json:"swaps"`
}
