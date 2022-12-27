package offchain_tip_bot_contract

import (
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) ([]model.OffchainTipBotContract, error) {
	var contracts []model.OffchainTipBotContract
	db := pg.db.Table("offchain_tip_bot_contracts").Joins("JOIN offchain_tip_bot_chains ON offchain_tip_bot_contracts.chain_id = offchain_tip_bot_chains.id")
	if q.ChainID != "" {
		db = db.Where("offchain_tip_bot_contracts.chain_id::TEXT = ?", q.ChainID)
	}
	if q.IsEVM != nil {
		db = db.Where("is_evm = ?", *q.IsEVM)
	}
	if q.SupportDeposit != nil {
		db = db.Where("support_deposit = ?", *q.SupportDeposit)
	}
	return contracts, db.Find(&contracts).Error
}

func (pg *pg) GetByID(id string) (model.OffchainTipBotContract, error) {
	var rs model.OffchainTipBotContract
	return rs, pg.db.Where("id = ?", id).First(&rs).Error
}

func (pg *pg) GetByAddress(addr string) (model.OffchainTipBotContract, error) {
	var rs model.OffchainTipBotContract
	return rs, pg.db.Where("contract_address = ?", addr).First(&rs).Error
}

func (pg *pg) CreateAssignContract(ac *model.OffchainTipBotAssignContract) (*model.OffchainTipBotAssignContract, error) {
	return ac, pg.db.Preload("OffchainTipBotContract").FirstOrCreate(&ac).Error
}

func (pg *pg) DeleteExpiredAssignContract() error {
	return pg.db.Delete(&model.OffchainTipBotAssignContract{}, "expired_time < ?", time.Now()).Error
}

func (pg *pg) GetAssignContract(address, tokenSymbol string) (*model.OffchainTipBotAssignContract, error) {
	result := model.OffchainTipBotAssignContract{}
	return &result, pg.db.Table("offchain_tip_bot_assign_contract").
		Select("offchain_tip_bot_assign_contract.*").
		Joins("JOIN offchain_tip_bot_contracts ON offchain_tip_bot_assign_contract.contract_id = offchain_tip_bot_contracts.id").
		Joins("JOIN offchain_tip_bot_tokens ON offchain_tip_bot_assign_contract.token_id = offchain_tip_bot_tokens.id").
		Where("offchain_tip_bot_contracts.contract_address = ?", address).
		Where("offchain_tip_bot_tokens.token_symbol ILIKE ?", tokenSymbol).
		Where("offchain_tip_bot_assign_contract.expired_time > now()").First(&result).Error
}
