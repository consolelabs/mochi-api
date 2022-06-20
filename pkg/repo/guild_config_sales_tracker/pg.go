package guild_config_sales_tracker

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

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigSalesTracker, error) {
	config := &model.GuildConfigSalesTracker{}
	return config, pg.db.Table("guild_config_sales_trackers").Where("guild_id = ?", guildID).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigSalesTracker) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_config_sales_trackers").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
