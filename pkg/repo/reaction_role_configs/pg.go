package reaction_role_configs

import (
	"fmt"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
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

func (pg *pg) UpdateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error {
	err := pg.db.Model(&model.ReactionRoleConfig{}).Where("guild_id = ? AND message_id = ?", req.GuildID, req.MessageID).Update("reaction_roles", updateJson).Error
	if err != nil {
		fmt.Errorf("failed to update role configs: %w", err)
	}

	return nil
}

func (pg *pg) CreateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error {
	config := model.ReactionRoleConfig{
		MessageID:     req.MessageID,
		GuildID:       req.GuildID,
		ReactionRoles: updateJson,
	}

	err := pg.db.Create(&config).Error
	if err != nil {
		return err
	}
	return nil
}
