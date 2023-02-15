package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type watchCoinPriceChanges struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewWatchCoinPriceChange(e *entities.Entity, l logger.Logger) Job {
	return &watchCoinPriceChanges{
		entity: e,
		log:    l,
	}
}

// TODO: implement this job later
func (job *watchCoinPriceChanges) Run() error {
	l := job.log
	// cfg := config.LoadConfig(config.DefaultConfigLoaders())

	l.Infof("watchCoinPriceChanges finished")
	return nil
}
