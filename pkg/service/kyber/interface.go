package kyber

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetSwapRoutesEVM(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
	GetSwapRoutesSolana(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
}
