package offchain_tip_bot_tokens

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

func (pg *pg) GetBySymbol(symbol string) (*model.OffchainTipBotToken, error) {
	var rs model.OffchainTipBotToken
	return &rs, pg.db.Where("token_symbol = ?", symbol).First(&rs).Error
}
