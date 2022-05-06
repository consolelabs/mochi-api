package token

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetBySymbol(symbol string, botSupported bool) (model.Token, error) {
	var token model.Token
	return token, pg.db.First(&token, "lower(symbol) = ? AND discord_bot_supported = ?", symbol, botSupported).Error
}

func (pg *pg) GetAllSupported() ([]model.Token, error) {
	var tokens []model.Token
	return tokens, pg.db.Where("discord_bot_supported = TRUE").Find(&tokens).Error
}

func (pg *pg) GetByAddress(address string, chainID int) (*model.Token, error) {
	var token model.Token
	return &token, pg.db.First(&token, "lower(address) = lower(?) and chain_id = ?", address, chainID).Error
}
