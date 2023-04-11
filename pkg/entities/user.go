package entities

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) CreateUser(req request.CreateUserRequest) error {

	user := &model.User{
		ID:       req.ID,
		Username: req.Username,
		GuildUsers: []*model.GuildUser{
			{
				GuildID:   req.GuildID,
				UserID:    req.ID,
				Nickname:  req.Nickname,
				InvitedBy: req.InvitedBy,
			},
		},
	}

	if err := e.repo.Users.Upsert(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (e *Entity) GetUser(discordID string) (*response.User, error) {
	user, err := e.repo.Users.GetOne(discordID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	guildUsers := []*response.GetGuildUserResponse{}
	for _, guildUser := range user.GuildUsers {
		guildUsers = append(guildUsers, &response.GetGuildUserResponse{
			GuildID:   guildUser.GuildID,
			UserID:    guildUser.UserID,
			Nickname:  guildUser.Nickname,
			InvitedBy: guildUser.InvitedBy,
		})
	}

	res := &response.User{
		ID:         user.ID,
		Username:   user.Username,
		GuildUsers: guildUsers,
		NrOfJoin:   user.NrOfJoin,
	}
	return res, nil
}

func (e *Entity) GetUserCurrentGMStreak(discordID, guildID string) (*model.DiscordUserGMStreak, int, error) {
	streak, err := e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID(discordID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get user's gm streak: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		return &model.DiscordUserGMStreak{}, http.StatusOK, nil
	}

	return streak, http.StatusOK, nil
}

func (e *Entity) GetAllGMStreak() ([]model.DiscordUserGMStreak, error) {
	streaks, err := e.repo.DiscordUserGMStreak.GetAll()
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetAllGMStreak] fail to get all gm streaks")
		return nil, fmt.Errorf("failed to get all gm streaks: %v", err)
	}
	return streaks, nil
}

func (e *Entity) UpsertBatchGMStreak(streaks []model.DiscordUserGMStreak) error {
	err := e.repo.DiscordUserGMStreak.UpsertBatch(streaks)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.UpsertOneGMStreak] fail to get all gm streaks")
		return fmt.Errorf("failed to upsert gm streaks: %v", err)
	}
	return nil
}

func (e *Entity) GetUserCurrentUpvoteStreak(discordID string) (*response.GetUserCurrentUpvoteStreakResponse, int, error) {
	streak, err := e.repo.DiscordUserUpvoteStreak.GetByDiscordID(discordID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetUserCurrentUpvoteStreak] fail to get user upvote streak")
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get user's upvote streak: %v", err)
	}
	if err == gorm.ErrRecordNotFound {
		e.log.Info("[e.GetUserCurrentUpvoteStreak] user upvote streak empty")
		return nil, http.StatusOK, nil
	}

	var resetTime, topggTime, dcBotTime float64 = 0, 0, 0
	expireTime := streak.LastStreakDate.Add(time.Hour * 12)
	now := time.Now()
	currTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
	if currTime.Before(expireTime) {
		resetTime = util.MinuteLeftUntil(currTime, expireTime)
	}

	upvoteLogs, err := e.repo.DiscordUserUpvoteLog.GetByDiscordID(discordID)
	if err != nil {
		e.log.Info("[e.GetUserCurrentUpvoteStreak] user first time upvote")
	}
	for _, log := range upvoteLogs {
		expireTime = log.LatestUpvoteTime.Add(time.Hour * 12)
		switch log.Source {
		case "topgg":
			topggTime = util.MinuteLeftUntil(currTime, expireTime)
		case "discordbotlist":
			dcBotTime = util.MinuteLeftUntil(currTime, expireTime)
		}

	}

	return &response.GetUserCurrentUpvoteStreakResponse{
		UserID:                  streak.DiscordID,
		ResetTime:               resetTime,
		ResetTimeTopGG:          topggTime,
		ResetTimeDiscordBotList: dcBotTime,
		SteakCount:              streak.StreakCount,
		TotalCount:              streak.TotalCount,
		LastStreakTime:          streak.LastStreakDate,
	}, http.StatusOK, nil
}

