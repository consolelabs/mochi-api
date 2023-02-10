package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
)

type updateUserMixRoles struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserMixRolesJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &updateUserMixRoles{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *updateUserMixRoles) Run() error {
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
		if err := job.updateMixRoles(guild.ID); err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "Run failed")
		}
	}

	return nil
}

func (job *updateUserMixRoles) updateMixRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateMixRoles] starting")
	MixConfigs, err := job.entity.ListGuildMixRoles(guildID)
	if err != nil {
		l.Error(err, "[updateMixRoles] entity.ListGuildMixRoles failed")
		return err
	}

	if len(MixConfigs) == 0 {
		l.Info("[updateMixRoles] entity.ListGuildMixRoles - no data found")
		return nil
	}

	isMixRoles := make(map[string]bool)
	for _, MixConfig := range MixConfigs {
		isMixRoles[MixConfig.RoleID] = true
	}

	members, err := job.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[updateMixRole] entity.ListGuildMembers failed")
		return err
	}

	rolesToAdd, err := job.entity.ListMemberMixRolesToAdd(MixConfigs, guildID)
	if err != nil {
		l.Error(err, "[updateMixRole] entity.ListMemberNFTRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isMixRoles[roleID] {
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
					gMemberRoleLog.Infof("[updateMixRoles] entity.RemoveGuildMemberRole failed: %v", err)
					continue
				}
				if err != nil {
					gMemberRoleLog.Error(err, "[updateMixRoles] entity.RemoveGuildMemberRole failed")
					continue
				}
				gMemberRoleLog.Info("[updateMixRoles] entity.RemoveGuildMemberRole executed successfully")
			}
		}
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateMixRole] entity.GetGuild failed")
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
			gMemberRoleLog.Infof("[updateMixRole] entity.AddGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			gMemberRoleLog.Error(err, "[updateMixRole] entity.AddGuildMemberRole failed")
			continue
		}

		// send logs to moderation channel
		gMemberRoleLog.Info("[updateMixRole] entity.AddGuildMemberRole executed successfully")
		err := job.service.Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "nft-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Info("[updateMixRole] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}
