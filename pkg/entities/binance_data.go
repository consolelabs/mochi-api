package entities

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

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

func (e *Entity) UnlinkBinance(req request.UnlinkBinance) error {
	res, err := e.svc.MochiProfile.GetByDiscordID(req.DiscordUserId, true)
	if err != nil {
		e.log.Fields(logger.Fields{"discordUserId": req.DiscordUserId}).Error(err, "[entities.UnlinkBinance] - fail to get profile by discord id")
		return err
	}

	err = e.svc.MochiProfile.UnlinkDex(res.ID, consts.PlatformBinance)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": res.ID, "platform": consts.PlatformBinance}).Error(err, "[entities.UnlinkBinance] - fail to unlink binance")
		return err
	}

	return nil
}

func binanceStartTime() time.Time {
	return time.Unix(1499943600, 0).UTC()
}
func (e *Entity) CrawlBinanceSpotTransactions() {
	e.log.Info("Watching Binance account from profile ...")
	// get all binance associated account
	res, err := e.svc.MochiProfile.GetAllBinanceAccount()
	if err != nil {
		e.log.Error(err, "[entities.CrawlBinanceSpotTransactions] - fail to get all binance associated account")
		return
	}

	for idx, binance := range res.Data {
		binanceTracking, err := e.repo.BinanceTracking.FirstOrCreate(&model.BinanceTracking{ProfileId: binance.ProfileId, SpotLastTime: binanceStartTime()})
		if err != nil {
			e.log.Fields(logger.Fields{"profileId": binance.ProfileId}).Error(err, "[entities.CrawlBinanceSpotTransactions] - fail to first or create binance tracking")
			continue
		}
		res.Data[idx].SpotLastTime = binanceTracking.SpotLastTime
	}

	for _, binance := range res.Data {
		e.log.Fields(logger.Fields{"profileId": binance.ProfileId}).Info("[entities.CrawlBinanceSpotTransactions] - start crawling binance spot transactions")
		// get spot transactions
		startTime := strconv.Itoa(int(binance.SpotLastTime.UnixMilli()))
		endTime := strconv.Itoa(int(time.Date(2024, 2, 20, 0, 0, 0, 0, time.UTC).UnixMilli()))
		_, err := e.svc.Binance.GetSpotTransactions(binance.ApiKey, binance.ApiSecret, startTime, endTime)
		if err != nil {
			e.log.Fields(logger.Fields{"profileId": binance.ProfileId}).Error(err, "[entities.CrawlBinanceSpotTransactions] - fail to get spot transactions")
			continue
		}
	}
}
