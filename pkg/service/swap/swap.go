package swap

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	GetAllRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.KyberSwapRoutes, error)
}
