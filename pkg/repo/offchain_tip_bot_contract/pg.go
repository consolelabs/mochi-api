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
	return ac, pg.db.
		Where("chain_id = ? AND token_id = ? AND contract_id = ? AND user_id = ? AND expired_time >= now()", ac.ChainID, ac.TokenID, ac.ContractID, ac.UserID).
		Preload("OffchainTipBotContract").FirstOrCreate(&ac).Error
}

func (pg *pg) DeleteExpiredAssignContract() error {
	return pg.db.Delete(&model.OffchainTipBotAssignContract{}, "expired_time < ?", time.Now()).Error
}

func (pg *pg) GetAssignContract(q GetAssignContractQuery) (*model.OffchainTipBotAssignContract, error) {
	result := &model.OffchainTipBotAssignContract{}
	db := pg.db.Table("offchain_tip_bot_assign_contract").
		Select("offchain_tip_bot_assign_contract.*").
		Joins("JOIN offchain_tip_bot_contracts ON offchain_tip_bot_assign_contract.contract_id = offchain_tip_bot_contracts.id").
		Joins("JOIN offchain_tip_bot_tokens ON offchain_tip_bot_assign_contract.token_id = offchain_tip_bot_tokens.id")
	if q.Address != "" {
		db = db.Where("offchain_tip_bot_contracts.contract_address = ?", q.Address)
	}
	if q.TokenSymbol != "" {
		db = db.Where("offchain_tip_bot_tokens.token_symbol ILIKE ?", q.TokenSymbol)
	}
	// get assigned contract at the given time OR get current assigned
	if q.SignedAt != nil {
		db = db.
			Where("EXTRACT(EPOCH FROM offchain_tip_bot_assign_contract.created_at) * 1000 <= ?", *q.SignedAt).
			Where("? <= EXTRACT(EPOCH FROM offchain_tip_bot_assign_contract.expired_time) * 1000", *q.SignedAt)
	} else {
		db = db.Where("offchain_tip_bot_assign_contract.expired_time > now()")
	}
	return result, db.First(result).Error
}

func (pg *pg) UpdateSweepTime(contractID string, sweepTime time.Time) error {
	return pg.db.Model(&model.OffchainTipBotContract{}).Where("id::TEXT = ?", contractID).Update("sweeped_time", sweepTime).Error
}
