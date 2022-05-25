package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"gorm.io/gorm"
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
		c.log.Infof("start updating users role - guild %s", guild.ID)
		lrConfigs, err := c.entity.GetGuildLevelRoleConfigs(guild.ID)
		if err != nil {
			return err
		}
		if len(lrConfigs) == 0 {
			c.log.Infof("no levelrole configs found - guild %s", guild.ID)
			continue
		}

		userXPs, err := c.entity.GetGuildUserXPs(guild.ID)
		if err != nil {
			return err
		}
		if len(userXPs) == 0 {
			c.log.Infof("no user XP found - guild %s", guild.ID)
			continue
		}

		rolesToAdd := make(map[string]string)
		rolesToRemove := make(map[string]string)
		for _, userXP := range userXPs {
			member, err := c.entity.GetGuildMember(guild.ID, userXP.UserID)
			if err != nil {
				c.log.Errorf(err, "cannot get guild member %s - guild %s", userXP.UserID, guild.ID)
				continue
			}

			userLevelRole, err := c.entity.GetUserRoleByLevel(guild.ID, userXP.Level)
			if err != nil {
				if err != gorm.ErrRecordNotFound {
					c.log.Errorf(err, "cannot get role by level %d - guild %s", userXP.Level, guild.ID)
					return err
				}
				c.log.Infof("no config found for level %d - guild %s", userXP.Level, guild.ID)
				continue
			}

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
			c.log.Errorf(err, "cannot remove guild member roles - guild %s", guild.ID)
			continue
		}

		if err := c.entity.AddGuildMemberRoles(guild.ID, rolesToAdd); err != nil {
			c.log.Errorf(err, "cannot add guild member roles - guild %s", guild.ID)
			continue
		}
	}

	return nil
}
