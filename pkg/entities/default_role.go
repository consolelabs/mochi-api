package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) GetDefaultRoleByGuildID(guildID string) (*response.DefaultRole, error) {
	role, err := e.repo.GuildConfigDefaultRole.GetAllByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetDefaultRoleByGuildID] failed to get default role by guild id")
		return nil, err
	}

	return &response.DefaultRole{
		RoleID:  role.RoleID,
		GuildID: guildID,
	}, nil
}

func (e *Entity) CreateDefaultRoleConfig(GuildID, RoleID string) error {
	return e.repo.GuildConfigDefaultRole.CreateDefaultRoleIfNotExist(model.GuildConfigDefaultRole{
		RoleID:  RoleID,
		GuildID: GuildID,
	})
}

func (e *Entity) DeleteDefaultRoleConfig(GuildID string) error {
	return e.repo.GuildConfigDefaultRole.DeleteByGuildID(GuildID)
}
