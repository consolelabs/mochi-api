package guild_config_level_role

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

func (pg *pg) GetHighest(guildID string, level int) (*model.GuildConfigLevelRole, error) {
	config := &model.GuildConfigLevelRole{}
	return config, pg.db.Where("guild_id = ? AND level <= ?", guildID, level).Order("level DESC").First(config).Error
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigLevelRole, error) {
	var configs []model.GuildConfigLevelRole
	return configs, pg.db.Where("guild_id = ?", guildID).Preload("LevelConfig").Find(&configs).Error
}

func (pg *pg) UpsertOne(config model.GuildConfigLevelRole) error {
	tx := pg.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "guild_id"}, {Name: "level"}},
		UpdateAll: true,
	}).Create(&config).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) DeleteOne(guildID string, level int) error {
	return pg.db.Where("guild_id = ? AND level = ?", guildID, level).Delete(&model.GuildConfigLevelRole{}).Error
}
