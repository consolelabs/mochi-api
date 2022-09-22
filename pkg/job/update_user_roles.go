package job

import (
	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
)

type updateUserRoles struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserRolesJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &updateUserRoles{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *updateUserRoles) Run() error {
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
		if err := job.updateLevelRoles(guild.ID); err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "Run failed")
		}

		if err := job.updateNFTRoles(guild.ID); err != nil {
			job.log.Fields(logger.Fields{"guildId": guild.ID}).Error(err, "Run failed")
		}
	}

	return nil
}

func (job *updateUserRoles) updateLevelRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateLevelRoles] starting")
	lrConfigs, err := job.entity.GetGuildLevelRoleConfigs(guildID)
	if err != nil {
		l.Error(err, "[updateLevelRoles] entity.GetGuildLevelRoleConfigs failed")
		return err
	}

	if len(lrConfigs) == 0 {
		l.Info("[updateLevelRoles] entity.GetGuildLevelRoleConfigs - no data found")
		return nil
	}

	userXPs, err := job.entity.GetGuildUserXPs(guildID)
	if err != nil {
		l.Error(err, "[updateLevelRoles] entity.GetGuildUserXPs failed")
		return err
	}
	if len(userXPs) == 0 {
		l.Info("[updateLevelRoles] entity.GetGuildUserXPs - no data found")
		return nil
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateLevelRoles] entity.GetGuild failed")
		return err
	}

	rolesToAdd := make(map[string]string)
	rolesToRemove := make(map[string]string)
	for _, userXP := range userXPs {
		var member *discordgo.Member
		var jobErr error
		err := util.RetryRequest(func() error {
			member, jobErr = job.entity.GetGuildMember(guildID, userXP.UserID)
			return jobErr
		})
		if util.IsAcceptableErr(err) {
			job.log.Fields(logger.Fields{
				"userId":  userXP.UserID,
				"guildId": guildID,
			}).Infof("[updateLevelRoles] user is no longer guild member: %v", err)
			continue
		}
		if err != nil {
			job.log.Fields(logger.Fields{
				"userId":  userXP.UserID,
				"guildId": guildID,
			}).Error(err, "[updateLevelRoles] entity.GetGuildMember failed")
			continue
		}

		userLevelRole, err := job.entity.GetUserRoleByLevel(guildID, userXP.Level)
		switch {
		case err == gorm.ErrRecordNotFound:
			job.log.Fields(logger.Fields{
				"level":   userXP.Level,
				"guildId": guildID,
			}).Info("[updateLevelRoles] entity.GetUserRoleByLevel no role found")
			continue
		case err != nil && err != gorm.ErrRecordNotFound:
			job.log.Fields(logger.Fields{
				"level":   userXP.Level,
				"guildId": guildID,
			}).Error(err, "[updateLevelRoles] entity.GetUserRoleByLevel failed")
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

	for userID, roleID := range rolesToRemove {
		err := util.RetryRequest(func() error {
			return job.entity.RemoveGuildMemberRole(guildID, userID, roleID)
		})
		if util.IsAcceptableErr(err) {
			job.log.Fields(logger.Fields{
				"guildId": guildID,
				"userId":  userID,
				"roleId":  roleID,
			}).Infof("[updateLevelRoles] entity.RemoveGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId": guildID,
				"userId":  userID,
				"roleId":  roleID,
			}).Error(err, "[updateLevelRoles] entity.RemoveGuildMemberRole failed")
			continue
		}
		job.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		}).Info("[updateLevelRoles] entity.RemoveGuildMemberRole executed successfully")
	}

	for userID, roleID := range rolesToAdd {
		err := util.RetryRequest(func() error {
			return job.entity.AddGuildMemberRole(guildID, userID, roleID)
		})
		if util.IsAcceptableErr(err) {
			job.log.Fields(logger.Fields{
				"guildId": guildID,
				"userId":  userID,
				"roleId":  roleID,
			}).Infof("[updateLevelRoles] entity.AddGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId": guildID,
				"userId":  userID,
				"roleId":  roleID,
			}).Error(err, "[updateLevelRoles] entity.AddGuildMemberRole failed")
			continue
		}

		job.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		}).Info("[updateLevelRoles] entity.AddGuildMemberRole executed successfully")

		// send logs to moderation channel
		err = job.service.Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "level-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Info("[updateLevelRoles] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}

func (job *updateUserRoles) updateNFTRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateNFTRoles] starting")
	hrConfigs, err := job.entity.ListGuildNFTRoleConfigs(guildID)
	if err != nil {
		l.Error(err, "[updateNFTRoles] entity.ListGuildNFTRoleConfigs failed")
		return err
	}

	if len(hrConfigs) == 0 {
		l.Info("[updateNFTRoles] entity.ListGuildNFTRoleConfigs - no data found")
		return nil
	}

	isNFTRoles := make(map[string]bool)
	for _, hrConfig := range hrConfigs {
		isNFTRoles[hrConfig.RoleID] = true
	}

	members, err := job.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[updateNFTRoles] entity.ListGuildMembers failed")
		return err
	}

	rolesToAdd, err := job.entity.ListMemberNFTRolesToAdd(hrConfigs, guildID)
	if err != nil {
		l.Error(err, "[updateNFTRoles] entity.ListMemberNFTRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isNFTRoles[roleID] {
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
					gMemberRoleLog.Infof("[updateNFTRoles] entity.RemoveGuildMemberRole failed: %v", err)
					continue
				}
				if err != nil {
					gMemberRoleLog.Error(err, "[updateNFTRoles] entity.RemoveGuildMemberRole failed")
					continue
				}
				gMemberRoleLog.Info("[updateNFTRoles] entity.RemoveGuildMemberRole executed successfully")
			}
		}
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateNFTRoles] entity.GetGuild failed")
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
			gMemberRoleLog.Infof("[updateNFTRoles] entity.AddGuildMemberRole failed: %v", err)
			continue
		}
		if err != nil {
			gMemberRoleLog.Error(err, "[updateNFTRoles] entity.AddGuildMemberRole failed")
			continue
		}

		// send logs to moderation channel
		gMemberRoleLog.Info("[updateNFTRoles] entity.AddGuildMemberRole executed successfully")
		err := job.service.Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "nft-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Info("[updateNFTRoles] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}
