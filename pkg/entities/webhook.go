package entities

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) GuildLatestInvites(guildID string) ([]*discordgo.Invite, error) {
	invites, err := e.discord.GuildInvites(guildID)
	if err != nil {
		return nil, err
	}

	return invites, nil
}

func (e *Entity) GuildCachedInvites(guildID string) (invites map[string]string, err error) {
	return e.cache.HashGet(consts.CachePrefixInviteTracker + guildID)
}

func (e *Entity) GuildLatestVanityUses(guildID string, invites []*discordgo.Invite) (int, error) {
	guild, err := e.discord.Guild(guildID)
	if err != nil {
		return 0, err
	}

	for _, invite := range invites {
		if invite.Code == guild.VanityURLCode {
			return invite.Uses, nil
		}
	}

	return 0, nil
}

func (e *Entity) GuildCachedVanityUses(guildID string, invites map[string]string) (int, error) {
	guild, err := e.discord.Guild(guildID)
	if err != nil {
		return 0, err
	}

	if uses, ok := invites[guild.VanityURLCode]; ok {
		return strconv.Atoi(uses)
	}

	return 0, nil
}

func (e *Entity) SetGuildCacheInvites(guildID string, invites map[string]string) error {
	return e.cache.HashSet(consts.CachePrefixInviteTracker+guildID, invites, 0)
}

func (e *Entity) FindInviter(guildID string) (inviter *discordgo.Member, isVanity bool, err error) {
	latestInvites, err := e.GuildLatestInvites(guildID)
	if err != nil {
		return nil, false, err
	}
	cachedInvites, err := e.GuildCachedInvites(guildID)
	if err != nil {
		return nil, false, err
	}

	latestVanityUses, err := e.GuildLatestVanityUses(guildID, latestInvites)
	if err != nil {
		return nil, false, err
	}
	cachedVanityUses, err := e.GuildCachedVanityUses(guildID, cachedInvites)
	if err != nil {
		return nil, false, err
	}
	if latestVanityUses > cachedVanityUses {
		return nil, true, nil
	}

	var inviterID string
	for _, invite := range latestInvites {
		var cachedUses int64
		cachedUsesStr, ok := cachedInvites[invite.Code]
		if ok {
			cachedUses, err = strconv.ParseInt(cachedUsesStr, 10, 64)
			if err != nil {
				continue
			}
		}

		if invite.Uses > int(cachedUses) {
			inviterID = invite.Inviter.ID
			cachedInvites[invite.Code] = strconv.Itoa(invite.Uses)
			break
		}
	}

	if err := e.SetGuildCacheInvites(guildID, cachedInvites); err != nil {
		return nil, false, err
	}

	member, err := e.discord.GuildMember(guildID, inviterID)
	if err != nil {
		return nil, false, err
	}

	return member, false, nil
}

func (e *Entity) GetUserGlobalInviteCodes(guildID, userID string) ([]string, error) {
	resp := make([]string, 0)
	invites, err := e.discord.GuildInvites(guildID)
	if err != nil {
		return resp, err
	}

	for _, invite := range invites {
		if invite.Inviter.ID == userID &&
			invite.TargetUser == nil &&
			!invite.Revoked &&
			(invite.MaxUses == 0 || invite.Uses < invite.MaxUses) {
			resp = append(resp, invite.Code)
		}
	}

	return resp, nil
}

