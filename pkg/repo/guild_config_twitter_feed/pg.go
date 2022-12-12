package guild_config_twitter_feed

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

func (pg *pg) UpsertOne(config *model.GuildConfigTwitterFeed) error {
	tx := pg.db.Begin()
	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) GetAll() ([]model.GuildConfigTwitterFeed, error) {
	configs := []model.GuildConfigTwitterFeed{}
	err := pg.db.Table("guild_config_twitter_feeds").Find(&configs)
	if err.Error != nil {
		return nil, err.Error
	}
	return configs, nil
}
