package jupiter

import (
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetSwapRoutesSolana(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
	BuildSwapRoutes(chainName string, req *request.JupiterBuildSwapRouteRequest) (*response.BuildRoute, error)
}
