package guild_config_dao_tracker

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
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
func (pg *pg) GetAllBySpaceAndSource(space, source string) (models []model.GuildConfigDaoTracker, err error) {
	return models, pg.db.Where("space = ? AND source = ?", space, source).Find(&models).Error
}
func (pg *pg) DeleteByID(id string) error {
	cfg := model.GuildConfigDaoTracker{}
	return pg.db.Where("id = ?", id).Delete(&cfg).Error
}
func (pg *pg) GetUsageStatsWithPaging(page int, size int) (models []response.DaoTrackerSpaceCountData, total int64, err error) {
	return models, total, pg.db.Table("guild_config_dao_trackers").
		Count(&total).
		Select("space, source, COUNT(space)").
		Group("space, source").
		Offset(size * page).
		Limit(size).
		Order("count DESC").
		Scan(&models).Error
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
