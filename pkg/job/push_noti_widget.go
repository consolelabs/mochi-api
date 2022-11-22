package job

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/ethereum/go-ethereum/log"
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
		log.Error("failed to connect to websocket")
		return err
	}
	conn.SetPongHandler(func(appData string) error {
		return conn.WriteControl(websocket.PingMessage, []byte(appData), time.Now().Add(5*time.Second))
	})

	data, err := b.entity.GetAllUserTokenAlert()
	if err != nil {
		log.Error("failed to get users alert")
		return err
	}
	userAlerts := data.Data

	for {
		_, message, _ := conn.ReadMessage()
		var trade response.WebsocketKlinesDataResponse
		json.Unmarshal(message, &trade)
		go b.checkAndNotify(&trade, &userAlerts, b.entity)
	}
}

func (b *binanceWebsocket) checkAndNotify(trade *response.WebsocketKlinesDataResponse, userAlerts *[]model.DiscordUserTokenAlert, e *entities.Entity) {
	openPrice, _ := strconv.ParseFloat(trade.Data.OPrice, 64)
	closePrice, _ := strconv.ParseFloat(trade.Data.CPrice, 64)
	for i, alert := range *userAlerts {
		direction := "down"
		if openPrice-closePrice < 0 {
			direction = "up"
		}
		//same symbol and same up-down trend
		if alert.IsEnable && strings.ToUpper(alert.TokenID) == strings.ToUpper(trade.Symbol[0:len(trade.Symbol)-4]) && alert.Trend == direction {
			// if up trend => current price is higher or equal to alert price and reverse
			if (direction == "up" && alert.PriceSet <= closePrice) || (direction == "down" && alert.PriceSet >= closePrice) {
				// disable alert after push noti
				go b.entity.UpsertUserTokenAlert(&request.UpsertDiscordUserAlertRequest{
					ID:        alert.ID.UUID.String(),
					IsEnable:  false,
					TokenID:   alert.TokenID,
					DiscordID: alert.DiscordID,
					PriceSet:  alert.PriceSet,
					Trend:     alert.Trend,
					DeviceID:  alert.DiscordUserDevice.ID,
				})
				(*userAlerts)[i].IsEnable = false
				e.GetSvc().Apns.PushNotificationToIos(alert.DiscordUserDevice.IosNotiToken, alert.PriceSet, alert.Trend, strings.ToUpper(alert.TokenID))
			}
		}
	}
}
