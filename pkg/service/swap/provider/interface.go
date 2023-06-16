package provider

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Provider interface {
	GetRoute(fromToken, toToken, chain, amount string) (*response.ProviderSwapRoutes, error)
	GetRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.ProviderSwapRoutes, error)
	BuildSwapRoutes(chainName string, req *request.KyberBuildSwapRouteRequest) (*response.BuildRoute, error)
}
