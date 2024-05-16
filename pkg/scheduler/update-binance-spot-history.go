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
	binancespottransaction "github.com/defipod/mochi/pkg/repo/binance_spot_transaction"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/binance"
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
	return time.Now().Add(-1 * time.Hour).UTC()
}
func (s *updateBinanceSpotHistory) Run() error {
	for {
		err := s.schedulerUpdate()
		if err != nil {
			s.log.Error(err, "[updateBinanceSpotHistory] - s.schedulerUpdate()")
			continue
		}
		// Sleep for an hour interval before checking the database again
		time.Sleep(1 * time.Hour)
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
	var pairs []string
	for _, d := range data.Symbols {
		pairs = append(pairs, d.Symbol)
	}

	for _, acc := range res.Data {
		binanceTracking, err := s.entity.GetRepo().BinanceTracking.FirstOrCreate(&model.BinanceTracking{ProfileId: acc.ProfileId, SpotLastTime: binanceStartTime()})
		if err != nil {
			s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] - BinanceTracking.FirstOrCreate() fail to first or create binance tracking ")
			continue
		}
		// update status of NEW order in case it filled or cancel
		newTxs, _ := s.entity.GetRepo().BinanceSpotTransaction.List(binancespottransaction.ListQuery{
			ProfileId: acc.ProfileId,
			Status:    "NEW",
		})
		for _, newTx := range newTxs {
			tx, err := s.svc.Binance.GetSpotTransactionByOrderId(acc.ApiKey, acc.ApiSecret, newTx.Pair, newTx.OrderId)
			if err != nil {
				continue
			}
			if tx.Status != newTx.Status {
				newTx.Status = tx.Status
				newTx.ExecutedQty = tx.ExecutedQty
				newTx.UpdateTime = tx.UpdateTime
				newTx.UpdatedAt = time.UnixMilli(tx.UpdateTime)
			}
			err = s.entity.GetRepo().BinanceSpotTransaction.Update(&newTx)
			if err != nil {
				continue
			}
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
				endTime := strconv.Itoa(int(binanceTracking.SpotLastTime.Add(1 * time.Hour).UnixMilli()))
				txs, err := s.svc.Binance.GetSpotTransactions(acc.ApiKey, acc.ApiSecret, p, startTime, endTime)
				if err != nil {
					s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] - svc.Binance.GetSpotTransactions() fail to get spot txs")
					break
				}
				for _, tx := range txs {
					// If order is filled, and
					// if the pair order is xxxbtc / xxxeth / xxxbnb we will get the price of symbol in usd
					// This data will be used to calculate the average price of the symbol
					isFilled := tx.Status == "FILLED"
					isBtcPair := strings.HasSuffix(p, "BTC")
					isEthPair := strings.HasSuffix(p, "ETH")
					isBnbPair := strings.HasSuffix(p, "BNB")
					priceInUsd := tx.Price
					if isFilled && (isBtcPair || isEthPair || isBnbPair) {
						// get price of the symbol at the time of the transaction
						usdtpair := fmt.Sprintf("%sUSDT", symbol)
						ticks, err := s.svc.Binance.Kline(usdtpair, binance.Interval1m, tx.Time, 0)
						if err != nil {
							s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] - svc.Binance.Kline() fail to get kline")
							break
						}
						if len(ticks) > 0 && len(ticks[0]) > 0 {
							priceInUsd = ticks[0][0]
						}
					}
					err = s.entity.GetRepo().BinanceSpotTransaction.Create(&model.BinanceSpotTransaction{
						ProfileId:               acc.ProfileId,
						Symbol:                  symbol,
						Pair:                    tx.Symbol,
						OrderId:                 tx.OrderId,
						OrderListId:             tx.OrderListId,
						ClientOrderId:           tx.ClientOrderId,
						Price:                   tx.Price,
						PriceInUsd:              priceInUsd,
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
						Time:                    tx.Time,
						UpdateTime:              tx.UpdateTime,
						CreatedAt:               time.UnixMilli(tx.Time),
						UpdatedAt:               time.UnixMilli(tx.UpdateTime),
					})
					if err != nil {
						fmt.Printf("err: %v", err)
						break
					}
				}
			}
		}
		binanceTracking.SpotLastTime = binanceTracking.SpotLastTime.Add(1 * time.Hour)
		err = s.entity.GetRepo().BinanceTracking.Update(binanceTracking)
		if err != nil {
			s.log.Fields(logger.Fields{"profileId": acc.ProfileId}).Error(err, "[updateBinanceSpotHistory] -BinanceTracking.Update() fail to update binance tracking")
			break
		}
	}
	return nil
}
