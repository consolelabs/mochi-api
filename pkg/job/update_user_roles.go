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
		err = c.updateLevelRoles(guild.ID)
		if err != nil {
			c.log.Fields(logger.Fields{"guildId": guild.ID}).Infof("job.updateLevelRoles failed, error: %v", err)
		}

		err = c.updateNFTRoles(guild.ID)
		if err != nil {
			c.log.Fields(logger.Fields{"guildId": guild.ID}).Infof("job.updateNFTRoles failed, error: %v", err)
		}
	}

	return nil
}

func (c *updateUserRoles) updateLevelRoles(guildID string) error {
	l := c.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("start scanning levelroles")
	lrConfigs, err := c.entity.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		l.Error(err, "entity.GetGuildLevelRoleConfigs failed")
		return err
	}

	if len(lrConfigs) == 0 {
		l.Info("entity.GetGuildLevelRoleConfigs - no data found")
		return nil
	}

	userXPs, err := c.entity.GetGuildUserXPs(guildID)
	if err != nil {
		l.Error(err, "entity.GetGuildUserXPs failed")
		return err
	}
	if len(userXPs) == 0 {
		l.Info("entity.GetGuildUserXPs - no data found")
		return nil
	}

	rolesToAdd := make(map[string]string)
	rolesToRemove := make(map[string]string)
	for _, userXP := range userXPs {
		member, err := c.entity.GetGuildMember(guildID, userXP.UserID)
		if err != nil {
			c.log.Fields(logger.Fields{
				"userId":  userXP.UserID,
				"guildId": guildID,
			}).Error(err, "entity.GetGuildMember failed")
			continue
		}

		userLevelRole, err := c.entity.GetUserRoleByLevel(guildID, userXP.Level)
		if err != nil {
			c.log.Fields(logger.Fields{
				"level":   userXP.Level,
				"guildId": guildID,
			}).Error(err, "entity.GetUserRoleByLevel failed")
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

	err = c.entity.RemoveGuildMemberRoles(guildID, rolesToRemove)
	if err != nil {
		c.log.Fields(logger.Fields{
			"guildId":       guildID,
			"rolesToRemove": rolesToRemove,
		}).Error(err, "entity.RemoveGuildMemberRoles failed")
	}
	c.log.Fields(logger.Fields{
		"guildId":       guildID,
		"rolesToRemove": rolesToRemove,
	}).Info("entity.RemoveGuildMemberRoles executed successfully")

	if err := c.entity.AddGuildMemberRoles(guildID, rolesToAdd); err != nil {
		c.log.Fields(logger.Fields{
			"guildId":    guildID,
			"rolesToAdd": rolesToAdd,
		}).Error(err, "entity.AddGuildMemberRoles failed")
	}
	c.log.Fields(logger.Fields{
		"guildId":       guildID,
		"rolesToRemove": rolesToAdd,
	}).Info("entity.AddGuildMemberRoles executed successfully")

	return nil
}

func (c *updateUserRoles) updateNFTRoles(guildID string) error {
	l := c.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("start scanning nftroles")
	hrConfigs, err := c.entity.ListGuildNFTRoleConfigs(guildID)
	if err != nil {
		l.Error(err, "entity.ListGuildNFTRoleConfigs failed")
		return err
	}

	if len(hrConfigs) == 0 {
		l.Info("entity.ListGuildNFTRoleConfigs - no data found")
		return nil
	}

	isNFTRoles := make(map[string]bool)
	for _, hrConfig := range hrConfigs {
		isNFTRoles[hrConfig.RoleID] = true
	}

	rolesToAdd, err := c.entity.ListMemberNFTRolesToAdd(guildID)
	if err != nil {
		l.Error(err, "entity.ListMemberNFTRolesToAdd failed")
		return err
	}

	members, err := c.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[job][pdateNFTRoles] entity.ListGuildMembers failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isNFTRoles[roleID] {
				if rolesToAdd[[2]string{member.User.ID, roleID}] {
					delete(rolesToAdd, [2]string{member.User.ID, roleID})
					continue
				}

				gMemberRoleLog := c.log.Fields(logger.Fields{
					"guildId": guildID,
					"userId":  member.User.ID,
					"roleId":  roleID,
				})
				err = c.entity.RemoveGuildMemberRole(guildID, member.User.ID, roleID)
				if err != nil {
					gMemberRoleLog.Error(err, "entity.RemoveGuildMemberRole failed")
				}
				gMemberRoleLog.Info("entity.RemoveGuildMemberRole executed successfully")
			}
		}
	}

	for roleToAdd := range rolesToAdd {
		gMemberRoleLog := c.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  roleToAdd[0],
			"roleId":  roleToAdd[1],
		})
		err = c.entity.AddGuildMemberRole(guildID, roleToAdd[0], roleToAdd[1])
		if err != nil {
			gMemberRoleLog.Error(err, "entity.AddGuildMemberRole failed")
		}
		gMemberRoleLog.Info("entity.AddGuildMemberRole executed successfully")
	}

	return nil
}
