package swap

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/swap/provider"
)

type SwapService struct {
	config  *config.Config
	logger  logger.Logger
	kyber   provider.Provider
	jupyter provider.Provider
}

func New(cfg *config.Config, l logger.Logger) Service {
	return &SwapService{
		config:  cfg,
		logger:  l,
		kyber:   provider.NewKyber(cfg, l),
		jupyter: provider.NewJupyter(cfg, l),
	}
}

func (s *SwapService) GetAllRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.ProviderSwapRoutes, error) {
	kyberRoutes, err := s.kyber.GetRoutes(fromTokens, toTokens, amount)
	if err != nil {
		return nil, err
	}

	jupyterRoutes, err := s.jupyter.GetRoutes(fromTokens, toTokens, amount)
	if err != nil {
		return nil, err
	}

	finalRoutes := append(kyberRoutes, jupyterRoutes...)

	return finalRoutes, nil
}

func (s *SwapService) GetBestRoute(routes []response.ProviderSwapRoutes) (*response.ProviderSwapRoutes, error) {
	if len(routes) == 0 {
		return nil, nil
	}

	for _, route := range routes {
		if route.Code == 1 {
			return &route, nil
		}
	}

	return nil, nil
}

func (s *SwapService) BuildSwapRoutes(chainName string, req *request.BuildSwapRouteRequest) (*response.BuildRoute, error) {
	if chainName == "solana" {
		return s.jupyter.BuildSwapRoutes(chainName, req)
	}

	return s.kyber.BuildSwapRoutes(chainName, req)
}
