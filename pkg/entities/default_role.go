package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetAllDefaultRoles(guildID string) (*response.DefaultRoleGetAllResponse, error) {
	roles, err := e.repo.GuildConfigDefaultRole.GetAllByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	var res response.DefaultRoleGetAllResponse
	res.Success = true
	res.Data = make([]*response.DefaultRole, 0)
	for _, r := range roles {
		res.Data = append(res.Data, &response.DefaultRole{
			RoleID:  r.RoleID,
			GuildID: r.GuildID,
		})
	}

	return &res, nil
}

func (e *Entity) CreateDefaultRoleConfig(GuildID, RoleID string) error {
	return e.repo.GuildConfigDefaultRole.CreateDefaultRoleIfNotExist(model.GuildConfigDefaultRole{
		RoleID:  RoleID,
		GuildID: GuildID,
	})
}