func (e *Entity) GetUpvoteLeaderboard(by, guildId string) ([]model.DiscordUserUpvoteStreak, error) {
	if guildId == "" {
		switch by {
		case "total":
			streaks, err := e.repo.DiscordUserUpvoteStreak.GetTopByTotal()
			if err != nil && err != gorm.ErrRecordNotFound {
				e.log.Errorf(err, "[e.GetUpvoteLeaderboard] fail to get upvote global leaderboard by total")
				return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
			}
			return streaks, nil
		case "streak":
			streaks, err := e.repo.DiscordUserUpvoteStreak.GetTopByStreak()
			if err != nil && err != gorm.ErrRecordNotFound {
				e.log.Errorf(err, "[e.GetUpvoteLeaderboard] fail to get upvote global leaderboard by streak")
				return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
			}
			return streaks, nil
		default:
			e.log.Infof("[e.GetUpvoteLeaderboard] invalid query string by=%s", by)
			return nil, fmt.Errorf("invalid query string")
		}
	} else {
		switch by {
		case "total":
			streaks, err := e.repo.DiscordUserUpvoteStreak.GetGuildTopByTotal(guildId)
			if err != nil && err != gorm.ErrRecordNotFound {
				e.log.Errorf(err, "[e.GetUpvoteLeaderboard] fail to get upvote guild leaderboard by total")
				return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
			}
			return streaks, nil
		case "streak":
			streaks, err := e.repo.DiscordUserUpvoteStreak.GetGuildTopByStreak(guildId)
			if err != nil && err != gorm.ErrRecordNotFound {
				e.log.Errorf(err, "[e.GetUpvoteLeaderboard] fail to get upvote guild leaderboard by streak")
				return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
			}
			return streaks, nil
		default:
			e.log.Infof("[e.GetUpvoteLeaderboard] invalid query string by=%s", by)
			return nil, fmt.Errorf("invalid query string")
		}
	}

}

func (e *Entity) GetAllUpvoteStreak() ([]model.DiscordUserUpvoteStreak, error) {
	streaks, err := e.repo.DiscordUserUpvoteStreak.GetAll()
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetAllUpvoteStreak] fail to get all upvote streaks")
		return nil, fmt.Errorf("failed to get all upvote streaks: %v", err)
	}
	return streaks, nil
}

func (e *Entity) UpsertBatchUpvoteStreak(streak []model.DiscordUserUpvoteStreak) error {
	err := e.repo.DiscordUserUpvoteStreak.UpsertBatch(streak)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetAllUpvoteStreak] fail to get all upvote streaks")
		return fmt.Errorf("failed to upsert upvote streaks: %v", err)
	}
	return nil
}

func (e *Entity) HandleUserActivities(req *request.HandleUserActivityRequest) (*response.HandleUserActivityResponse, error) {
	userXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	earnedXP := int(req.CustomXP)
	if earnedXP == 0 {
		gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, req.Action)
		if err != nil {
			return nil, fmt.Errorf("failed to get guild config activity: %v", err.Error())
		}
		earnedXP = gActivityConfig.Activity.XP
	}

	if err := e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
		GuildID:      req.GuildID,
		UserID:       req.UserID,
		ActivityName: req.Action,
		EarnedXP:     earnedXP,
		CreatedAt:    req.Timestamp,
	}); err != nil {
		e.log.
			Fields(logger.Fields{"guildID": req.GuildID, "userID": req.UserID, "action": req.Action}).
			Error(err, "[Entity][HandleUserActivities] failed to create guild_user_activity_logs")
		return nil, err
	}

	latestUserXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil {
		return nil, err
	}

	res := &response.HandleUserActivityResponse{
		GuildID:      req.GuildID,
		ChannelID:    req.ChannelID,
		UserID:       req.UserID,
		Action:       req.Action,
		AddedXP:      earnedXP,
		CurrentXP:    latestUserXP.TotalXP,
		CurrentLevel: latestUserXP.Level,
		Timestamp:    req.Timestamp,
		LevelUp:      latestUserXP.Level > userXP.Level,
	}

	role, err := e.GetRoleByGuildLevelConfig(req.GuildID, req.UserID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildId": req.GuildID,
			"userId":  req.UserID,
		}).Errorf(err, "[HandleUserActivities] - SendLevelUpMessage failed")
	} else if res.LevelUp {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "userID": req.UserID}).Info("User leveled up")
		// get level up config
		config, err := e.repo.GuildConfigLevelUpMessage.GetByGuildId(req.GuildID)
		if err != nil && err != gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"guildId": req.GuildID}).Errorf(err, "[HandleUserActivities] - e.repo.GuildConfigLevelUpMessage.GetByGuildId failed")
			return nil, err
		}
		e.svc.Discord.SendLevelUpMessage(config, role, res)
	}
	return res, nil
}

