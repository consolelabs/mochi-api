package offchain_tip_bot_contract

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByID(id string) (model.OffchainTipBotContract, error) {
	var rs model.OffchainTipBotContract
	return rs, pg.db.Where("id = ?", id).First(&rs).Error
}

func (pg *pg) GetByAddress(addr string) (model.OffchainTipBotContract, error) {
	var rs model.OffchainTipBotContract
	return rs, pg.db.Where("contract_address = ?", addr).First(&rs).Error
}

func (pg *pg) GetByChainID(id string) ([]model.OffchainTipBotContract, error) {
	var rs []model.OffchainTipBotContract
	return rs, pg.db.Where("chain_id = ?", id).Find(&rs).Error
}

func (pg *pg) CreateAssignContract(ac *model.OffchainTipBotAssignContract) error {
	return pg.db.Create(ac).Error
}

func (pg *pg) DeleteExpiredAssignContract() error {
	return pg.db.Delete(&model.OffchainTipBotAssignContract{}, "expired_time < ?", time.Now()).Error
}
