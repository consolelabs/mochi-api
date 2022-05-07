package guild_config_token

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

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigToken, error) {
	var configs []model.GuildConfigToken
	return configs, pg.db.Where("guild_id = ? AND active = TRUE", guildID).Preload("Token").Find(&configs).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigToken) error {
	tx := pg.db.Begin()

	err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
