package dune

import (
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/dune/adapter"
)

type Dune struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
	apiKey string
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &Dune{
		config: cfg,
		logger: l,
		cache:  cache,
		apiKey: cfg.DuneApiKey,
	}
}

func (d *Dune) ExecuteQuery(queryId int64, param map[string]any) (res *adapter.ExecuteQueryResponse, err error) {
	d.logger.Debug("start dune.ExecuteQuery()")
	defer d.logger.Debug("end dune.ExecuteQuery()")

	return adapter.ExecuteQuery(d.apiKey, queryId, param)
}

func (d *Dune) GetExecutionResult(executionId string, limit, offset int64) (res *adapter.GetExecutionResultResponse, err error) {
	d.logger.Debug("start dune.GetExecutionResult()")
	defer d.logger.Debug("end dune.GetExecutionResult()")

	return adapter.GetExecutionResult(d.apiKey, executionId, limit, offset)
}

func (d *Dune) GetExecutionStatus(executionId string) (res *adapter.GetExecutionResultResponse, err error) {
	d.logger.Debug("start dune.GeExecutionStatus()")
	defer d.logger.Debug("end dune.GetExecutionStatus()")

	return adapter.GetExecutionStatus(d.apiKey, executionId)
}
