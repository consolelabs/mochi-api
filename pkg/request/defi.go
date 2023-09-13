package request

import (
	"github.com/gin-gonic/gin"

	"github.com/defipod/mochi/pkg/model"
)

type GetMarketChartRequest struct {
	CoinID           string `json:"coin_id" form:"coin_id" binding:"required"`
	Currency         string `json:"currency" form:"currency,default=usd"`
	Days             int    `json:"days" form:"days,default=7"`
	DiscordID        string `json:"discord_id" form:"discord_id"`
	IsDominanceChart bool   `form:"is_dominance_chart,default=false"`
}

type TransferRequest struct {
	Sender         string   `json:"sender"`
	Recipients     []string `json:"recipients"`
	GuildID        string   `json:"guildId"`
	ChannelID      string   `json:"channelId"`
	Amount         float64  `json:"amount"`
	Cryptocurrency string   `json:"cryptocurrency"`
	Each           bool     `json:"each"`
	All            bool     `json:"all"`
	TransferType   string   `json:"transferType"`
}

func (input *TransferRequest) Bind(c *gin.Context) (err error) {
	err = c.BindJSON(input)
	if err != nil {
		return err
	}

	return err
}

type ListTrackingTokensRequest struct {
	ProfileID   string
	CoinGeckoID string `json:"coin_gecko_id" form:"coin_gecko_id"`
	Page        int    `json:"page" form:"page,default=0"`
	Size        int    `json:"size" form:"size,default=16"`
}

type AddToWatchlistRequest struct {
	WatchlistBaseRequest
	Symbol      string `json:"symbol"`
	CoinGeckoID string `json:"coin_gecko_id"`
	IsFiat      bool   `json:"is_fiat"`
}

type RemoveFromWatchlistRequest struct {
	WatchlistBaseRequest
	Symbol string `json:"symbol"`
}

type GetFiatHistoricalExchangeRatesRequest struct {
	Days   int    `json:"days" form:"days,default=7" binding:"min=7,max=365"`
	Base   string `json:"base" form:"base" binding:"required"`
	Target string `json:"target" form:"target,default=usd"`
}

type AddTokenPriceAlertRequest struct {
	UserDiscordID  string               `json:"user_discord_id"`
	Symbol         string               `json:"symbol"`
	AlertType      model.AlertType      `json:"alert_type" enums:"price_reaches,price_rises_above,price_drops_to,change_is_over,change_is_under"`
	Frequency      model.AlertFrequency `json:"frequency" enums:"only_once,once_a_day,always"`
	Value          float64              `json:"value"`
	PriceByPercent float64              `json:"price_by_percent"`
}

type GetUserListPriceAlertRequest struct {
	UserDiscordID string `json:"user_discord_id" form:"user_discord_id" binding:"required"`
	Page          int    `json:"page" form:"page"`
	Size          int    `json:"size" form:"size"`
}

type GetMarketDataRequest struct {
	Page     string `json:"page" form:"page,default=1"`
	PageSize string `json:"page_size" form:"page_size,default=100"`
	Order    string `json:"order" form:"order"`
}

type TopGainerLoserRequest struct {
	Duration string `json:"duration" form:"duration,default=24h"`
}

type SearchCoinRequest struct {
	Query     string `form:"query"`
	NoDefault bool   `form:"no_default,default=false"`
	GuildId   string `form:"guild_id"`
}

type SearchFriendTechKeysRequest struct {
	Query string `form:"query"`
	Limit int    `form:"limit,default=200"`
}
