package entities

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"

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
	var (
		discordID      = message.Author.ID
		authorAvatar   = message.Author.Avatar
		authorUsername = message.Author.Username
		guildID        = message.GuildID
		sentAt         = message.Timestamp
		channelID      = message.ChannelID
		isGmMessage    = strings.EqualFold("gm", message.Content) || strings.EqualFold("gn", message.Content)
	)

	switch {
	case isGmMessage:
		guildConfigGm, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
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
		return nil, fmt.Errorf("failed to get user's gm streak: %v", err)
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
			return nil, fmt.Errorf("failed to create new user gm streak: %v", err)
		}

		// // send data to processor to calculate user's xp
		err = e.sendGmGnMessage(discordID, guildID, &newStreak)
		if err != nil {
			e.log.Fields(logger.Fields{
				"guildID":   guildID,
				"discordID": discordID,
			}).Error(err, "[Entity][newUserGM] failed to send GmGn message")
		}
		return nil, nil
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
		return nil, fmt.Errorf("failed to update user gm streak: %v", err)
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
			}).Error(err, "[Entity][newUserGM] failed to send GmGn message")
		}

		return nil, e.replyGmGn(streak, channelID, discordID, authorAvatar, authorUsername, "", true)
	}

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
		}).Error(err, "[Entity][newUserGM] failed to handle user activity")
		return nil, fmt.Errorf("failed to handle user activity: %v", err.Error())
	}

	// send data to processor to calculate user's xp
	err = e.sendGmGnMessage(discordID, guildID, streak)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":   guildID,
			"discordID": discordID,
		}).Error(err, "[Entity][newUserGM] failed to send GmGn message")
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
	if newStreakRecorded && streak.StreakCount >= 3 {
		_, err := e.discord.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title:       "GM / GN",
			Description: fmt.Sprintf("<@%s>, you've said gm-gn %d days in a row :fire: and %d days in total.", discordID, streak.StreakCount, streak.TotalCount),
			Color:       3447003,
			Footer: &discordgo.MessageEmbedFooter{
				Text:    authorUsername,
				IconURL: authorAvatar,
			},
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		})
		return err
	}

	if !newStreakRecorded && durationTilNextGoal != "" {
		_, err := e.discord.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title:       "GM / GN",
			Description: fmt.Sprintf("<@%s>, you've already said gm-gn today. \nYou need to wait `%s` :alarm_clock: to reach your next streak goal :dart:", discordID, durationTilNextGoal),
			Color:       3447003,
			Footer: &discordgo.MessageEmbedFooter{
				Text:    authorUsername,
				IconURL: authorAvatar,
			},
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		})
		return err
	}

	return nil
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

func (e *Entity) WebhookUpvoteStreak(userID string) error {
	sentAt := time.Now()
	chatDate := time.Date(sentAt.Year(), sentAt.Month(), sentAt.Day(), sentAt.Hour(), 0, 0, 0, time.UTC)
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
		return nil
	}

	nextStreakDate := streak.LastStreakDate.Add(time.Hour * 12)

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
	return nil
}
