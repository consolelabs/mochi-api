package reaction_role_configs

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

func (pg *pg) Gets(guildId string) ([]model.ReactionRoleConfig, error) {
	var guilds []model.ReactionRoleConfig
	err := pg.db.Model(&model.ReactionRoleConfig{}).Find(&guilds, "guild_id = ?", guildId).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get role configs: %w", err)
	}

	return guilds, nil
}
