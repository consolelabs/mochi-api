package kyber

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetSwapRoutesEVM(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
	GetSwapRoutesSolana(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
	BuildSwapRoutes(chainName string, req *request.KyberBuildSwapRouteRequest) (*response.BuildRoute, error)
	BuildSwapRoutesSol(chainName, centralizedAddress string, req *request.KyberBuildSwapRouteRequest) (*response.BuildRoute, error)
}
