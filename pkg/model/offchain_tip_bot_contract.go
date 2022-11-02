package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotContract struct {
	ID               uuid.UUID            `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	ChainID          string               `json:"chain_id"`
	ContractAddress  string               `json:"contract_address"`
	Status           int                  `json:"status"`
	AssignStatus     int                  `json:"assign_status"`
	CentralizeWallet string               `json:"centralize_wallet"`
	SweepedTime      *time.Time           `json:"sweeped_time"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
	DeletedAt        *time.Time           `json:"-"`
	Chain            *OffchainTipBotChain `json:"chain"`
}

func (OffchainTipBotContract) TableName() string {
	return "offchain_tip_bot_contracts"
}
