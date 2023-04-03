package kyber

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetSwapRoutes(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error)
}
