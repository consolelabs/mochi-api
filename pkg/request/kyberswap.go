package request

import (
	"math/big"

	"github.com/defipod/mochi/pkg/model"
)

type KyberBuildSwapRouteRequest struct {
	Recipient         string             `json:"recipient"`
	Sender            string             `json:"sender"`
	Source            string             `json:"source"`
	SlippageTolerance int64              `json:"slippageTolerance"`
	SkipSimulateTx    bool               `json:"skipSimulateTx"`
	RouteSummary      model.RouteSummary `json:"routeSummary"`
}

type KyberSwapRequest struct {
	Amount        *big.Int `json:"amount" binding:"required"`
	ChainName     string   `json:"chain_name" binding:"required"`
	EncodedData   string   `json:"encoded_data" binding:"required"`
	RouterAddress string   `json:"router_address" binding:"required"`
	Gas           string   `json:"gas" binding:"required"`
}
