package swap

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetAllRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.ProviderSwapRoutes, error)
	GetBestRoute(routes []response.ProviderSwapRoutes) (*response.ProviderSwapRoutes, error)
}
