package model

import (
	"time"

	"github.com/google/uuid"
)

type OffchainTipBotAssignContractLog struct {
	ID          uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID     string    `json:"token_id"`
	ChainID     string    `json:"chain_id"`
	UserID      string    `json:"user_id"`
	ContractID  string    `json:"contract_id"`
	Status      int       `json:"status" gorm:"default:0"`
	ExpiredTime time.Time `json:"expired_time"`
}

func (OffchainTipBotAssignContractLog) TableName() string {
	return "offchain_tip_bot_assign_contract_logs"
}