func (e *Entity) InitGuildDefaultActivityConfigs(guildID string) error {
	activities, err := e.repo.Activity.GetDefaultActivities()
	if err != nil {
		return err
	}

	var configs []model.GuildConfigActivity
	for _, a := range activities {
		configs = append(configs, model.GuildConfigActivity{
			GuildID:    guildID,
			ActivityID: a.ID,
			Active:     true,
		})
	}

	return e.repo.GuildConfigActivity.UpsertMany(configs)
}

func (e *Entity) GetTopUsers(guildID, userID, query, sort string, limit, page int) (*response.TopUser, error) {
	offset := page * limit
	leaderboard, err := e.repo.GuildUserXP.GetTopUsers(guildID, query, sort, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range leaderboard {
		item := &leaderboard[i]

		if item.User != nil && len(item.User.GuildUsers) > 0 {
			memberInfo := item.User.GuildUsers[0]

			rolesByte := memberInfo.Roles

			roles := make([]string, 0)

			if err := json.Unmarshal(rolesByte, &roles); err != nil {
				return nil, err
			}

			item.User.GuildUsers[0].RoleSlice = roles
		}

		currentLevel, err := e.repo.ConfigXPLevel.GetNextLevel(item.TotalXP, false)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		nextLevel, err := e.repo.ConfigXPLevel.GetNextLevel(item.TotalXP, true)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		item.Progress = math.Min(float64(item.TotalXP-currentLevel.MinXP)/float64(nextLevel.MinXP-currentLevel.MinXP), 1)
		if nextLevel.Level == 0 {
			item.Progress = 1
		}
	}

	author, err := e.repo.GuildUserXP.GetOne(guildID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	total, err := e.repo.GuildUserXP.GetTotalTopUsersCount(guildID, query)
	if err != nil {
		return nil, err
	}

	return &response.TopUser{
		Metadata: response.PaginationResponse{
			Pagination: model.Pagination{
				Page: int64(page),
				Size: int64(limit),
			},
			Total: total,
		},
		Author:      author,
		Leaderboard: leaderboard,
	}, nil
}

func (e *Entity) GetGuildUserXPs(guildID string) ([]model.GuildUserXP, error) {
	return e.repo.GuildUserXP.GetByGuildID(guildID)
}

func (e *Entity) GetGuildMember(guildID, userID string) (*discordgo.Member, error) {
	member, err := e.discord.GuildMember(guildID, userID)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (e *Entity) ListGuildMembers(guildID string) ([]*discordgo.Member, error) {
	var afterID string
	res := make([]*discordgo.Member, 0)
	for {
		members, err := e.discord.GuildMembers(guildID, afterID, 100)
		if err != nil {
			return nil, err
		}
		res = append(res, members...)
		if len(members) < 100 {
			break
		}
		afterID = members[len(members)-1].User.ID
	}
	return res, nil
}

func (e *Entity) GetUserProfile(guildID, userID string) (*response.GetUserProfileResponse, error) {
	gUserXP, err := e.repo.GuildUserXP.GetOne(guildID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	currentLevel, err := e.repo.ConfigXPLevel.GetNextLevel(gUserXP.TotalXP, false)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	nextLevel, err := e.repo.ConfigXPLevel.GetNextLevel(gUserXP.TotalXP, true)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	progress := math.Min(float64(gUserXP.TotalXP-currentLevel.MinXP)/float64(nextLevel.MinXP-currentLevel.MinXP), 1)
	if nextLevel.Level == 0 {
		progress = 1
	}

	if gUserXP.Guild == nil {
		if gUserXP.Guild, err = e.repo.DiscordGuilds.GetByID(guildID); err != nil {
			return nil, err
		}
	}

	userWallet, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(userID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	userFactionXp, err := e.svc.Processor.GetUserFactionXp(userID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
		}).Error(err, "[e.svc.Processor.GetUserFactionXp] - get user faction xp from Processor failed")
		return nil, err
	}

	return &response.GetUserProfileResponse{
		ID:           userID,
		CurrentLevel: currentLevel,
		NextLevel:    nextLevel,
		GuildXP:      gUserXP.TotalXP,
		NrOfActions:  gUserXP.NrOfActions,
		Progress:     progress,
		Guild:        gUserXP.Guild,
		GuildRank:    gUserXP.GuildRank,
		UserWallet:   userWallet,
		UserFactionXps: &model.UserFactionXpsMapping{
			ImperialXp: userFactionXp.Data.NobilityXp,
			RebellioXp: userFactionXp.Data.FameXp,
			MerchantXp: userFactionXp.Data.LoyaltyXp,
			AcademyXp:  userFactionXp.Data.ReputationXp,
		},
	}, nil
}

func (e *Entity) ListAllWalletAddresses() ([]model.WalletAddress, error) {
	was, err := e.repo.UserWallet.ListWalletAddresses("evm")
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet addresses: %v", err.Error())
	}
	return was, nil
}

func (e *Entity) GetRoleByGuildLevelConfig(guildID, userID string) (string, error) {
	if e.discord == nil {
		return "", nil
	}
	configs, err := e.repo.GuildConfigLevelRole.GetByGuildID(guildID)
	if err != nil {
		return "", err
	}

	gMember, err := e.discord.GuildMember(guildID, userID)
	if err != nil {
		return "", err
	}
	if gMember.Roles == nil {
		return "", fmt.Errorf("Member %s of guild %s has no role", userID, guildID)
	}

	for _, cfg := range configs {
		for _, memRole := range gMember.Roles {
			if cfg.RoleID == memRole {
				return fmt.Sprintf("<@&%s>", cfg.RoleID), nil
			}
		}
	}

	return "", nil
}

func (e *Entity) HandleInviteTracker(inviter *discordgo.Member, invitee *discordgo.Member) (*response.HandleInviteHistoryResponse, error) {
	res := &response.HandleInviteHistoryResponse{}
	var guildID string

	if inviter != nil {
		// create inviter if not exist
		if _, err := e.GetOneOrUpsertUser(inviter.User.ID); err != nil {
			e.log.Fields(logger.Fields{"userID": inviter.User.ID, "username": inviter.User.Username}).
				Error(err, "[Entity][CreateInvite] failed to create user for inviter")
			return nil, err
		}
		if err := e.CreateGuildUserIfNotExists(inviter.GuildID, inviter.User.ID, inviter.Nick); err != nil {
			e.log.Fields(logger.Fields{"userID": inviter.User.ID, "username": inviter.User.Username}).
				Error(err, "[Entity][CreateInvite] GetOneOrUpsertUser() failed to create guild user for inviter")
			return nil, err
		}

		totalInvites, err := e.CountInviteHistoriesByGuildUser(inviter.GuildID, inviter.User.ID)
		if err != nil {
			e.log.Fields(logger.Fields{"inviterID": invitee.User.ID, "guildID": inviter.GuildID}).
				Error(err, "[Entity][CreateInvite] failed to count inviter invites")
			return nil, err
		}

		res.InvitesAmount = int(totalInvites)
		res.InviterID = inviter.User.ID
		res.IsBot = inviter.User.Bot
	}

	if invitee != nil {
		// create invitee if not exist
		user, err := e.GetOneOrUpsertUser(invitee.User.ID)
		if err != nil {
			e.log.Fields(logger.Fields{"userID": invitee.User.ID, "username": invitee.User.Username}).
				Error(err, "[Entity][CreateInvite] GetOneOrUpsertUser() failed to create user for invitee")
			return nil, err
		}

		roles, err := json.Marshal(invitee.Roles)
		if err != nil {
			e.log.Fields(logger.Fields{"userID": invitee.User.ID, "username": invitee.User.Username}).Error(err, "[entity.FetchAndSaveGuildMembers] json.Marshal() failed")
			return nil, err
		}

		if err := e.repo.GuildUsers.Create(&model.GuildUser{
			GuildID:   invitee.GuildID,
			UserID:    invitee.User.ID,
			Nickname:  invitee.Nick,
			InvitedBy: res.InviterID,
			Roles:     roles,
		}); err != nil {
			e.log.Fields(logger.Fields{
				"guildID":         invitee.GuildID,
				"inviteeID":       invitee.User.ID,
				"inviteeNickname": invitee.Nick,
				"inviterID":       res.InviterID,
			}).Error(err, "[Entity][CreateInvite] failed to create guild user for invitee")
			return nil, err
		}

		err = e.repo.Users.UpdateNrOfJoin(invitee.User.ID, user.NrOfJoin+1)
		if err != nil {
			e.log.Fields(logger.Fields{
				"guildID":   invitee.GuildID,
				"inviteeID": invitee.User.ID,
			}).Error(err, "[Entity][CreateInvite] failed to update nr_of_join of user")
			return nil, err
		}

		res.InviteeID = invitee.User.ID
		res.IsInviteeABot = invitee.User.Bot
		guildID = invitee.GuildID
	}

	// create invite history
	inviteType := model.INVITE_TYPE_NORMAL
	if inviter == nil {
		inviteType = model.INVITE_TYPE_LEFT
	}

	// TODO: Can't find age of user now
	// if time.Now().Unix()-invit < 60*60*24*3 {
	// 	inviteType = model.INVITE_TYPE_FAKE
	// }

	if res.InviteeID != "" && res.InviterID != "" && guildID != "" {
		if err := e.repo.InviteHistories.Create(&model.InviteHistory{
			GuildID:   guildID,
			UserID:    res.InviteeID,
			InvitedBy: res.InviterID,
			Type:      inviteType,
		}); err != nil {
			e.log.Fields(logger.Fields{
				"guildID":   invitee.GuildID,
				"userID":    res.InviteeID,
				"invitedBy": res.InviterID,
				"type":      inviteType,
			}).Error(err, "[Entity][CreateInvite] failed to create invite history")
			return nil, err
		}
	}

	return res, nil
}

func (e *Entity) GetOneOrUpsertUser(discordID string) (*model.User, error) {
	u, err := e.repo.Users.GetOne(discordID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"discord_id": discordID}).Error(err, "[entity.GetOneOrUpsertUser] repo.Users.GetOne() failed")
		return nil, err
	}
	u.ID = discordID
	err = e.repo.Users.UpsertMany([]model.User{*u})
	if err != nil {
		e.log.Fields(logger.Fields{"user": *u}).Error(err, "[entity.GetOneOrUpsertUser] repo.Users.UpsertMany() failed")
		return nil, err
	}
	return u, nil
}

