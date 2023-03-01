package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotContract struct {
	ID              uuid.UUID            `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	ChainID         uuid.UUID            `json:"chain_id" swaggertype:"string"`
	ContractAddress string               `json:"contract_address"`
	PrivateKey      string               `json:"-"`
	SweepedTime     *time.Time           `json:"sweeped_time"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	DeletedAt       *time.Time           `json:"-"`
	Chain           *OffchainTipBotChain `json:"chain"`
}

func (OffchainTipBotContract) TableName() string {
	return "offchain_tip_bot_contracts"
}
