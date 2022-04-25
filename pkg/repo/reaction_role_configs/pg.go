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

func (pg *pg) GetByMessageID(guildId, messageID string) (model.ReactionRoleConfig, error) {
	var config model.ReactionRoleConfig
	err := pg.db.Model(&model.ReactionRoleConfig{}).Where("guild_id = ? AND message_id = ?", guildId, messageID).First(&config).Error
	if err != nil {
		return config, fmt.Errorf("failed to get role configs: %w", err)
	}

	return config, nil
}
