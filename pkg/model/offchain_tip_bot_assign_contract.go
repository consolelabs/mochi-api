package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OffchainTipBotAssignContract struct {
	ID                     uuid.UUID               `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	TokenID                uuid.UUID               `json:"token_id"`
	ChainID                uuid.UUID               `json:"chain_id"`
	UserID                 string                  `json:"user_id"`
	ContractID             uuid.UUID               `json:"contract_id"`
	CreatedAt              time.Time               `json:"created_at"`
	ExpiredTime            time.Time               `json:"expired_time"`
	OffchainTipBotContract *OffchainTipBotContract `json:"contract" gorm:"foreignkey:ContractID"`
}

func (OffchainTipBotAssignContract) TableName() string {
	return "offchain_tip_bot_assign_contract"
}

func (o *OffchainTipBotAssignContract) BeforeCreate(tx *gorm.DB) (err error) {
	if err := tx.First(&OffchainTipBotAssignContract{},
		"token_id = ? AND contract_id = ? AND chain_id = ? AND expired_time > ?", o.TokenID, o.ContractID, o.ChainID, time.Now()).Error; err == nil {
		return errors.New("contract not found or already assigned")
	}
	return nil
}

func (o *OffchainTipBotAssignContract) AfterCreate(tx *gorm.DB) (err error) {
	if err := tx.Create(&OffchainTipBotAssignContractLog{
		ID:          o.ID,
		TokenID:     o.TokenID,
		ChainID:     o.ChainID,
		UserID:      o.UserID,
		ContractID:  o.ContractID,
		CreatedAt:   o.CreatedAt,
		ExpiredTime: o.ExpiredTime,
	}).Error; err != nil {
		return err
	}
	return nil
}
