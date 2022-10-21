package offchain_tip_bot_activity_logs

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

func (pg *pg) CreateActivityLog(al *model.OffchainTipBotActivityLog) (*model.OffchainTipBotActivityLog, error) {
	return al, pg.db.Create(al).Error
}
