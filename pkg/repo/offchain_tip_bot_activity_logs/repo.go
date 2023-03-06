package offchain_tip_bot_activity_logs

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

func (pg *pg) CreateActivityLog(al *model.OffchainTipBotActivityLog) error {
	return pg.db.Create(al).Error
}

func (pg *pg) List(q ListQuery) ([]model.OffchainTipBotActivityLog, error) {
	var result []model.OffchainTipBotActivityLog
	db := pg.db
	if q.Action != "" {
		db = db.Where("action = ?", q.Action)
	}
	if q.UserID != "" {
		db = db.Where("user_id = ?", q.UserID)
	}
	if q.Message != "" {
		db = db.Where("message = ?", q.Message)
	}
	if q.TokenID != "" {
		db = db.Where("token_id = ?", q.TokenID)
	}
	return result, db.Find(&result).Error
}
