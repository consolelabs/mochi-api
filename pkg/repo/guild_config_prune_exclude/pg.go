package guild_config_prune_exclude

import (
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigWhitelistPrune, error) {
	configs := []model.GuildConfigWhitelistPrune{}
	return configs, pg.db.Where("guild_id = ?", guildID).Find(&configs).Error
}

func (pg *pg) UpsertOne(config *model.GuildConfigWhitelistPrune) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "role_id"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) DeleteOne(config *model.GuildConfigWhitelistPrune) error {
	return pg.db.Where("guild_id = ? AND role_id = ?", config.GuildID, config.RoleID).Delete(&config).Error
}
