package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type updateUserRoles struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateUserRolesJob(e *entities.Entity, l logger.Logger) Job {
	return &updateUserRoles{
		entity: e,
		log:    l,
	}
}

func (c *updateUserRoles) Run() error {
	guilds, err := c.entity.GetGuilds()
	if err != nil {
		return err
	}

	for _, guild := range guilds.Data {
		userXPs, err := c.entity.GetGuildUserXPs(guild.ID)
		if err != nil {
			return err
		}

		lrConfigs, err := c.entity.GetGuildLevelRoleConfigs(guild.ID)
		if err != nil {
			return err
		}

		rolesToAdd := make(map[string]string)
		rolesToRemove := make(map[string]string)
		for _, userXP := range userXPs {
			member, err := c.entity.GetGuildMember(guild.ID, userXP.UserID)
			if err != nil {
				return err
			}

			userLevelRole, err := c.entity.GetUserRoleByLevel(guild.ID, userXP.Level)

			memberRoles := make(map[string]bool)
			for _, roleID := range member.Roles {
				memberRoles[roleID] = true
			}

			// add role if not assigned yet
			if _, ok := memberRoles[userLevelRole]; !ok {
				rolesToAdd[userXP.UserID] = userLevelRole
			}

			for _, lrConfig := range lrConfigs {
				if _, ok := memberRoles[lrConfig.RoleID]; ok && lrConfig.RoleID != userLevelRole {
					rolesToRemove[userXP.UserID] = lrConfig.RoleID
				}
			}
		}

		if err := c.entity.RemoveGuildMemberRoles(guild.ID, rolesToRemove); err != nil {
			return err
		}

		if err := c.entity.AddGuildMemberRoles(guild.ID, rolesToAdd); err != nil {
			return err
		}
	}

	return nil
}
