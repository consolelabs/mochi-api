package guild_config_default_currency

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

func (pg *pg) GetByGuildID(guildID string) (*model.GuildConfigDefaultCurrency, error) {
	config := &model.GuildConfigDefaultCurrency{}
	return config, pg.db.Table("guild_config_default_currencies").Preload("TipBotToken").Where("guild_id = ?", guildID).First(config).Error
}
func (pg *pg) Upsert(config *model.UpsertGuildConfigDefaultCurrency) error {
	tx := pg.db.Begin()
	err := tx.Table("guild_config_default_currencies").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}},
		UpdateAll: true,
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (pg *pg) DeleteByGuildID(guildID string) error {
	config := &model.GuildConfigDefaultCurrency{}
	return pg.db.Table("guild_config_default_currencies").Where("guild_id = ?", guildID).Delete(config).Error
}
