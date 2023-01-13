package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
)

type updateUserTokenRoles struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserTokenRolesJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &updateUserTokenRoles{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *updateUserTokenRoles) Run() error {
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
		if err := job.updateTokenRoles(guild.ID); err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "Run failed")
		}
	}

	return nil
}

func (job *updateUserTokenRoles) updateTokenRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateTokenRoles] starting")
	trConfigs, err := job.entity.ListGuildTokenRoles(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRoles] entity.ListGuildTokenRoles failed")
		return err
	}

	if len(trConfigs) == 0 {
		l.Info("[updateTokenRoles] entity.ListGuildTokenRoles - no data found")
		return nil
	}

	isTokenRoles := make(map[string]bool)
	for _, trConfig := range trConfigs {
		isTokenRoles[trConfig.RoleID] = true
	}

	members, err := job.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRole] entity.ListGuildMembers failed")
		return err
	}

	rolesToAdd, err := job.entity.ListMemberTokenRolesToAdd(trConfigs, guildID)
	if err != nil {
		l.Error(err, "[updateTokenRole] entity.ListMemberNFTRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isTokenRoles[roleID] {
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
					gMemberRoleLog.Infof("[updateTokenRoles] entity.RemoveGuildMemberRole failed: %v", err)
					continue
				}
				if err != nil {
					gMemberRoleLog.Error(err, "[updateTokenRoles] entity.RemoveGuildMemberRole failed")
					continue
				}
				gMemberRoleLog.Info("[updateTokenRoles] entity.RemoveGuildMemberRole executed successfully")
			}
		}
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRole] entity.GetGuild failed")
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
			gMemberRoleLog.Infof("[updateTokenRole] entity.AddGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			gMemberRoleLog.Error(err, "[updateTokenRole] entity.AddGuildMemberRole failed")
			continue
		}

		// send logs to moderation channel
		gMemberRoleLog.Info("[updateTokenRole] entity.AddGuildMemberRole executed successfully")
		err := job.service.Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "nft-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Info("[updateTokenRole] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}

func (job *updateUserTokenRoles) userBalances() (map[string]string, error) {

	return nil, nil
}
