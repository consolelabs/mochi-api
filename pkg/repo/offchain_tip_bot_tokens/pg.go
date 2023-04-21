package offchain_tip_bot_tokens

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetBySymbol(symbol string) (*model.OffchainTipBotToken, error) {
	var rs model.OffchainTipBotToken
	return &rs, pg.db.Where("token_symbol ILIKE ?", symbol).First(&rs).Error
}

func (pg *pg) Create(t *model.OffchainTipBotToken) error {
	return pg.db.Create(t).Error
}
