package model

import (
	"github.com/lib/pq"
)

type MoneySource struct {
	Platform           string `json:"platform"`
	PlatformIdentifier string `json:"platform_identifier"`
}

type DefaultMessageSetting struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type TxLimitSetting struct {
	Action  string    `json:"action"`
	Min     float64   `json:"min"`
	Max     float64   `json:"max"`
	TokenId string    `json:"token_id"`
	Token   *PayToken `json:"token" gorm:"foreignkey:TokenId;<-:false"`
}

type UserPaymentSetting struct {
	ProfileId               string
	DefaultMoneySource      MoneySource             `json:"default_money_source" gorm:"type:json"`
	DefaultReceiverPlatform string                  `json:"default_receiver_platform"`
	PrioritizedTokenIds     pq.StringArray          `json:"prioritized_token_ids" gorm:"type:text[]"`
	DefaultTokenId          string                  `json:"default_token_id"`
	DefaultMessageEnable    bool                    `json:"default_message_enable"`
	DefaultMessageSettings  []DefaultMessageSetting `json:"default_message_settings" gorm:"type:json"`
	TxLimitEnable           bool                    `json:"tx_limit_enable"`
	TxLimitSettings         []TxLimitSetting        `json:"tx_limit_settings" gorm:"type:json"`
	DefaultToken            *PayToken               `json:"default_token" gorm:"foreignkey:TokenId;<-:false"`
	PrioritizedTokens       []PayToken              `json:"prioritized_token" gorm:"-"`
}

// TODO: remove this
type PayToken struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	Decimal     int64   `json:"decimal"`
	ChainId     string  `json:"chain_id"`
	Native      bool    `json:"native"`
	Address     string  `json:"address"`
	Icon        string  `json:"icon"`
	Price       float64 `json:"price"`
	Chain       Chain   `json:"chain"`
	CoinGeckoId string  `json:"coin_gecko_id"`
}
