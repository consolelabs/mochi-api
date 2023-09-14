package message

import (
	"time"

	"github.com/consolelabs/mochi-typeset/typeset"
	"github.com/shopspring/decimal"
)

type KeyPriceChangeAlertMessage struct {
	Type                        typeset.NotificationType    `json:"type"`
	KeyPriceChangeAlertMetadata KeyPriceChangeAlertMetadata `json:"key_price_change_alert_metadata"`
}

type KeyPriceChangeAlertMetadata struct {
	Timestamp      time.Time       `json:"timestamp"`
	ProfileID      string          `json:"profile_id"`
	KeyAddressID   string          `json:"key_address_id"`
	KeyAddress     FriendTechKey   `json:"key_address"`
	Change         decimal.Decimal `json:"change"`
	CurrentPrice   decimal.Decimal `json:"current_price"`
	YesterdayPrice decimal.Decimal `json:"yesterday_price"`
}

type FriendTechKey struct {
	ID              int64           `json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Address         string          `json:"address"`
	TwitterUsername string          `json:"twitter_username"`
	TwitterPfpUrl   string          `json:"twitter_pfp_url"`
	ProfileChecked  bool            `json:"profile_checked"`
	Price           decimal.Decimal `json:"price"`
	Supply          int             `json:"supply"`
	Holders         int             `json:"holders"`
}