func (e *Entity) GetUserDevice(deviceID string) (*response.UserDeviceResponse, error) {
	data, err := e.repo.DiscordUserDevice.GetByDeviceID(deviceID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"deviceID": deviceID}).Error(err, "[e.repo.DiscordUserDevice.GetByDeviceID] - failed to get user device")
		return nil, err
	}
	return &response.UserDeviceResponse{
		DeviceID:     data.ID,
		IosNotiToken: data.IosNotiToken,
	}, nil
}
func (e *Entity) UpsertUserDevice(req *request.UpsertUserDeviceRequest) error {
	return e.repo.DiscordUserDevice.UpsertOne(&model.DiscordUserDevice{
		ID:           req.DeviceID,
		IosNotiToken: req.IosNotiToken,
		UpdatedAt:    time.Now().UTC(),
	})
}
func (e *Entity) DeleteUserDevice(req *request.DeleteUserDeviceRequest) error {
	return e.repo.DiscordUserDevice.RemoveByDeviceID(req.DeviceID)
}

func (e *Entity) GetUserWalletByGuildIDAddress(guildID, address string) (*model.UserWallet, error) {
	uw, err := e.repo.UserWallet.GetOneByGuildIDAndAddress(guildID, address)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return uw, nil
}

func (e *Entity) TotalActiveUsers(guildId string) (*response.Metric, error) {
	discordGuilds, err := e.repo.DiscordGuilds.GetNonLeftGuilds()
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "[entities.TotalActiveUsers] - cannot get discord guilds")
		return nil, err
	}

	totalMembersAllGuilds := make([]response.MetricActiveUser, 0)
	for _, guild := range discordGuilds {
		var guildActiveUsers int64

		guildInfo, err := e.discord.GuildWithCounts(guild.ID)
		if err != nil {
			e.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "[entities.TotalActiveUsers] - failed to get total active users of current guild")
			return nil, err
		}
		guildActiveUsers += int64(guildInfo.ApproximateMemberCount)

		totalMembersAllGuilds = append(totalMembersAllGuilds, response.MetricActiveUser{
			GuildId:     guild.ID,
			ActiveUsers: guildActiveUsers,
		})
	}

	// sum total active users of all guilds
	var sumActiveUsers, currentGuildActiveUser int64
	for _, guildMember := range totalMembersAllGuilds {
		sumActiveUsers += guildMember.ActiveUsers
		if guildMember.GuildId == guildId {
			currentGuildActiveUser = guildMember.ActiveUsers
		}
	}

	return &response.Metric{
		TotalActiveUsers:  sumActiveUsers,
		ServerActiveUsers: currentGuildActiveUser,
	}, nil
}

