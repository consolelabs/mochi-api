package guildconfigwalletverificationmessage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetOne(guildID string) (*model.GuildConfigWalletVerificationMessage, error) {
	var gcv model.GuildConfigWalletVerificationMessage
	return &gcv, pg.db.Where("guild_id = ?", guildID).First(&gcv).Error
}

func (pg *pg) UpsertOne(gcv model.GuildConfigWalletVerificationMessage) error {

	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
		},
		UpdateAll: true,
	}).Create(&gcv).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) DeleteOne(guildID string) error {
	return pg.db.Where("guild_id = ?", guildID).Delete(&model.GuildConfigWalletVerificationMessage{}).Error
}
