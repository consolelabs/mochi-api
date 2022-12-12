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

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigSalesTracker, error) {
	config := &model.GuildConfigSalesTracker{}
	return config, pg.db.Table("guild_config_sales_trackers").Where("guild_id = ?", guildID).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigSalesTracker) error {
	return pg.db.Create(&config).Error
}
