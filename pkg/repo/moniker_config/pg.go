package moniker_config

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetByGuildID(guildID string) ([]model.MonikerConfig, error) {
	var configs []model.MonikerConfig
	return configs, pg.db.Preload("Token").Where("guild_id = ?", guildID).Find(&configs).Error
}

func (pg *pg) UpsertOne(record model.MonikerConfig) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "moniker"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"plural", "token_id", "amount", "updated_at"}),
	}).Create(&record).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) DeleteOne(guildID, moniker string) error {
	return pg.db.Where("guild_id = ? AND moniker = ?", guildID, moniker).Delete(&model.MonikerConfig{}).Error
}

func (pg *pg) GetDefaultMoniker() (configs []model.MonikerConfig, err error) {
	return configs, pg.db.Preload("Token").Where("guild_id = ?", "*").Find(&configs).Error
}
