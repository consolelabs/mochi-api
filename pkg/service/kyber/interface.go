package kyber

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetSwapRoutes(amount string) (*response.KyberSwapRoutes, error)
}
