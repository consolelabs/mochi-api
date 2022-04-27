package guild_config_default_roles

import (
	"fmt"
	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{db: db}
}

func (pg *pg) GetAllByGuildID(guildID string) ([]model.GuildConfigDefaultRole, error) {
	var roles []model.GuildConfigDefaultRole
	err := pg.db.Model(&model.GuildConfigDefaultRole{}).Where("guild_id = ?", guildID).Scan(&roles).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query default roles: %w", err)
	}

	return roles, nil
}

func (pg *pg) CreateDefaultRoleIfNotExist(config model.GuildConfigDefaultRole) error {
	return pg.db.Create(&config).Error
}
