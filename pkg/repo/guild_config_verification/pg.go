package guildconfigverification

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetOne(guildID string) (*model.GuildConfigVerification, error) {
	var gcv model.GuildConfigVerification
	return &gcv, pg.db.Where("guild_id = ?", guildID).First(&gcv).Error
}

func (pg *pg) UpsertOne(gcv model.GuildConfigVerification) error {

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
