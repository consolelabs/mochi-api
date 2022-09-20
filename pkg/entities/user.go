package entities

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
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

	if user.InDiscordWalletAddress.String == "" {
		if err = e.generateInDiscordWallet(user); err != nil {
			err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
			return nil, err
		}
	}

	res := &response.User{
		ID:                     user.ID,
		Username:               user.Username,
		InDiscordWalletAddress: &user.InDiscordWalletAddress.String,
		InDiscordWalletNumber:  &user.InDiscordWalletNumber.Int64,
		GuildUsers:             guildUsers,
		NrOfJoin:               user.NrOfJoin,
	}
	return res, nil
}

func (e *Entity) GetUserCurrentGMStreak(discordID, guildID string) (*model.DiscordUserGMStreak, int, error) {
	streak, err := e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID(discordID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get user's gm streak: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		return nil, http.StatusBadRequest, fmt.Errorf("user has no gm streak")
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
	expireTime := streak.LastStreakDate.Add(time.Hour * 13)
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
		expireTime = log.LatestUpvoteTime.Add(time.Hour * 13)
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

func (e *Entity) GetUpvoteLeaderboardByStreak() ([]model.DiscordUserUpvoteStreak, error) {
	streaks, err := e.repo.DiscordUserUpvoteStreak.GetTopByStreak()
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetUpvoteLeaderboardByStreak] fail to get upvote leaderboard by streak")
		return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
	}
	return streaks, nil
}

func (e *Entity) GetUpvoteLeaderboardByTotal() ([]model.DiscordUserUpvoteStreak, error) {
	streaks, err := e.repo.DiscordUserUpvoteStreak.GetTopByTotal()
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetUpvoteLeaderboardByTotal] fail to get upvote leaderboard by total")
		return nil, fmt.Errorf("failed to get upvote leaderboard: %v", err)
	}
	return streaks, nil
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
	} else {
		e.svc.Discord.SendLevelUpMessage(latestUserXP.Guild.LogChannel, role, res)
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

func (e *Entity) GetTopUsers(guildID, userID string, limit, page int) (*response.TopUser, error) {
	offset := page * limit
	leaderboard, err := e.repo.GuildUserXP.GetTopUsers(guildID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range leaderboard {
		item := &leaderboard[i]
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

	return &response.TopUser{
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

func (e *Entity) SendGiftXp(req request.GiftXPRequest) (*response.HandleUserActivityResponse, error) {
	res, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		UserID:    req.UserDiscordID,
		Action:    "gifted",
		CustomXP:  int64(req.XPAmount),
	})
	if err != nil {
		e.log.Errorf(err, "[SendGiftXp] - HandleUserActivities failed: %v %v %v", req.GuildID, req.UserDiscordID, req.XPAmount)
		return nil, err
	}

	return res, nil
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
		if err := e.repo.GuildUsers.Create(&model.GuildUser{
			GuildID:   invitee.GuildID,
			UserID:    invitee.User.ID,
			Nickname:  invitee.Nick,
			InvitedBy: res.InviterID,
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

		// create invite history
		inviteType := model.INVITE_TYPE_NORMAL
		if inviter == nil {
			inviteType = model.INVITE_TYPE_LEFT
		}

		// TODO: Can't find age of user now
		// if time.Now().Unix()-invit < 60*60*24*3 {
		// 	inviteType = model.INVITE_TYPE_FAKE
		// }

		if err := e.repo.InviteHistories.Create(&model.InviteHistory{
			GuildID:   invitee.GuildID,
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
	switch err {
	case gorm.ErrRecordNotFound:
		u.ID = discordID
		return e.createNewUser(u)
	case nil:
		return e.createNewUser(u)
	default:
		e.log.Fields(logger.Fields{"discord_id": discordID}).Error(err, "[entity.GetOneOrUpsertUser] repo.Users.GetOne() failed")
		return nil, err
	}
}

func (e *Entity) createNewUser(u *model.User) (*model.User, error) {
	dcUser, err := e.discord.User(u.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"discord_id": u.ID}).Error(err, "[entity.createNewUser] discord.User() failed")
		return nil, err
	}

	switch {
	case u.InDiscordWalletAddress.String == "":
		u.Username = dcUser.Username
		if err := e.generateInDiscordWallet(u); err != nil {
			e.log.Fields(logger.Fields{"user": u}).Error(err, "[entity.createNewUser] generateInDiscordWallet() failed")
			return nil, err
		}
		return u, nil
	case u.Username != dcUser.Username:
		u.Username = dcUser.Username
		if err := e.repo.Users.Upsert(u); err != nil {
			e.log.Fields(logger.Fields{"user": u}).Error(err, "[entity.generateInDiscordWallet] repo.Users.Upsert() failed")
			return nil, err
		}
		return u, nil
	default:
		return u, nil
	}
}