func (e *Entity) HandleDiscordMessage(message *discordgo.Message) (*response.HandleUserActivityResponse, error) {
	// TODO(trkhoi): temp keep hardcode of Pod Town emoji, remove when it has been added to the database
	guildConfigGm, err := e.repo.GuildConfigGmGn.GetByGuildID(message.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	isGmEmoji := strings.EqualFold(guildConfigGm.Emoji, message.Content) || strings.EqualFold("<:gm:967285238306840576>", message.Content) || strings.EqualFold("<:gm:930840080761880626>", message.Content)
	isGmSticker := false
	for _, sticker := range message.StickerItems {
		if sticker.ID == guildConfigGm.Sticker || sticker.ID == "928509218171006986" || sticker.ID == "1039136044836200549" {
			isGmSticker = true
			break
		}
	}
	var (
		discordID      = message.Author.ID
		authorAvatar   = message.Author.Avatar
		authorUsername = message.Author.Username
		guildID        = message.GuildID
		sentAt         = message.Timestamp
		channelID      = message.ChannelID
		isGmMessage    = strings.EqualFold(guildConfigGm.Msg, message.Content) || strings.EqualFold("gm", message.Content) || strings.EqualFold("gn", message.Content) || isGmEmoji || isGmSticker
	)

	switch {
	case isGmMessage:
		if guildConfigGm.ChannelID != channelID {
			// do nothing if not gm channel
			return nil, nil
		}
		return e.newUserGM(authorAvatar, authorUsername, discordID, guildID, channelID, sentAt)
	}
	return nil, nil
}

func (e *Entity) HandleMochiSalesMessage(message *request.TwitterSalesMessage) error {
	err := e.repo.MochiNFTSales.CreateOne(message)
	if err != nil {
		e.log.Errorf(err, "[e.HandleMochiSalesMessage] failed to create mochi nft sales: %s", err)
		return err
	}
	return nil
}

func (e *Entity) newUserGM(authorAvatar, authorUsername, discordID, guildID, channelID string, sentAt time.Time) (*response.HandleUserActivityResponse, error) {
	chatDate := time.Date(sentAt.Year(), sentAt.Month(), sentAt.Day(), 0, 0, 0, 0, time.UTC)
	streak, err := e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID(discordID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"discordID": discordID, "guildID": guildID}).Error(err, "[entity.newUserGM] repo.DiscordUserGMStreak.GetByDiscordIDGuildID() failed")
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		newStreak := model.DiscordUserGMStreak{
			DiscordID:      discordID,
			GuildID:        guildID,
			StreakCount:    1,
			TotalCount:     1,
			LastStreakDate: chatDate,
		}
		err = e.repo.DiscordUserGMStreak.UpsertOne(newStreak)
		if err != nil {
			e.log.Fields(logger.Fields{"newStreak": newStreak}).Error(err, "[entity.newUserGM] repo.DiscordUserGMStreak.UpsertOne() failed")
			return nil, err
		}

		// handle quest logs
		log := &model.QuestUserLog{
			GuildID: &guildID,
			UserID:  discordID,
			Action:  model.QuestAction(model.GM),
		}
		if err := e.UpdateUserQuestProgress(log); err != nil {
			e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.newUserGM] entity.UpdateUserQuestProgress() failed")
		}

		// // send data to processor to calculate user's xp
		err = e.sendGmGnMessage(discordID, guildID, &newStreak)
		if err != nil {
			e.log.Fields(logger.Fields{
				"guildID":   guildID,
				"discordID": discordID,
			}).Error(err, "[entity.newUserGM] entity.sendGmGnMessage() failed")
		}
		return nil, e.replyGmGn(&newStreak, channelID, discordID, authorAvatar, authorUsername, "", true)
	}

	nextStreakDate := streak.LastStreakDate.Add(time.Hour * 24)

	switch {
	case chatDate.Before(nextStreakDate):
		durationTilNextGoal := nextStreakDate.Sub(sentAt)
		durationString := fmt.Sprintf("%d hours and %.0f minutes",
			int(durationTilNextGoal.Hours()),
			durationTilNextGoal.Minutes()-float64(int(durationTilNextGoal.Hours()))*60)
		return nil, e.replyGmGn(streak, channelID, discordID, authorAvatar, authorUsername, durationString, false)
	case chatDate.Equal(nextStreakDate):
		streak.StreakCount++
	case chatDate.After(nextStreakDate):
		streak.StreakCount = 1
	}
	streak.LastStreakDate = chatDate
	streak.TotalCount++

	if err := e.repo.DiscordUserGMStreak.UpsertOne(*streak); err != nil {
		e.log.Fields(logger.Fields{"streak": *streak}).Error(err, "[entity.newUserGM] repo.DiscordUserGMStreak.UpsertOne() failed")
		return nil, err
	}

	// handle quest logs
	log := &model.QuestUserLog{
		GuildID: &guildID,
		UserID:  discordID,
		Action:  model.QuestAction(model.GM),
	}
	if err := e.UpdateUserQuestProgress(log); err != nil {
		e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.newUserGM] entity.UpdateUserQuestProgress() failed")
	}

	// add new feature : GmExIncrease
	///////
	if streak.StreakCount < 2 {
		// // send data to processor to calculate user's xp
		err = e.sendGmGnMessage(discordID, guildID, streak)
		if err != nil {
			e.log.Fields(logger.Fields{
				"guildID":   guildID,
				"discordID": discordID,
			}).Error(err, "[entity.newUserGM] entity.sendGmGnMessage() failed")
		}
		return nil, e.replyGmGn(streak, channelID, discordID, authorAvatar, authorUsername, "", true)
	}

	// handle activity logs
	res, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID:   guildID,
		ChannelID: channelID,
		UserID:    discordID,
		Action:    "gm_streak",
		Timestamp: sentAt,
	})
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":   guildID,
			"channelID": channelID,
			"userID":    discordID,
			"action":    "gm_streak",
		}).Error(err, "[entity.newUserGM] entity.HandleUserActivities() failed")
		return nil, err
	}

	// send data to processor to calculate user's xp
	err = e.sendGmGnMessage(discordID, guildID, streak)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":   guildID,
			"discordID": discordID,
		}).Error(err, "[entity.newUserGM] entity.sendGmGnMessage() failed")
	}

	return res, e.replyGmGn(streak, channelID, discordID, authorAvatar, authorUsername, "", true)
}

