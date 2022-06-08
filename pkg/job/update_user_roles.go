package job

import (
	"fmt"

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
		err = c.updateLevelRoles(guild.ID)
		if err != nil {
			return err
		}

		err = c.updateNFTRoles(guild.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *updateUserRoles) updateLevelRoles(guildID string) error {
	c.log.Infof("start updating users role - guild %s", guildID)
	lrConfigs, err := c.entity.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		return err
	}

	if len(lrConfigs) == 0 {
		c.log.Infof("no levelrole configs found - guild %s", guildID)
		return nil
	}

	userXPs, err := c.entity.GetGuildUserXPs(guildID)
	if err != nil {
		return err
	}
	if len(userXPs) == 0 {
		c.log.Infof("no user XP found - guild %s", guildID)
		return nil
	}

	rolesToAdd := make(map[string]string)
	rolesToRemove := make(map[string]string)
	for _, userXP := range userXPs {
		member, err := c.entity.GetGuildMember(guildID, userXP.UserID)
		if err != nil {
			c.log.Errorf(err, "cannot get guild member %s - guild %s", userXP.UserID, guildID)
			return err
		}

		userLevelRole, err := c.entity.GetUserRoleByLevel(guildID, userXP.Level)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				c.log.Errorf(err, "cannot get role by level %d - guild %s", userXP.Level, guildID)
				return err
			}
			c.log.Infof("no config found for level %d - guild %s", userXP.Level, guildID)
			return err
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

	if err := c.entity.RemoveGuildMemberRoles(guildID, rolesToRemove); err != nil {
		c.log.Errorf(err, "cannot remove guild member roles - guild %s", guildID)
		return err
	}

	if err := c.entity.AddGuildMemberRoles(guildID, rolesToAdd); err != nil {
		c.log.Errorf(err, "cannot add guild member roles - guild %s", guildID)
		return err
	}

	return nil
}

func (c *updateUserRoles) updateNFTRoles(guildID string) error {
	hrConfigs, err := c.entity.ListGuildNFTRoleConfigs(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild nft role configs: %v", err.Error())
	}

	if len(hrConfigs) == 0 {
		c.log.Infof("no nftrole configs found - guild %s", guildID)
		return nil
	}

	isNFTRoles := make(map[string]bool)
	for _, hrConfig := range hrConfigs {
		isNFTRoles[hrConfig.RoleID] = true
	}

	rolesToAdd, err := c.entity.ListMemberNFTRolesToAdd(guildID)
	if err != nil {
		return fmt.Errorf("failed to get member nft roles to add: %v", err.Error())
	}

	members, err := c.entity.ListGuildMembers(guildID)
	if err != nil {
		return fmt.Errorf("failed to list guild members: %v", err.Error())
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isNFTRoles[roleID] {
				if rolesToAdd[[2]string{member.User.ID, roleID}] {
					delete(rolesToAdd, [2]string{member.User.ID, roleID})
					continue
				}

				err = c.entity.RemoveGuildMemberRole(guildID, member.User.ID, roleID)
				if err != nil {
					c.log.Errorf(err, "cannot remove guild member role %s of %s - guild %s", roleID, member.User.ID, guildID)
				}
			}
		}
	}

	for roleToAdd := range rolesToAdd {
		err = c.entity.AddGuildMemberRole(guildID, roleToAdd[0], roleToAdd[1])
		if err != nil {
			c.log.Errorf(err, "cannot add guild member %s role %s - guild %s", roleToAdd[0], roleToAdd[1], guildID)
		}
	}

	return nil
}
