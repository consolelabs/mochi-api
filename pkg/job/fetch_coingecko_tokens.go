package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type fetchCoingeckoTokens struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewFetchCoingeckoTokensJob(e *entities.Entity, l logger.Logger) Job {
	return &fetchCoingeckoTokens{
		entity: e,
		log:    l,
	}
}

func (j *fetchCoingeckoTokens) Run() error {
	j.log.Infof("Start fetching coingecko supported tokens ...")
	updatedRows, err := j.entity.RefreshCoingeckoSupportedTokensList()
	if err != nil {
		j.log.Error(err, "entity.RefreshCoingeckoSupportedTokensList() failed")
	} else {
		j.log.Infof("Successfully refresh coingecko tokens list, %d created", updatedRows)
	}
	return err
}
