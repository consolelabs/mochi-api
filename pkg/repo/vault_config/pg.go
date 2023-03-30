package vaultconfig

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) Create(vaultConfig *model.VaultConfig) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "channel_id"},
		},
		UpdateAll: true,
	}).Create(&vaultConfig).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func (pg *pg) GetByGuildId(guildId string) (vaultConfig *model.VaultConfig, err error) {
	return vaultConfig, pg.db.Where("guild_id = ?", guildId).First(&vaultConfig).Error
}
