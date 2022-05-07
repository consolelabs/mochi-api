package guild_config_default_roles

import (
	"fmt"
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

func (pg *pg) GetAllByGuildID(guildID string) (model.GuildConfigDefaultRole, error) {
	var role model.GuildConfigDefaultRole
	err := pg.db.Model(&model.GuildConfigDefaultRole{}).Where("guild_id = ?", guildID).First(&role).Error
	if err != nil {
		return role, fmt.Errorf("failed to query default roles: %w", err)
	}

	return role, nil
}

func (pg *pg) CreateDefaultRoleIfNotExist(config model.GuildConfigDefaultRole) error {
	//return pg.db.Create(&config).Error
	tx := pg.db.Begin()

	err := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "guild_id"}},
		DoUpdates: clause.Set{
			{
				Column: clause.Column{Name: "role_id"},
				Value:  config.RoleID,
			},
		},
	}).Create(&config).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (pg *pg) DeleteByGuildID(guildID string) error {
	return pg.db.Where("guild_id = ?", guildID).Delete(&model.GuildConfigDefaultRole{}).Error
}
