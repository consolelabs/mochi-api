package guild_config_reaction_roles

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

func (pg *pg) ListAllByGuildID(guildId string) ([]model.GuildConfigReactionRole, error) {
	var configs []model.GuildConfigReactionRole
	err := pg.db.Model(&model.GuildConfigReactionRole{}).Where("guild_id = ?", guildId).Scan(&configs).Error
	if err != nil {
		return configs, fmt.Errorf("failed to list role configs: %w", err)
	}

	return configs, nil
}

func (pg *pg) GetByMessageID(guildId, messageID string) (model.GuildConfigReactionRole, error) {
	var config model.GuildConfigReactionRole
	return config, pg.db.Model(&model.GuildConfigReactionRole{}).Where("guild_id = ? AND message_id = ?", guildId, messageID).First(&config).Error
}

func (pg *pg) GetByRoleID(guildID, roleID string) (*model.GuildConfigReactionRole, error) {
	config := &model.GuildConfigReactionRole{}
	return config, pg.db.Where("guild_id = ? AND role_id = ?", guildID, roleID).First(config).Error
}

func (pg *pg) UpdateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error {
	err := pg.db.Model(&model.GuildConfigReactionRole{}).Where("guild_id = ? AND message_id = ?", req.GuildID, req.MessageID).Update("reaction_roles", updateJson).Error
	if err != nil {
		return fmt.Errorf("failed to update role configs: %w", err)
	}

	return nil
}

func (pg *pg) ClearMessageConfig(guildID, messageID string) error {
	return pg.db.Where("guild_id = ? AND message_id = ?", guildID, messageID).Delete(&model.GuildConfigReactionRole{}).Error
}

func (pg *pg) CreateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error {
	config := model.GuildConfigReactionRole{
		MessageID:     req.MessageID,
		ChannelID:     req.ChannelID,
		GuildID:       req.GuildID,
		ReactionRoles: updateJson,
	}

	err := pg.db.Create(&config).Error
	if err != nil {
		return err
	}
	return nil
}
