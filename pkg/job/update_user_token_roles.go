package job

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

type updateUserTokenRoles struct {
	entity *entities.Entity
	svc    *service.Service
	log    logger.Logger
	sentry *sentry.Client
	opts   *UpdateUserTokenRolesOptions
}

type UpdateUserTokenRolesOptions struct {
	// GuildID is the guild ID to update token roles
	GuildID string
	//  RolesToRemove is a list of roles to remove from users when using cmd "/tokenrole remove"
	RolesToRemove []string
}

func NewUpdateUserTokenRolesJob(e *entities.Entity, sentry *sentry.Client, opts *UpdateUserTokenRolesOptions) Job {
	if opts == nil {
		opts = &UpdateUserTokenRolesOptions{}
	}
	return &updateUserTokenRoles{
		entity: e,
		sentry: sentry,
		svc:    e.GetSvc(),
		log:    e.GetLogger(),
		opts:   opts,
	}
}

func (job *updateUserTokenRoles) Run() error {
	guildIDs := []string{}
	var err error

	switch {
	case job.opts.GuildID != "":
		guildIDs = append(guildIDs, job.opts.GuildID)
	default:
		guildIDs, err = job.entity.ListTokenRoleConfigGuildIds()
		if err != nil {
			job.log.Error(err, "entity.ListTokenRoleConfigGuildIds failed")
			job.captureSentry(fmt.Sprintf("entity.ListTokenRoleConfigGuildIds() failed: %v", err), nil)
			return err
		}
	}

	for _, guildId := range guildIDs {
		_, err := job.entity.GetGuildById(guildId)
		if err != nil {
			job.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "entity.GetGuildById failed")
			job.captureSentry(fmt.Sprintf("entity.GetGuildById() failed: %v", err), map[string]interface{}{
				"guildID": guildId,
			})
			continue
		}
		if err := job.updateTokenRoles(guildId); err != nil {
			job.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "Run failed")
			job.captureSentry(fmt.Sprintf("updateTokenRoles failed: %v", err), map[string]interface{}{
				"guildID": guildId,
			})
		}
	}

	return nil
}

func (job *updateUserTokenRoles) updateTokenRoles(guildID string) error {
	l := job.log.Fields(logger.Fields{"guildId": guildID})
	l.Info("[updateTokenRoles] starting...")

	trConfigs, err := job.entity.ListGuildTokenRoles(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRoles] entity.ListGuildTokenRoles failed")
		return err
	}

	if len(trConfigs) == 0 {
		l.Info("[updateTokenRoles] entity.ListGuildTokenRoles - no data found")
		return nil
	}

	// we only manage discord roles that are in db
	isTokenRoles := make(map[string]bool)
	for _, trConfig := range trConfigs {
		isTokenRoles[trConfig.RoleID] = true
	}

	// because we removed role from db before fetching them again, we need to keep track of roles to remove
	for _, roleID := range job.opts.RolesToRemove {
		isTokenRoles[roleID] = true
	}

	members, err := job.entity.ListGuildMembers(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRole] entity.ListGuildMembers failed")
		return err
	}

	rolesToAdd, err := job.listMemberTokenRolesToAdd(guildID, trConfigs, members)
	if err != nil {
		l.Error(err, "[updateTokenRole] job.listMemberTokenRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if !isTokenRoles[roleID] {
				continue
			}

			key := [2]string{member.User.ID, roleID}
			valid, ok := rolesToAdd[key]
			// if error occurs while fetching balance -> skip
			if ok && !valid {
				continue
			}

			// if user already has the role -> no need to add and skip removing
			if ok && valid {
				delete(rolesToAdd, key)
				continue
			}

			// if not a role to add -> remove
			gMemberRoleLog := job.log.Fields(logger.Fields{
				"guildId": guildID,
				"userId":  member.User.ID,
				"roleId":  roleID,
			})
			err = job.entity.RemoveGuildMemberRole(guildID, member.User.ID, roleID)
			if err != nil {
				gMemberRoleLog.Error(err, "[updateTokenRoles] entity.RemoveGuildMemberRole failed")
				continue
			}
			gMemberRoleLog.Info("[updateTokenRoles] entity.RemoveGuildMemberRole executed successfully")
		}
	}

	guild, err := job.entity.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[updateTokenRole] entity.GetGuild failed")
		return err
	}

	for roleToAdd, valid := range rolesToAdd {
		// if error occurs while fetching balance -> skip
		if !valid {
			continue
		}
		userID := roleToAdd[0]
		roleID := roleToAdd[1]
		gMemberRoleLog := job.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		})
		err = job.entity.AddGuildMemberRole(guildID, userID, roleID)
		if err != nil {
			gMemberRoleLog.Error(err, "[updateTokenRole] entity.AddGuildMemberRole failed")
			// if role is not found in discord, remove it from db
			if util.IsRoleNotFoundErr(err.Error()) {
				gMemberRoleLog.Infof("[updateTokenRole] entity.AddGuildMemberRole - remove role from db")
				for _, trConfig := range trConfigs {
					if trConfig.RoleID == roleID {
						job.entity.RemoveGuildTokenRole(trConfig.ID)
						break
					}
				}
			}
			continue
		}

		// send logs to moderation channel
		gMemberRoleLog.Info("[updateTokenRole] entity.AddGuildMemberRole executed successfully")
		err := job.entity.GetSvc().Discord.SendUpdateRolesLog(guildID, guild.LogChannel, userID, roleID, "nft-role")
		if err != nil {
			job.log.Fields(logger.Fields{
				"guildId":   guildID,
				"channelId": guild.LogChannel,
				"roleId":    roleID,
			}).Error(err, "[updateTokenRole] service.Discord.SendUpdateRolesLog failed")
			continue
		}
	}
	return nil
}

