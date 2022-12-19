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
	db := pg.db
	if q.ChainID != "" {
		db = db.Where("chain_id = ?", q.ChainID)
	}
	return contracts, db.Preload("OffchainTipBotChain").Find(&contracts).Error
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
		Joins("LEFT JOIN offchain_tip_bot_contracts ON offchain_tip_bot_assign_contract.contract_id = offchain_tip_bot_contracts.id").
		Joins("LEFT JOIN offchain_tip_bot_tokens ON offchain_tip_bot_assign_contract.token_id = offchain_tip_bot_tokens.id").
		Where("offchain_tip_bot_contracts.contract_address = ? AND offchain_tip_bot_tokens.token_symbol ILIKE ?", address, tokenSymbol).First(&result).Error
}