func (e *Entity) sendGmGnMessage(discordID string, guildID string, streak *model.DiscordUserGMStreak) error {
	// send data to processor to calculate user's xp
	data := model.UserTxData{
		UserDiscordId: discordID,
		Guild:         guildID,
		StreakCount:   streak.StreakCount,
		TotalCount:    streak.TotalCount,
	}

	podTownXps, err := e.svc.Processor.CreateUserTransaction(model.CreateUserTransaction{
		Dapp:   consts.NekoBot,
		Action: consts.GmStreak,
		Data:   data,
	})
	if err != nil {
		e.log.Fields(logger.Fields{
			"dapp":   consts.NekoBot,
			"action": consts.GmStreak,
			"data":   data,
		}).Error(err, "[Entity][sendGmGnMessage] failed to send data to Processor")
		return err
	}

	// send message to log channel
	guild, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[Entity][sendGmGnMessage] failed to get guild data")
		return err
	}
	err = e.svc.Discord.NotifyGmStreak(guild.LogChannel, discordID, streak.StreakCount, *podTownXps)
	if err != nil {
		e.log.Fields(logger.Fields{
			"channelID": guild.LogChannel,
			"discordID": discordID,
			"streak":    streak.StreakCount,
		}).Error(err, "[Entity][sendGmGnMessage] failed to notify gm streak log")
		return err
	}
	return nil
}

func (e *Entity) replyGmGn(streak *model.DiscordUserGMStreak, channelID, discordID, authorAvatar, authorUsername, durationTilNextGoal string, newStreakRecorded bool) error {
	if newStreakRecorded && streak.StreakCount >= 2 {
		embed := e.composeGmGnMessageEmbed(fmt.Sprintf("<@%s>, you've said gm-gn %d days in a row :fire: and %d days in total.", discordID, streak.StreakCount, streak.TotalCount), authorUsername, authorAvatar)
		_, err := e.discord.ChannelMessageSendEmbed(channelID, embed)
		return err
	}

	if !newStreakRecorded && durationTilNextGoal != "" {
		embed := e.composeGmGnMessageEmbed(fmt.Sprintf("<@%s>, you've already said gm-gn today. \nYou need to wait `%s` :alarm_clock: to reach your next streak goal :dart:", discordID, durationTilNextGoal), authorUsername, authorAvatar)
		_, err := e.discord.ChannelMessageSendEmbed(channelID, embed)
		return err
	}

	if streak.StreakCount == 1 {
		embed := e.composeGmGnMessageEmbed(fmt.Sprintf("<@%s>, you've started a gm-gn streak :sparkles: Keep it up!", discordID), authorUsername, authorAvatar)
		_, err := e.discord.ChannelMessageSendEmbed(channelID, embed)
		return err
	}

	return nil
}

