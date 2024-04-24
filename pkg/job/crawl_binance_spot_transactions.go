package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type crawlBinanceSpotTransactions struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewCrawlBinanceSpotTransactionsJob(e *entities.Entity, l logger.Logger) Job {
	return &crawlBinanceSpotTransactions{
		entity: e,
		log:    l,
	}
}

func (j *crawlBinanceSpotTransactions) Run() error {
	j.log.Infof("Start crawling binance spot transaction ...")
	j.entity.CrawlBinanceSpotTransactions()
	return nil
}
