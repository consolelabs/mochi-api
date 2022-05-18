package entities

import (
	"fmt"
	"strconv"
	"time"

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

func (e *Entity) HandleDiscordMessage(message *discordgo.Message) error {
	var (
		discordID = message.Author.ID
		guildID   = message.GuildID
		sentAt    = message.Timestamp
		channelID = message.ChannelID
	)

	isGmMessage := message.Content == "gm" || message.Content == "gn"

	switch {
	case isGmMessage:
		guildConfigGm, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
		if err != nil {
			return err
		}
		if guildConfigGm.ChannelID != channelID {
			// do nothing if not gm channel
			return nil
		}
		return e.newUserGM(discordID, guildID, channelID, sentAt)
	}
	return nil
}

func (e *Entity) newUserGM(discordID, guildID, channelID string, sentAt time.Time) error {
	chatDate := time.Date(sentAt.Year(), sentAt.Month(), sentAt.Day(), 0, 0, 0, 0, time.UTC)
	streak, err := e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID(discordID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to get user's gm streak: %v", err)
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
			return fmt.Errorf("failed to create new user gm streak: %v", err)
		}
		return nil
	}

	nextStreakDate := streak.LastStreakDate.Add(time.Hour * 24)

	switch {
	case chatDate.Before(nextStreakDate):
		durationTilNextGoal := nextStreakDate.Sub(sentAt).String()
		return e.replyGmGn(streak, channelID, discordID, durationTilNextGoal, false)
	case chatDate.Equal(nextStreakDate):
		streak.StreakCount++
	case chatDate.After(nextStreakDate):
		streak.StreakCount = 1
	}
	streak.LastStreakDate = chatDate
	streak.TotalCount++

	if err := e.repo.DiscordUserGMStreak.UpsertOne(*streak); err != nil {
		return fmt.Errorf("failed to update user gm streak: %v", err)
	}
	return e.replyGmGn(streak, channelID, discordID, "", true)
}

func (e *Entity) replyGmGn(streak *model.DiscordUserGMStreak, channelID, discordID, durationTilNextGoal string, newStreakRecorded bool) error {
	if newStreakRecorded && streak.StreakCount >= 3 {
		_, err := e.discord.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title:       "GM / GN",
			Description: fmt.Sprintf("<@%s>, you've said gm-gn %d days in a row :fire: and %d days in total.", discordID, streak.StreakCount, streak.TotalCount),
		})
		return err
	}

	if !newStreakRecorded && durationTilNextGoal != "" {
		_, err := e.discord.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title:       "GM / GN",
			Description: fmt.Sprintf("<@%s>, you've already said gm-gn today. You need to wait `%s` :alarm_clock: to reach your next streak goal :dart:.", discordID, durationTilNextGoal),
		})
		return err
	}

	return nil
}

func (e *Entity) ChatXPIncrease(message *discordgo.Message) (*response.HandleUserActivityResponse, error) {
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
			return nil, fmt.Errorf("failed to handle user activity: %v", err.Error())
		}

		err = e.cache.Set(xpID, true, time.Minute)
		if err != nil {
			return nil, fmt.Errorf(`failed to set chat xp cooldown: %v`, err.Error())
		}
	}

	return resp, nil
}
