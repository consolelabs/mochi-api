package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) ConfigLevelRole(req request.ConfigLevelRoleRequest) error {
	err := e.checkRoleIDBeenConfig(req.GuildID, req.RoleID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID}).Error(err, "[entity.ConfigLevelRole] check roleID config failed")
		return err
	}
	return e.repo.GuildConfigLevelRole.UpsertOne(model.GuildConfigLevelRole{
		GuildID: req.GuildID,
		RoleID:  req.RoleID,
		Level:   req.Level,
	})
}

func (e *Entity) GetGuildLevelRoleConfigs(guildID string) ([]model.GuildConfigLevelRole, error) {
	return e.repo.GuildConfigLevelRole.GetByGuildID(guildID)
}

func (e *Entity) RemoveGuildLevelRoleConfig(guildID string, level int) error {
	return e.repo.GuildConfigLevelRole.DeleteOne(guildID, level)
}
