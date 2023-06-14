package provider

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Provider interface {
	GetRoute(fromToken, toToken, chain, amount string) (*response.KyberSwapRoutes, error)
	GetRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.KyberSwapRoutes, error)
}
