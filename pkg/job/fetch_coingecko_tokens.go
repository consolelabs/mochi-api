package job

import (
	"fmt"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/sentrygo"
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
		j.entity.GetSvc().Sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[CJ prod mochi] - fetch_coingecko_tokens failed - %v", err),
			Tags: map[string]string{
				"type": "system",
			},
			Extra: map[string]interface{}{
				"task": "RefreshCoingeckoSupportedTokenList",
			},
		})
		return err
	}

	j.log.Infof("Successfully refresh coingecko tokens list, %d created", updatedRows)
	return nil
}
