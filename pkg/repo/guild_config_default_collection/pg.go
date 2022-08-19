package guild_config_default_collection

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

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigDefaultCollection, error) {
	config := []model.GuildConfigDefaultCollection{}
	err := pg.db.Table("guild_config_default_collections").Where("guild_id = ?", guildID).Find(&config).Error
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (pg *pg) Upsert(config *model.GuildConfigDefaultCollection) error {
	tx := pg.db.Begin()

	// update on conflict
	err := tx.Table("guild_config_default_collections").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "address"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