func (job *updateUserTokenRoles) listMemberTokenRolesToAdd(guildID string, cfgs []model.GuildConfigTokenRole, members []*discordgo.Member) (map[[2]string]bool, error) {
	tokens, err := job.entity.ListAllConfigTokens(guildID)
	if err != nil {
		job.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Job.UpdateUserTokenRoles] entity.ListAllConfigTokens() failed")
		return nil, err
	}
	userBals := make(map[struct {
		UserDiscordID string
		TokenID       int
	}]*big.Int)

	discordIds := []string{}
	for _, mem := range members {
		discordIds = append(discordIds, mem.User.ID)
	}
	logrus.WithField("discordIds", discordIds).Info("[Job.UpdateUserTokenRoles] getting profiles with discord Ids")
	profiles, err := job.svc.MochiProfile.GetByDiscordIds(discordIds)
	if err != nil {
		logrus.Error(err, "[Job.UpdateUserTokenRoles] service.MochiProfile.GetByDiscordIds() failed")
		return nil, err
	}

	for _, profile := range profiles {
		userDiscordID := ""
		for _, acc := range profile.AssociatedAccounts {
			if acc.Platform == mochiprofile.PlatformDiscord {
				userDiscordID = acc.PlatformIdentifier
				break
			}
		}
		for _, token := range tokens {
			bal, err := job.entity.CalculateTokenBalance(int64(token.ChainID), token.Address, profile)
			if err != nil {
				job.log.Error(err, "[Job.UpdateUserTokenRoles] entity.CalculateTokenBalance() failed")
				userBals[struct {
					UserDiscordID string
					TokenID       int
				}{UserDiscordID: userDiscordID, TokenID: token.ID}] = nil
				continue
			}
			userBals[struct {
				UserDiscordID string
				TokenID       int
			}{UserDiscordID: userDiscordID, TokenID: token.ID}] = bal
		}
	}

	// rolesToAdd: key = [userID, roleID] | value = valid balance (no error)
	rolesToAdd := make(map[[2]string]bool)
	for _, mem := range members {
		for _, cfg := range cfgs {
			userBal := userBals[struct {
				UserDiscordID string
				TokenID       int
			}{UserDiscordID: mem.User.ID, TokenID: cfg.TokenID}]
			// cannot fetch user balance
			if userBal == nil {
				rolesToAdd[[2]string{mem.User.ID, cfg.RoleID}] = false
				continue
			}
			decimalsBigFloat := new(big.Float).SetInt(math.BigPow(10, int64(cfg.Token.Decimals)))
			requiredAmountBigFloat := new(big.Float).Mul(big.NewFloat(cfg.RequiredAmount), decimalsBigFloat)
			requiredAmount := new(big.Int)
			requiredAmountBigFloat.Int(requiredAmount)
			if userBal.Cmp(requiredAmount) != -1 {
				rolesToAdd[[2]string{mem.User.ID, cfg.RoleID}] = true
			}
		}
	}

	return rolesToAdd, nil
}

func (j *updateUserTokenRoles) captureSentry(message string, data map[string]interface{}) {
	scope := sentry.NewScope()
	scope.SetLevel(sentry.LevelError)
	event := sentry.NewEvent()
	event.Level = sentry.LevelError
	event.Message = message
	event.Extra = data
	j.sentry.CaptureEvent(event, &sentry.EventHint{
		Data:              data,
		EventID:           message,
		OriginalException: errors.New(message),
	}, scope)
}
