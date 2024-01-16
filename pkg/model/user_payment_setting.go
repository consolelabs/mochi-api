package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

type MoneySource struct {
	Platform           string `json:"platform"`
	PlatformIdentifier string `json:"platform_identifier"`
}

// db explaination for select query
func (m *MoneySource) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("money_source has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, m)
}

// db explaination for insert/update query
func (m MoneySource) Value() (driver.Value, error) {
	bytes, err := json.Marshal(m)
	return string(bytes), err
}

type DefaultMessageSetting struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Enable  bool   `json:"enable"`
}

type DefaultMessageSettings []DefaultMessageSetting

// db explaination for select query
func (s *DefaultMessageSettings) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("default_messaging_settings has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, s)
}

// db explaination for insert/update query
func (s DefaultMessageSettings) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

type TxLimitSetting struct {
	Action string  `json:"action"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
}

type TxLimitSettings []TxLimitSetting

// db explaination for select query
func (s *TxLimitSettings) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("tx_limit_settings has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, s)
}

// db explaination for insert/update query
func (s TxLimitSettings) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

type UserPaymentSetting struct {
	ProfileId               string                 `json:"profile_id"`
	DefaultMoneySource      MoneySource            `json:"default_money_source" gorm:"type:jsonb"`
	DefaultReceiverPlatform string                 `json:"default_receiver_platform"`
	PrioritizedTokenIds     pq.StringArray         `json:"prioritized_token_ids" gorm:"type:text[]"`
	DefaultMessageEnable    bool                   `json:"default_message_enable"`
	DefaultMessageSettings  DefaultMessageSettings `json:"default_message_settings" gorm:"type:jsonb"`
	TxLimitEnable           bool                   `json:"tx_limit_enable"`
	TxLimitSettings         TxLimitSettings        `json:"tx_limit_settings" gorm:"type:jsonb"`
	PrioritizedTokens       []PayToken             `json:"prioritized_token" gorm:"-"`
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
