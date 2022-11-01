package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotToken struct {
	ID          uuid.UUID              `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID     string                 `json:"token_id"`
	TokenName   string                 `json:"token_name"`
	TokenSymbol string                 `json:"token_symbol"`
	Icon        *string                `json:"icon"`
	Status      int                    `json:"status"`
	Chains      []*OffchainTipBotChain `json:"-" gorm:"many2many:offchain_tip_bot_tokens_chains;"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	DeletedAt   *time.Time             `json:"-"`
	CoinGeckoID string                 `json:"coin_gecko_id"`
}

func (OffchainTipBotToken) TableName() string {
	return "offchain_tip_bot_tokens"
}
