package message

import (
	"github.com/consolelabs/mochi-typeset/typeset"
	"github.com/shopspring/decimal"
)

type KeyPriceChangeAlertMessage struct {
	Type                        typeset.NotificationType    `json:"type"`
	KeyPriceChangeAlertMetadata KeyPriceChangeAlertMetadata `json:"key_price_change_alert_metadata"`
}

type KeyPriceChangeAlertMetadata struct {
	ProfileID      string          `json:"profile_id"`
	KeyAddressID   string          `json:"key_address_id"`
	Change         decimal.Decimal `json:"change"`
	CurrentPrice   decimal.Decimal `json:"current_price"`
	YesterdayPrice decimal.Decimal `json:"yesterday_price"`
}
