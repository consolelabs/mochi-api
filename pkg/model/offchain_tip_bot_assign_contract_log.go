package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotAssignContractLog struct {
	ID          uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID     uuid.UUID `json:"token_id" swaggertype:"string"`
	ChainID     uuid.UUID `json:"chain_id" swaggertype:"string"`
	UserID      string    `json:"user_id"`
	ContractID  uuid.UUID `json:"contract_id" swaggertype:"string"`
	Status      int       `json:"status" gorm:"default:0"`
	ExpiredTime time.Time `json:"expired_time"`
}

func (OffchainTipBotAssignContractLog) TableName() string {
	return "offchain_tip_bot_assign_contract_logs"
}
