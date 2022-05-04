package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetDefaultRoleByGuildID(guildID string) (*response.DefaultRoleResponse, error) {
	role, err := e.repo.GuildConfigDefaultRole.GetAllByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	var res response.DefaultRoleResponse
	res.Success = true
	res.Data = response.DefaultRole{
		RoleID:  role.RoleID,
		GuildID: role.GuildID,
	}

	return &res, nil
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
