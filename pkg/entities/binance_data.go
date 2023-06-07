package entities

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) IntegrateBinanceData(req request.IntegrationBinanceData) (*model.KafkaIntegrateMessage, error) {
	res, err := e.svc.MochiProfile.GetByDiscordID(req.DiscordUserId, true)
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	return &kafkaMsg, nil
}