func (e *Entity) TotalVerifiedWallets(guildId string) (*response.Metric, error) {
	totalVerfiedWallets, err := e.repo.DiscordWalletVerification.TotalVerifiedWallets()
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "[entities.TotalVerifiedWallets] - cannot get total verified wallets")
		return nil, err
	}
	guildVerifiedWallets, err := e.repo.DiscordWalletVerification.TotalVerifiedWalletsByGuildID(guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "[entities.TotalVerifiedWallets] - cannot get total verified wallets by guild id")
		return nil, err
	}

	return &response.Metric{
		TotalVerifiedWallets:  totalVerfiedWallets,
		ServerVerifiedWallets: guildVerifiedWallets,
	}, nil
}

func (e *Entity) FetchAndSaveGuildMembers(guildID string) (int, error) {
	roles, err := e.svc.Discord.GetGuildRoles(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entity.FetchAndSaveGuildMembers] entity.GetGuildRolesFromDiscord() failed")
		return 0, err
	}

	// Sort roles by position DESC
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Position > roles[j].Position
	})

	// Create a map of roleID to position
	rolePositionMap := make(map[string]int)
	for _, r := range roles {
		rolePositionMap[r.ID] = r.Position
	}

	members, err := e.GetGuildUsersFromDiscord(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entity.FetchAndSaveGuildMembers] entity.GetGuildUsersFromDiscord() failed")
		return 0, err
	}
	upsertUsersPayload := make([]model.User, 0, len(members))
	upsertGuildUsersPayload := make([]model.GuildUser, 0, len(members))
	for _, m := range members {
		upsertUsersPayload = append(upsertUsersPayload, model.User{
			ID:            m.User.ID,
			Username:      m.User.Username,
			Discriminator: m.User.Discriminator,
		})

		// Map each member role to proper position
		userRolePositionMap := make(map[string]int)
		for _, r := range m.Roles {
			userRolePositionMap[r] = rolePositionMap[r]
		}

		// Sort roles by position DESC
		sort.SliceStable(m.Roles, func(i, j int) bool {
			return userRolePositionMap[m.Roles[i]] > userRolePositionMap[m.Roles[j]]
		})

		roles, err := json.Marshal(m.Roles)
		if err != nil {
			e.log.Fields(logger.Fields{"guildID": guildID, "user": m.User.ID}).Error(err, "[entity.FetchAndSaveGuildMembers] json.Marshal() failed")
			return 0, err
		}

		upsertGuildUsersPayload = append(upsertGuildUsersPayload, model.GuildUser{
			UserID:   m.User.ID,
			Nickname: m.Nickname,
			GuildID:  guildID,
			Avatar:   m.Avatar,
			JoinedAt: m.JoinedAt,
			Roles:    roles,
		})
	}

	if err = e.repo.Users.UpsertMany(upsertUsersPayload); err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID, "users": len(upsertUsersPayload)}).Error(err, "[entity.FetchAndSaveGuildMembers] repo.Users.UpsertMany() failed")
		return 0, err
	}
	if err = e.repo.GuildUsers.UpsertMany(upsertGuildUsersPayload); err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID, "gUsers": len(upsertGuildUsersPayload)}).Error(err, "[entity.FetchAndSaveGuildMembers] repo.GuildUsers.UpsertMany() failed")
		return 0, err
	}
	return len(members), nil
}
