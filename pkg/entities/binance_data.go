package entities

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) IntegrateBinanceData(req request.IntegrationBinanceData) (*model.KafkaIntegrateMessage, error) {
	// get profile by discord id
	res, err := e.svc.MochiProfile.GetByDiscordID(req.DiscordUserId, true)
	if err != nil {
		e.log.Fields(logger.Fields{"discordUserId": req.DiscordUserId}).Error(err, "[entities.IntegrateBinanceData] - fail to get profile by discord id")
		return nil, err
	}

	// check permission of binance api key
	permission, err := e.svc.Binance.GetApiKeyPermission(req.ApiKey, req.ApiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"apiKey": req.ApiKey, "apiSecret": req.ApiSecret}).Error(err, "[entities.IntegrateBinanceData] - fail to get api key permission")
		return nil, err
	}

	if !permission.EnableReading {
		return nil, errors.ErrApiKeyBinancePermissionReadingDisabled
	}

	// update api key to profile
	err = e.svc.MochiProfile.AssociateDex(res.ID, consts.PlatformBinance, req.ApiKey, req.ApiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": res.ID, "apiKey": req.ApiKey, "apiSecret": req.ApiSecret}).Error(err, "[entities.IntegrateBinanceData] - fail to update api key to profile")
		return nil, err
	}

	// send key to kafka
	kafkaMsg := model.KafkaIntegrateMessage{
		Type: "binance_integrated_key",
		Data: model.KafkaIntegrateData{
			IntegrationMessage: &model.IntegrationMessage{
				ApiKey:    req.ApiKey,
				ApiSecret: req.ApiSecret,
				ProfileId: res.ID,
			},
		},
	}

	value, err := json.Marshal(kafkaMsg)
	if err != nil {
		return nil, err
	}

	key := strconv.Itoa(rand.Intn(100000))

	err = e.kafka.Produce(e.cfg.Kafka.BinanceDataTopic, key, value)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.IntegrateBinanceData] - fail to produce msg to send api key")
		return nil, err
	}

	return &kafkaMsg, nil
}
