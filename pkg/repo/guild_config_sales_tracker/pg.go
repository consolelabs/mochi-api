package guild_config_sales_tracker

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

func (pg *pg) GetByGuildID(guildID string) (config []model.GuildConfigSalesTracker, err error) {
	return config, pg.db.Where("guild_id = ?", guildID).Find(&config).Error
}

func (pg *pg) Create(config *model.GuildConfigSalesTracker) error {
	return pg.db.Create(&config).Error
}

func (pg *pg) GetAllSalesTrackerConfig() (config []model.GuildConfigSalesTracker, err error) {
	return config, pg.db.Find(&config).Error
}