func (e *Entity) composeGmGnMessageEmbed(description, authorUsername, authorAvatar string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "GM / GN",
		Description: description,
		Color:       3447003,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    authorUsername,
			IconURL: authorAvatar,
		},
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (e *Entity) ChatXPIncrease(message *discordgo.Message) (*response.HandleUserActivityResponse, error) {
	if message.Content == "" {
		return &response.HandleUserActivityResponse{
			GuildID:      message.GuildID,
			ChannelID:    message.ChannelID,
			UserID:       message.Author.ID,
			Action:       "default",
			AddedXP:      0,
			CurrentXP:    0,
			CurrentLevel: 0,
			Timestamp:    message.Timestamp,
			LevelUp:      false,
		}, nil
	}

	xpID := fmt.Sprintf(`%s_%s_chat_xp_cooldown`, message.Author.ID, message.GuildID)

	exists, err := e.cache.GetBool(xpID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat xp cooldown: %v", err.Error())
	}

	var resp *response.HandleUserActivityResponse

	if !exists {
		resp, err = e.HandleUserActivities(&request.HandleUserActivityRequest{
			GuildID:   message.GuildID,
			ChannelID: message.ChannelID,
			UserID:    message.Author.ID,
			Action:    "chat",
			Timestamp: message.Timestamp,
		})
		if err != nil {
			e.log.Fields(logger.Fields{
				"content":   message.Content,
				"guildID":   message.GuildID,
				"channelID": message.ChannelID,
				"userID":    message.Author.ID,
				"action":    "chat",
			}).Error(err, "[Entity][ChatXPIncrease] failed to handle user activity")
			return nil, fmt.Errorf("failed to handle user activity: %v", err.Error())
		}

		err = e.cache.Set(xpID, true, time.Minute)
		if err != nil {
			return nil, fmt.Errorf(`failed to set chat xp cooldown: %v`, err.Error())
		}
	}

	return resp, nil
}

func (e *Entity) BoostXPIncrease(message *discordgo.Message) (*response.HandleUserActivityResponse, error) {
	log := logger.NewLogrusLogger()
	log.Infof("New boost event start at guildID %v by user %v", message.GuildID, message.Author.ID)
	var resp *response.HandleUserActivityResponse

	resp, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID: message.GuildID,
		UserID:  message.Author.ID,
		Action:  "boost",
	})
	if err != nil {
		log.Info("Failed to handle user boost activity")
		return nil, fmt.Errorf("failed to handle user boost activity: %v", err.Error())
	}

	return resp, nil
}

