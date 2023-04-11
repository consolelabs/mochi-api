package request

import (
	"math/big"

	"github.com/defipod/mochi/pkg/model"
)

type BuildRouteRequest struct {
	Recipient    string             `json:"recipient" binding:"required"`
	Sender       string             `json:"sender" binding:"required"`
	ChainName    string             `json:"chain_name" binding:"required"`
	RouteSummary model.RouteSummary `json:"route_summary" binding:"required"`
}

type KyberBuildSwapRouteRequest struct {
	Recipient         string             `json:"recipient"`
	Sender            string             `json:"sender"`
	Source            string             `json:"source"`
	SlippageTolerance int64              `json:"slippageTolerance"`
	SkipSimulateTx    bool               `json:"skipSimulateTx"`
	RouteSummary      model.RouteSummary `json:"routeSummary"`
}

type SwapRequest struct {
	From            string `json:"from" binding:"required"`
	To              string `json:"to" binding:"required"`
	Amount          string `json:"amount" binding:"required"`
	ChainName       string `json:"chain_name" binding:"required"`
	Gas             string `json:"gas" binding:"required"`
	MinReturnAmount string `json:"min_return_amount" binding:"required"`
	EncodedData     string `json:"encoded_data" binding:"required"`
	RouterAddress   string `json:"router_address" binding:"required"`
}

type KyberSwapRequest struct {
	FromTokenAddress   string   `json:"fromTokenAddress" binding:"required"`
	ToTokenAddress     string   `json:"toTokenAddress" binding:"required"`
	Amount             *big.Int `json:"amount" binding:"required"`
	ChainName          string   `json:"chain_name" binding:"required"`
	CentralizedAddress string   `json:"centralized_address"`
	EncodedData        string   `json:"encoded_data" binding:"required"`
	RouterAddress      string   `json:"router_address" binding:"required"`
	Gas                string   `json:"gas" binding:"required"`
	MinReturnAmount    *big.Int `json:"min_return_amount" binding:"required"`
}
