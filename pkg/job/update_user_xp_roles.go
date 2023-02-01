package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
)

type updateUserXPRoles struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserXPRolesJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &updateUserXPRoles{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *updateUserXPRoles) Run() error {
	guilds, err := job.entity.GetGuilds()
	if err != nil {
		job.log.Error(err, "entity.GetGuilds failed")
		return err
	}

	for _, guild := range guilds.Data {
		_, err := job.entity.GetGuildById(guild.ID)
		if util.IsAcceptableErr(err) {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Infof("entity.GetGuildById - bot has no permission or access to this guild: %v", err)
			continue
		}
		if err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "entity.GetGuildById failed")
			continue
		}
		if err := job.updateXPRoles(guild.ID); err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "Run failed")
		}
	}

	return nil
}

func (job *updateUserXPRoles) updateXPRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateXPRoles] starting")
	xpConfigs, err := job.entity.ListGuildXPRoles(guildID)
	if err != nil {
		l.Error(err, "[updateXPRoles] entity.ListGuildXPRoles failed")
		return err
	}

	if len(xpConfigs) == 0 {
		l.Info("[updateXPRoles] entity.ListGuildXPRoles - no data found")
		return nil
	}

	isXPRoles := make(map[string]bool)
	for _, xpConfig := range xpConfigs {
		isXPRoles[xpConfig.RoleID] = true
	}

	members, err := job.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[updateXPRole] entity.ListGuildMembers failed")
		return err
	}

	rolesToAdd, err := job.entity.ListMemberXPRolesToAdd(xpConfigs, guildID)
	if err != nil {
		l.Error(err, "[updateXPRole] entity.ListMemberNFTRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isXPRoles[roleID] {
				if rolesToAdd[[2]string{member.User.ID, roleID}] {
					delete(rolesToAdd, [2]string{member.User.ID, roleID})
					continue
				}

				gMemberRoleLog := job.log.Fields(logger.Fields{
					"guildId": guildID,
					"userId":  member.User.ID,
					"roleId":  roleID,
				})
				err = job.entity.RemoveGuildMemberRole(guildID, member.User.ID, roleID)
				if util.IsAcceptableErr(err) {
					gMemberRoleLog.Infof("[updateXPRoles] entity.RemoveGuildMemberRole failed: %v", err)
					continue
				}
				if err != nil {
					gMemberRoleLog.Error(err, "[updateXPRoles] entity.RemoveGuildMemberRole failed")
					continue
				}
				gMemberRoleLog.Info("[updateXPRoles] entity.RemoveGuildMemberRole executed successfully")
			}
		}
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateXPRole] entity.GetGuild failed")
		return err
	}

	for roleToAdd := range rolesToAdd {
		userID := roleToAdd[0]
		roleID := roleToAdd[1]
		gMemberRoleLog := job.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		})
		err = job.entity.AddGuildMemberRole(guildID, userID, roleID)
		if util.IsAcceptableErr(err) {
			gMemberRoleLog.Infof("[updateXPRole] entity.AddGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			gMemberRoleLog.Error(err, "[updateXPRole] entity.AddGuildMemberRole failed")
			continue
		}

		// send logs to moderation channel
		gMemberRoleLog.Info("[updateXPRole] entity.AddGuildMemberRole executed successfully")
		err := job.service.Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "nft-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Info("[updateXPRole] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}
