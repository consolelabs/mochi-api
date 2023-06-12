package job

import (
	"math/big"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
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
	guildIds, err := job.entity.ListTokenRoleConfigGuildIds()
	if err != nil {
		job.log.Error(err, "entity.ListTokenRoleConfigGuildIds failed")
		return err
	}

	for _, guildId := range guildIds {
		_, err := job.entity.GetGuildById(guildId)
		if util.IsAcceptableErr(err) {
			job.log.Fields(logger.Fields{"guildId": guildId}).Infof("entity.GetGuildById - bot has no permission or access to this guild: %v", err)
			continue
		}
		if err != nil {
			job.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "entity.GetGuildById failed")
			continue
		}
		if err := job.updateTokenRoles(guildId); err != nil {
			job.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "Run failed")
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

	rolesToAdd, err := job.listMemberTokenRolesToAdd(guildID, trConfigs, members)
	if err != nil {
		l.Error(err, "[updateTokenRole] job.listMemberTokenRolesToAdd failed")
		return err
	}

	for _, member := range members {
		for _, roleID := range member.Roles {
			if isTokenRoles[roleID] {
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

func (job *updateUserTokenRoles) listMemberTokenRolesToAdd(guildID string, cfgs []model.GuildConfigTokenRole, members []*discordgo.Member) (map[[2]string]bool, error) {
	tokens, err := job.entity.ListAllConfigTokens(guildID)
	if err != nil {
		job.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Job.UpdateUserTokenRoles] entity.ListAllConfigTokens() failed")
		return nil, err
	}
	userBals := make(map[struct {
		UserID  string
		TokenID int
	}]*big.Int)
	for _, mem := range members {
		for _, token := range tokens {
			bal, err := job.entity.CalculateTokenBalance(int64(token.ChainID), token.Address, mem.User.ID)
			if err != nil {
				job.log.Error(err, "[Job.UpdateUserTokenRoles] entity.CalculateTokenBalance() failed")
				userBals[struct {
					UserID  string
					TokenID int
				}{UserID: mem.User.ID, TokenID: token.ID}] = nil
				continue
			}
			userBals[struct {
				UserID  string
				TokenID int
			}{UserID: mem.User.ID, TokenID: token.ID}] = bal
		}
	}

	// rolesToAdd: key = [userID, roleID] | value = valid balance (no error)
	rolesToAdd := make(map[[2]string]bool)
	for _, mem := range members {
		for _, cfg := range cfgs {
			userBal := userBals[struct {
				UserID  string
				TokenID int
			}{UserID: mem.User.ID, TokenID: cfg.TokenID}]
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
