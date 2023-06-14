package provider

import (
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type JupyterProvider struct {
	config *config.Config
	logger logger.Logger
}

func NewJupyter(cfg *config.Config, l logger.Logger) Provider {
	return &JupyterProvider{
		config: cfg,
		logger: l,
	}
}

func (j *JupyterProvider) GetRoute(fromToken, toToken, chain, amount string) (*response.KyberSwapRoutes, error) {
	return nil, nil
}

func (j *JupyterProvider) GetRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.KyberSwapRoutes, error) {
	return nil, nil
}
