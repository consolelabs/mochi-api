package kyber

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetSwapRoutesEVM(chain, fromAddress, toAddress, amount string) (*response.ProviderSwapRoutes, error)
	BuildSwapRoutes(chainName string, req *request.KyberBuildSwapRouteRequest) (*response.BuildRoute, error)
}
