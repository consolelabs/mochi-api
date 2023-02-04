package guild_config_dao_tracker

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

func (pg *pg) GetAllByGuildID(guildId string) (models *[]model.GuildConfigDaoTracker, err error) {
	return models, pg.db.Where("guild_id = ?", guildId).Find(&models).Error
}
func (pg *pg) GetAllBySpace(space string) (models []model.GuildConfigDaoTracker, err error) {
	return models, pg.db.Where("space = ?", space).Find(&models).Error
}
func (pg *pg) DeleteByID(id string) error {
	cfg := model.GuildConfigDaoTracker{}
	return pg.db.Where("id = ?", id).Delete(&cfg).Error
}
func (pg *pg) Upsert(cfg model.GuildConfigDaoTracker) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "guild_id"},
			{Name: "space"},
		},
		UpdateAll: true,
	}).Create(&cfg).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
