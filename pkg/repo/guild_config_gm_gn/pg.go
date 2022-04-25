package guild_config_gm_gn

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

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigGmGn, error) {
	config := &model.GuildConfigGmGn{}
	return config, pg.db.Table("guild_config_gm_gn").Where("guild_id = ?", guildID).First(config).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigGmGn) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_config_gm_gn").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