func (e *Entity) WebhookUpvoteStreak(userID, source string) error {
	sentAt := time.Now().UTC()
	chatDate := time.Date(sentAt.Year(), sentAt.Month(), sentAt.Day(), sentAt.Hour(), sentAt.Minute(), 0, 0, time.UTC)
	streak, err := e.repo.DiscordUserUpvoteStreak.GetByDiscordID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.WebhookUpvoteStreak] fail to get user upvote streak")
		return fmt.Errorf("failed to get user's upvote streak: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		err = e.repo.DiscordUserUpvoteStreak.UpsertOne(model.DiscordUserUpvoteStreak{
			DiscordID:      userID,
			StreakCount:    1,
			TotalCount:     1,
			LastStreakDate: chatDate,
		})
		if err != nil {
			e.log.Errorf(err, "[e.WebhookUpvoteStreak] fail to create new streak")
			return fmt.Errorf("failed to create new user upvote streak: %v", err)
		}
		// handle quest logs
		log := &model.QuestUserLog{
			UserID: userID,
			Action: model.QuestAction(model.VOTE),
		}
		if err := e.UpdateUserQuestProgress(log); err != nil {
			e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.WebhookUpvoteStreak] entity.UpdateUserQuestProgress() failed")
		}
		return nil
	}

	nextStreakDate := streak.LastStreakDate.Add(time.Hour * 24)

	switch {
	case chatDate.Before(nextStreakDate):
		streak.StreakCount++
	case chatDate.Equal(nextStreakDate):
		streak.StreakCount++
	case chatDate.After(nextStreakDate):
		streak.StreakCount = 1
	}
	streak.LastStreakDate = chatDate
	streak.TotalCount++

	if err := e.repo.DiscordUserUpvoteStreak.UpsertOne(*streak); err != nil {
		e.log.Errorf(err, "[e.WebhookUpvoteStreak] fail to upsert upvote streak")
		return fmt.Errorf("failed to update user upvote streak: %v", err)
	}

	// handle quest logs
	log := &model.QuestUserLog{
		UserID: userID,
		Action: model.QuestAction(model.VOTE),
	}
	if err := e.UpdateUserQuestProgress(log); err != nil {
		e.log.Fields(logger.Fields{"log": log}).Error(err, "[entity.WebhookUpvoteStreak] entity.UpdateUserQuestProgress() failed")
	}

	e.handleUpvoteXPBonus(streak)
	e.logUserUpvote(userID, source, chatDate)

	_, err = e.repo.Users.GetOne(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.WebhookUpvoteStreak] failed to get user")
		return err
	}
	isStranger := err != nil
	e.svc.Discord.SendUpvoteMessage(userID, source, isStranger)

	cache, err := e.GetUpvoteMessageCache(userID)
	if err != nil {
		e.log.Errorf(err, "[e.WebhookUpvoteStreak] failed to get cache")
		return err
	}
	if cache == nil {
		return nil
	}
	e.svc.Discord.ReplyUpvoteMessage(cache, source)
	e.RemoveUpvoteMessageCache(userID)
	return nil
}

func (e *Entity) handleUpvoteXPBonus(streak *model.DiscordUserUpvoteStreak) error {
	tier, err := e.repo.UpvoteStreakTier.GetByUpvoteCount(streak.StreakCount)
	if err != nil {
		e.log.Errorf(err, "[e.handleUpvoteXPBonus] failed to get upvote tier")
		return err
	}

	earnedXP := 0

	// increase x for every x votes
	if streak.TotalCount%tier.VoteInterval == 0 {
		earnedXP = tier.XPPerInterval
		// send data to processor to calculate user's xp
		data := model.UserTxData{
			UserDiscordId: streak.DiscordID,
		}
		_, err = e.svc.Processor.CreateUserTransaction(model.CreateUserTransaction{
			Dapp:   consts.NekoBot,
			Action: consts.Vote,
			Data:   data,
		})
		if err != nil {
			e.log.Fields(logger.Fields{
				"dapp":   consts.NekoBot,
				"action": consts.Vote,
				"data":   data,
			}).Error(err, "[Entity][handleUpvoteXPBonus] failed to send data to Processor")
		}
	}
	if err := e.repo.GuildUserActivityLog.CreateOneNoGuild(model.GuildUserActivityLog{
		UserID:       streak.DiscordID,
		ActivityName: "upvote",
		EarnedXP:     earnedXP,
		CreatedAt:    time.Now(),
	}); err != nil {
		e.log.
			Fields(logger.Fields{"userID": streak.DiscordID}).
			Error(err, "[Entity][handleUpvoteXPBonus] failed to create guild_user_activity_logs")
		return err
	}
	return nil
}

func (e *Entity) logUserUpvote(userID, source string, chatTime time.Time) error {
	err := e.repo.DiscordUserUpvoteLog.UpsertOne(model.DiscordUserUpvoteLog{
		DiscordID:        userID,
		Source:           source,
		LatestUpvoteTime: chatTime,
	})
	if err != nil {
		e.log.Errorf(err, "[e.logUpvoteActivity] failed to log upvote")
		return err
	}
	return nil
}

