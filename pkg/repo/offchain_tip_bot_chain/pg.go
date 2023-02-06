package offchain_tip_bot_chain

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetAll(f Filter) ([]model.OffchainTipBotChain, error) {
	var token model.OffchainTipBotToken
	db := pg.db

	switch {
	case f.TokenID != "":
		db = db.Where("token_id = ?", f.TokenID)
	case f.TokenSymbol != "":
		db = db.Where("lower(token_symbol) = ?", strings.ToLower(f.TokenSymbol))
	}
	if err := db.First(&token).Error; err != nil {
		return nil, err
	}

	var rs []model.OffchainTipBotChain

	db = pg.db.
		Group("offchain_tip_bot_chains.id,offchain_tip_bot_chains.chain_id").
		Order("offchain_tip_bot_chains.chain_name").
		Preload("Tokens").
		Preload("Contracts").
		Joins(`
	JOIN offchain_tip_bot_tokens_chains tc ON tc.chain_id = offchain_tip_bot_chains.id 
	JOIN offchain_tip_bot_tokens t ON tc.token_id = t.id
	JOIN offchain_tip_bot_contracts c ON c.chain_id = offchain_tip_bot_chains.id`).
		Where("t.id = ?", token.ID)

	if f.IsContractAvailable {
		assignContractQ := pg.db.Table("offchain_tip_bot_assign_contract").
			Select("contract_id").
			Where("token_id = ?", token.ID).
			Where("expired_time > now()")
		if f.UserID != "" {
			assignContractQ = assignContractQ.Where("user_id != ?", f.UserID)
		}
		db = db.Where("c.id NOT IN (?)", assignContractQ)
	}

	return rs, db.Find(&rs).Error
}

func (pg *pg) GetByID(id uuid.UUID) (model.OffchainTipBotChain, error) {
	var chain model.OffchainTipBotChain
	return chain, pg.db.First(&chain, "id = ?", id).Error
}

func (pg *pg) GetByChainID(chainID int) (model.OffchainTipBotChain, error) {
	var rs model.OffchainTipBotChain
	return rs, pg.db.Where("chain_id = ?", chainID).First(&rs).Error
}
