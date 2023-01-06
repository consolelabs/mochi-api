package onchaintipbottransaction

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/plugin/dbresolver"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) List(q ListQuery) ([]model.OnchainTipBotTransaction, error) {
	var list []model.OnchainTipBotTransaction
	db := pg.db
	if q.SenderDiscordID != "" {
		db = db.Where("sender_discord_id = ?", q.SenderDiscordID)
	}
	if q.RecipientDiscordID != "" {
		db = db.Where("recipient_discord_id = ?", q.RecipientDiscordID)
	}
	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}
	return list, db.Clauses(dbresolver.Write).Find(&list).Error
}

func (pg *pg) GetOnePending(ID int) (*model.OnchainTipBotTransaction, error) {
	tx := &model.OnchainTipBotTransaction{}
	return tx, pg.db.Where("status = ? AND id = ?", "pending", ID).First(tx).Error
}

func (pg *pg) UpsertMany(list []*model.OnchainTipBotTransaction) error {
	tx := pg.db.Begin()
	for _, item := range list {
		tx = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"recipient_address", "status", "claimed_at", "tx_hash"}),
		})
		if err := tx.Create(item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
