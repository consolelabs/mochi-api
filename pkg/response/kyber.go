package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type ProviderSwapRoutes struct {
	Code       int64            `json:"code"`
	Message    string           `json:"message"`
	Data       RouteSummaryData `json:"data"`
	Aggregator string           `json:"aggregator"`
	SwapData   interface{}      `json:"swapData"`
}

type RouteSummaryData struct {
	RouteSummary  model.RouteSummary `json:"routeSummary"`
	RouterAddress string             `json:"routerAddress"`
	TokenIn       RouteToken         `json:"tokenIn"`
	TokenOut      RouteToken         `json:"tokenOut"`
}

type RouteToken struct {
	Id          int64     `json:"id"`
	Address     string    `json:"address"`
	ChainId     int64     `json:"chain_id"`
	ChainName   string    `json:"chain_name"`
	Decimals    int64     `json:"decimals"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	LogoUri     string    `json:"logo_uri"`
	CoingeckoId string    `json:"coingecko_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BuildRoute struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    BuildRouteData `json:"data"`
}

type BuildRouteData struct {
	AmountIn      string       `json:"amountIn"`
	AmountInUsd   string       `json:"amountInUsd"`
	AmountOut     string       `json:"amountOut"`
	AmountOutUsd  string       `json:"amountOutUsd"`
	Gas           string       `json:"gas"`
	GasUsd        string       `json:"gasUsd"`
	OutputChange  OutputChange `json:"outputChange"`
	Data          string       `json:"data"`
	RouterAddress string       `json:"routerAddress"`
}

type OutputChange struct {
	Amount  string `json:"amount"`
	Percent int64  `json:"percent"`
	Level   int64  `json:"level"`
}
