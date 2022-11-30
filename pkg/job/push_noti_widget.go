package job

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/gorilla/websocket"
)

type binanceWebsocket struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewPushNotiWidgetJob(e *entities.Entity, l logger.Logger) Job {
	return &binanceWebsocket{
		log:    l,
		entity: e,
	}
}

type WSrequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

func (b *binanceWebsocket) Run() error {
	conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/bnbusdt@kline_1m/btcusdt@kline_1m/ethusdt@kline_1m/ftmusdt@kline_1m/avaxusdt@kline_1m/maticusdt@kline_1m/solusdt@kline_1m/icpusdt@kline_1m", nil)
	defer conn.Close()
	if err != nil {
		b.log.Error(err, "failed to connect to websocket")
		return err
	}

	// binance will send ping message
	conn.SetPingHandler(func(appData string) error {
		b.log.Info("ping frame responded")
		return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(5*time.Second))
	})

	for {
		_, message, _ := conn.ReadMessage()
		var trade response.WebsocketKlinesDataResponse
		json.Unmarshal(message, &trade)
		go b.checkAndNotify(&trade, b.entity)
	}
}

func (b *binanceWebsocket) checkAndNotify(trade *response.WebsocketKlinesDataResponse, e *entities.Entity) {
	openPrice, _ := strconv.ParseFloat(trade.Data.OPrice, 64)
	closePrice, _ := strconv.ParseFloat(trade.Data.CPrice, 64)
	tokenSymbol := trade.Symbol[0 : len(trade.Symbol)-4] //format: <symbol>USDT
	alertCache := []string{}
	direction := "up"
	if openPrice-closePrice < 0 {
		alertCache = b.entity.GetTokenAlertZCache(strings.ToLower(tokenSymbol), direction, "0", trade.Data.CPrice)
	} else {
		direction = "down"
		alertCache = b.entity.GetTokenAlertZCache(strings.ToLower(tokenSymbol), direction, trade.Data.CPrice, "inf")
	}
	for _, id := range alertCache {
		alert, err := b.entity.GetUserTokenAlertByID(id)
		if err != nil || alert == nil {
			b.log.Errorf(err, "failed to get users alert id: %s", id)
			continue
		}
		// if up trend => current price is higher or equal to alert price and reverse
		if alert.IsEnable {
			b.log.Infof("sending alert to id: %s", id)
			// disable alert after push noti
			err := b.entity.UpsertUserTokenAlert(&request.UpsertDiscordUserAlertRequest{
				ID:        alert.ID.UUID.String(),
				IsEnable:  false,
				TokenID:   alert.TokenID,
				Symbol:    alert.Symbol,
				DiscordID: alert.DiscordID,
				PriceSet:  alert.PriceSet,
				Trend:     alert.Trend,
				DeviceID:  alert.DiscordUserDevice.ID,
			})
			if err != nil {
				b.log.Errorf(err, "failed to update alert id: %s", id)
			}

			res, err := e.GetSvc().Apns.PushNotificationToIos(alert.DiscordUserDevice.IosNotiToken, alert.PriceSet, alert.Trend, strings.ToUpper(alert.Symbol))
			if err != nil {
				b.log.Errorf(err, "failed to push notification id: %s", id)
			}
			if res.Sent() {
				b.log.Infof("successfully sent alert to token: %s", alert.DiscordUserDevice.IosNotiToken)
			}
			// remove cache after use
			err = b.entity.DeleteTokenAlertZCache(tokenSymbol, direction, id)
			if err != nil {
				b.log.Errorf(err, "failed to remove cache id: %s", id)
			}
		}

	}
}
