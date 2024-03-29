package guild_config_default_ticker

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

func (pg *pg) GetOneByGuildIDAndQuery(guildID, query string) (*model.GuildConfigDefaultTicker, error) {
	config := &model.GuildConfigDefaultTicker{}
	return config, pg.db.Table("guild_config_default_ticker").Where("guild_id = ? AND query ILIKE ?", guildID, query).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigDefaultTicker) error {
	tx := pg.db.Begin()
	err := tx.Table("guild_config_default_ticker").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "query"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (pg *pg) GetList(guildID string) ([]model.GuildConfigDefaultTicker, error) {
	var configs []model.GuildConfigDefaultTicker
	return configs, pg.db.Table("guild_config_default_ticker").Where("guild_id = ?", guildID).Find(&configs).Error
}
