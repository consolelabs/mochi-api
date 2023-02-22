package job

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

type watchCoinPriceChanges struct {
	entity *entities.Entity
	log    logger.Logger
	cache  cache.Cache
}

type watchCoinPriceChangePayload struct {
	UserID    string  `json:"user_id"`
	Price     float64 `json:"price"`
	Symbol    string  `json:"symbol"`
	Direction string  `json:"direction"`
}

func (m *watchCoinPriceChangePayload) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, m)
}

func (m watchCoinPriceChangePayload) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func NewWatchCoinPriceChange(e *entities.Entity, l logger.Logger) Job {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "[WatchCoinPriceChanges] failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "[WatchCoinPriceChanges] failed to init redis cache")
	}
	return &watchCoinPriceChanges{
		entity: e,
		log:    l,
		cache:  cache,
	}
}

func (job *watchCoinPriceChanges) Run() error {
	var COMMUNICATION_CHANNEL = "channel_dm_price_alert"

	// TODO: Refactor to dynamically fetch list of symbols from database
	// alertSymbols, err := job.entity.GetListAlertSymbols()
	// if err != nil {
	// 	job.log.Error(err, "failed to fetch list alert tokens")
	// 	return err
	// }
	// paramStr := ""
	// for _, v := range alertSymbols {
	// 	paramStr += "/" + strings.ToLower(v) + "usdt@kline_1s"
	// }
	// conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("wss://stream.binance.com:9443/ws%s", paramStr), nil)

	conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/bnbusdt@kline_1s/btcusdt@kline_1s/solusdt@kline_1s/ftmusdt@kline_1s/magicusdt@kline_1s", nil)
	defer conn.Close()
	if err != nil {
		job.log.Error(err, "failed to connect to websocket")
		return err
	}

	// binance will send ping message
	conn.SetPingHandler(func(appData string) error {
		job.log.Info("ping frame responded")
		return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(5*time.Second))
	})

	cooldownMap := make(map[string]bool)
	for {
		_, message, _ := conn.ReadMessage()
		if err != nil {
			break
		}
		var data response.WebsocketKlinesDataResponse
		json.Unmarshal(message, &data)
		if !strings.Contains(data.Symbol, "USDT") {
			continue
		}
		alertCache := []response.ZSetWithScoreData{}
		openPrice, _ := strconv.ParseFloat(data.Data.OPrice, 64)
		closePrice, _ := strconv.ParseFloat(data.Data.CPrice, 64)
		tokenSymbol := data.Symbol[0 : len(data.Symbol)-4] //format: <symbol>USDT
		direction := "up"
		if openPrice-closePrice < 0 {
			alertCache = job.entity.GetPriceAlertZCache(strings.ToLower(data.Symbol), direction, "0", data.Data.CPrice)
		} else {
			direction = "down"
			alertCache = job.entity.GetPriceAlertZCache(strings.ToLower(data.Symbol), direction, data.Data.CPrice, "inf")
		}

		for _, v := range alertCache {
			job.log.Infof("Received %s - price %v :", tokenSymbol, data.Data.CPrice, alertCache)
			payload := watchCoinPriceChangePayload{}
			payload.Price = v.Score
			payload.Direction = direction
			payload.Symbol = data.Symbol
			payload.UserID = v.Member
			cooldownKey := fmt.Sprintf("%v:%v:%v", v.Member, data.Symbol, v.Score)
			if cooldownMap[cooldownKey] {
				continue
			} else {
				cooldownMap[cooldownKey] = true
				time.AfterFunc(60*time.Second, func() {
					job.log.Infof("User with ID %v - Symbol %v - Price %v - 60s cool down end", v.Member, data.Symbol, v.Score)
					cooldownMap[cooldownKey] = false
				})
				job.cache.Publish(COMMUNICATION_CHANNEL, payload)
			}
		}
	}
	return nil
}
