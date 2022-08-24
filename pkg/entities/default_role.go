package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) GetDefaultRoleByGuildID(guildID string) (*response.DefaultRoleResponse, error) {
	role, err := e.repo.GuildConfigDefaultRole.GetAllByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetDefaultRoleByGuildID] failed to get default role by guild id")
		return nil, err
	}

	var res response.DefaultRoleResponse
	res.Ok = true
	res.Data = response.DefaultRole{
		RoleID:  role.RoleID,
		GuildID: guildID,
	}

	return &res, nil
}

func (e *Entity) CreateDefaultRoleConfig(GuildID, RoleID string) error {
	err := e.checkRoleIDBeenConfig(GuildID, RoleID)
	if err != nil {
		e.log.Error(err, "[entities.CreateDefaultRoleConfig] check roleID been configed failed")
	}
	return e.repo.GuildConfigDefaultRole.CreateDefaultRoleIfNotExist(model.GuildConfigDefaultRole{
		RoleID:  RoleID,
		GuildID: GuildID,
	})
}

func (e *Entity) DeleteDefaultRoleConfig(GuildID string) error {
	return e.repo.GuildConfigDefaultRole.DeleteByGuildID(GuildID)
}