func (e Entity) RemoveAllMessageReactions(message *discordgo.Message) error {
	cfg, err := e.repo.GuildConfigReactionRole.GetByMessageID(message.GuildID, message.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": message.GuildID, "messageID": message.ID}).
			Info("[e.RemoveAllMessageReactions] this message is not reaction role config for guild")
		return nil
	}

	roles := []response.Role{}
	if err := json.Unmarshal([]byte(cfg.ReactionRoles), &roles); err != nil {
		e.log.Fields(logger.Fields{"ReactionRoles": cfg.ReactionRoles}).
			Error(err, "[e.RemoveAllMessageReactions] failed to unmarshal reaction roles")
		return err
	}
	rolesMap := map[string]string{}

	for _, role := range roles {
		rolesMap[role.Reaction] = role.ID
	}

	msgReactions, err := e.repo.MessageReaction.GetByMessageID(message.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"ReactionRoles": cfg.ReactionRoles}).
			Error(err, "[e.RemoveAllMessageReactions] failed to get message reactions")
		return err
	}

	for _, msgReact := range msgReactions {
		if _, ok := rolesMap[msgReact.Reaction]; ok {
			err := util.RetryRequest(func() error {
				return e.RemoveGuildMemberRole(msgReact.GuildID, msgReact.UserID, rolesMap[msgReact.Reaction])
			}, 10, time.Second)

			if err != nil {
				e.log.Fields(logger.Fields{
					"guildID": msgReact.GuildID,
					"userID":  msgReact.UserID,
					"roleID":  rolesMap[msgReact.Reaction],
				}).Infof("[e.RemoveAllMessageReactions] failed to get message reactions %v", err)
			}
		}
	}

	if err := e.repo.MessageReaction.DeleteByMessageID(message.ID); err != nil {
		e.log.Fields(logger.Fields{"messageID": message.ID}).Error(err, "[e.RemoveAllMessageReactions] failed to delete message reactions")
		return err
	}

	if err := e.repo.GuildConfigReactionRole.ClearMessageConfig(message.GuildID, message.ID); err != nil {
		e.log.Fields(logger.Fields{"messageID": message.ID, "guildID": message.GuildID}).
			Error(err, "[e.RemoveAllMessageReactions] failed to clear message config")
		return err
	}

	return nil
}

func (e *Entity) HandleMemberRemove(req *request.MemberRemoveWebhookRequest) error {
	jlChannel, err := e.GetJoinLeaveChannelConfig(req.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[e.GetJoinLeaveChannelConfig] failed to get channel config")
		return err
	}

	//joinleave channel not set
	if jlChannel == nil {
		return nil
	}
	err = e.svc.Discord.NotifyMemberLeave(req, jlChannel.ChannelID)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[discord.NotifyMemberLeave] failed to send notification")
		return err
	}
	return nil
}

func (e *Entity) HandleMemberAdd(member *discordgo.Member) error {
	jlChannel, err := e.GetJoinLeaveChannelConfig(member.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guildID": member.GuildID}).Error(err, "[e.GetJoinLeaveChannelConfig] failed to get channel config")
		return err
	}
	//joinleave channel not set
	if jlChannel == nil {
		return nil
	}

	guild, err := e.discord.GuildWithCounts(member.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": member.GuildID}).Error(err, "[e.discord.GuildWithCounts] failed to get guild")
		return err
	}

	err = e.svc.Discord.NotifyMemberJoin(member.User.ID, member.AvatarURL("80x80"), jlChannel.ChannelID, int64(guild.ApproximateMemberCount)) //guild.MemberCount only exists in guildCreate event
	if err != nil {
		e.log.Fields(logger.Fields{"member": member}).Error(err, "[discord.NotifyMemberJoin] failed to send notification")
		return err
	}
	return nil
}
