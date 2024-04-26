package scheduler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
)

type updateBinanceSpotHistory struct {
	entity *entities.Entity
	log    logger.Logger
	svc    *service.Service
	config config.Config
}

// NewCheckInvalidateEmoji returns a new job that checks for invalid emojis
func NewUpdateBinanceSpotHistory(e *entities.Entity, l logger.Logger, s *service.Service, cfg config.Config) Scheduler {
	return &updateBinanceSpotHistory{
		entity: e,
		log:    l,
		svc:    s,
		config: cfg,
	}
}
func binanceStartTime() time.Time {
	return time.Now().Add(-90 * 24 * time.Hour).UTC()
}
func (s *updateBinanceSpotHistory) Run() error {
	for {
		s.schedulerUpdate()
		// Sleep for an hour interval before checking the database again
		time.Sleep(1 * time.Minute)
	}
}

func (s *updateBinanceSpotHistory) schedulerUpdate() error {

	res, err := s.svc.MochiProfile.GetAllBinanceAccount()
	if err != nil {
		s.log.Error(err, "[updateBinanceSpotHistory] - MochiProfile.GetAllBinanceAccount() fail to get all binance associated account")
		return err
	}
	// get binance exchangeInfo
	data, _, _ := s.svc.Binance.GetExchangeInfo("")
	pairs := []string{}
	for _, d := range data.Symbols {
		pairs = append(pairs, d.Symbol)
	}

	for _, acc := range res.Data {
		binanceTracking, err := s.entity.GetRepo().BinanceTracking.FirstOrCreate(&model.BinanceTracking{ProfileId: acc.ProfileId, SpotLastTime: binanceStartTime()})
		if err != nil {
			s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] - BinanceTracking.FirstOrCreate() fail to first or create binance tracking ")
			continue
		}

		symbols := []string{}
		assetBal, _, _, _ := s.entity.GetBinanceAssets(request.GetBinanceAssetsRequest{
			Id:       acc.ProfileId,
			Platform: "binance",
		})
		for _, asset := range assetBal.Asset {
			symbols = append(symbols, asset.Token.Symbol)
		}
		symbolPairs := make(map[string][]string)

		// Populate the map
		for _, pair := range pairs {
			for _, symbol := range symbols {
				if strings.HasPrefix(pair, symbol) {
					symbolPairs[symbol] = append(symbolPairs[symbol], pair)
					break
				}
			}
		}

		// Print pairs sorted by the order of symbols
		for _, symbol := range symbols {
			pairs := symbolPairs[symbol]
			for _, p := range pairs {
				startTime := strconv.Itoa(int(binanceTracking.SpotLastTime.UnixMilli()))
				endTime := strconv.Itoa(int(binanceTracking.SpotLastTime.Add(24 * time.Hour).UnixMilli()))
				txs, err := s.svc.Binance.GetSpotTransactions(acc.ApiKey, acc.ApiSecret, p, startTime, endTime)
				if err != nil {
					s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] - svc.Binance.GetSpotTransactions() fail to get spot txs")
					break
				}
				for _, tx := range txs {
					err = s.entity.GetRepo().BinanceSpotTransaction.Create(&model.BinanceSpotTransaction{
						ProfileId:               acc.ProfileId,
						Symbol:                  tx.Symbol,
						OrderId:                 tx.OrderId,
						OrderListId:             tx.OrderListId,
						ClientOrderId:           tx.ClientOrderId,
						Price:                   tx.Price,
						OrigQty:                 tx.OrigQty,
						ExecutedQty:             tx.ExecutedQty,
						CumulativeQuoteQty:      tx.CumulativeQuoteQty,
						Status:                  tx.Status,
						TimeInForce:             tx.TimeInForce,
						Type:                    tx.Type,
						Side:                    tx.Side,
						StopPrice:               tx.StopPrice,
						IcebergQty:              tx.IcebergQty,
						IsWorking:               tx.IsWorking,
						OrigQuoteOrderQty:       tx.OrigQuoteOrderQty,
						SelfTradePreventionMode: tx.SelfTradePreventionMode,
						Time:                    time.Unix(tx.Time, 0),
						UpdateTime:              time.Unix(tx.UpdateTime, 0),
					})
					if err != nil {
						fmt.Printf("err: %v", err)
						break
					}
				}
			}
		}
		binanceTracking.SpotLastTime = binanceTracking.SpotLastTime.Add(24 * time.Hour)
		err = s.entity.GetRepo().BinanceTracking.Update(binanceTracking)
		if err != nil {
			s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] -BinanceTracking.Update() fail to update binance tracking")
			break
		}
	}
	return nil
}
