package guild_config_level_role

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/defipod/mochi/pkg/model"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *pg {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetCurrentLevelRole(guildID string, level int) (cfg *model.GuildConfigLevelRole, err error) {
	return cfg, pg.db.Where("guild_id = ? AND level <= ?", guildID, level).Order("level DESC").Preload("LevelConfig").First(&cfg).Error
}

func (pg *pg) GetNextLevelRole(guildID string, currentLevel int) (cfg *model.GuildConfigLevelRole, err error) {
	return cfg, pg.db.Where("guild_id = ? AND level > ?", guildID, currentLevel).Order("level ASC").Preload("LevelConfig").First(&cfg).Error
}

func (pg *pg) GetByGuildID(guildID string) ([]model.GuildConfigLevelRole, error) {
	var configs []model.GuildConfigLevelRole
	return configs, pg.db.Where("guild_id = ?", guildID).Preload("LevelConfig").Find(&configs).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigLevelRole, error) {
	config := &model.GuildConfigLevelRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
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
