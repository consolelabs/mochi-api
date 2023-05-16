package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateGuildAdminRoles(req request.CreateGuildAdminRoleRequest) error {
	var config []model.GuildConfigAdminRole
	for _, r := range req.RoleIds {
		config = append(config, model.GuildConfigAdminRole{
			GuildId: req.GuildID,
			RoleId:  r,
		})
	}

	err := e.repo.GuildConfigAdminRole.Create(config)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": req.GuildID,
			"roleIds": req.RoleIds,
		}).Error(err, "[e.CreateGuildAdminRoles] - repo.CreateGuildAdminRoles.Create failed")
		return err
	}

	return nil
}

func (e *Entity) ListGuildAdminRoles(guildID string) ([]model.GuildConfigAdminRole, error) {
	configs, err := e.repo.GuildConfigAdminRole.ListByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[e.ListGuildAdminRoles] - repo.ListGuildAdminRoles.ListByGuildID failed")
		return nil, err
	}
	return configs, nil
}

func (e *Entity) RemoveGuildAdminRole(id int) error {
	return e.repo.GuildConfigAdminRole.Delete(id)
}
