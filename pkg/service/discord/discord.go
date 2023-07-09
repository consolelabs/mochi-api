package discord

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Discord struct {
	session                *discordgo.Session
	log                    logger.Logger
	repo                   *repo.Repo
	mochiGuildID           string
	mochiLogChannelID      string
	mochiSaleChannelID     string
	mochiActivityChannelID string
	mochiFeedbackChannelID string
}

const (
	mochiLogColor     = 0x62A1FE
	mochiTipLogColor  = 0xFFDC50
	mochiErrorColor   = 0xD94F50
	mochiSuccessColor = 0x5cd97d
)

func NewService(
	cfg config.Config,
	log logger.Logger,
	repo *repo.Repo,
) (Service, error) {
	// *** discord ***
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}
	return &Discord{
		session:                discord,
		log:                    log,
		mochiGuildID:           cfg.MochiGuildID,
		mochiLogChannelID:      cfg.MochiLogChannelID,
		mochiSaleChannelID:     cfg.MochiSaleChannelID,
		mochiActivityChannelID: cfg.MochiActivityChannelID,
		mochiFeedbackChannelID: cfg.MochiFeedbackChannelID,
		repo:                   repo,
	}, nil
}

func (d *Discord) NotifyNewGuild(guildID string, count int) error {
	// get new guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		d.log.Errorf(err, "[discord.NotifyNewGuild] - failed to get guild: %s", guildID)
		return fmt.Errorf("failed to get guild info: %w", err)
	}
	inviteUrl := d.generateGuildInviteLink(guild)
	res, err := d.session.GuildWithCounts(guildID)
	if err != nil {
		d.log.Errorf(err, "[discord.NotifyNewGuild] - cannot get total members of guild: %s", guildID)
		return err
	}

	postfix := util.NumberPostfix(count)

	newGuildMsg := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label: "Jump to server",
						Style: discordgo.LinkButton,
						URL:   inviteUrl,
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       "New Guild!",
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: guild.IconURL()},
			Description: fmt.Sprintf("**Name**: `%s`\n**Members**: `%v`", guild.Name, res.ApproximateMemberCount),
			Color:       mochiLogColor,
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("The %v%s guild", count, postfix)},
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		},
	}

	_, err = d.session.ChannelMessageSendComplex(d.mochiLogChannelID, newGuildMsg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) NotifyAddNewCollection(guildID string, collectionName string, symbol string, chain string, image string) error {
	// get guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild info: %w", err)
	}

	msgEmbed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s just added a new collection", guild.Name),
		Description: fmt.Sprintf("%s (%s) on chain %s", collectionName, symbol, strings.ToUpper(chain)),
		Color:       mochiLogColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err = d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) NotifyGmStreak(channelID string, userDiscordID string, streakCount int, podTownXps model.CreateUserTxResponse) error {
	color, _ := strconv.ParseInt("6FC1D1", 16, 64)
	approveIcon := ""
	if streakCount <= 100 {
		for i := 0; i < streakCount; i++ {
			approveIcon += "<:approve:1077631110047080478>"
		}
	} else {
		for i := 0; i < 100; i++ {
			approveIcon += "<:approve:1077631110047080478>"
		}
		approveIcon += "(+" + strconv.Itoa(int(streakCount-100)) + "<:approve:1077631110047080478>)"
	}

	msgEmbed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Good Morning!",
			IconURL: "https://cdn.discordapp.com/attachments/701029345795375114/1013773058068201482/mochi.jpeg",
		},
		Description: "<@" + userDiscordID + "> just said hi to everyone.\nGM streak: **" +
			strconv.Itoa(streakCount) + "**\n" + approveIcon + "\n\n" +
			"**Faction XP Update**\n<:rebelio:932605621914701875> Rebellio EXP: **" +
			strconv.Itoa(int(podTownXps.Data.TotalFameXps)) + "/" + strconv.Itoa(int(podTownXps.Data.NextFameXps)) +
			"`(+" + strconv.Itoa(int(podTownXps.Data.FameXp)) + ")`" +
			"**\n<:academia:932605621730160680> Academy EXP: **" +
			strconv.Itoa(int(podTownXps.Data.TotalReputationXps)) + "/" + strconv.Itoa(int(podTownXps.Data.NextReputationXps)) +
			"**`(+" + strconv.Itoa(int(podTownXps.Data.ReputationXp)) + ")`",
		Color:     int(color),
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err := d.session.ChannelMessageSendEmbed(channelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) SendGuildActivityLogs(channelID, userID, title, description string) error {
	dcUser, err := d.session.User(userID)
	if err != nil {
		d.log.Errorf(err, "[SendGuildActivityLogs] - get discord user failed %s", userID)
		return err
	}
	msgEmbed := discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: dcUser.AvatarURL(""),
		},
	}

	_, err = d.session.ChannelMessageSendEmbed(channelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("[SendGuildActivityLogs] - ChannelMessageSendEmbed failed - channel %s: %s", channelID, err.Error())
	}

	return nil
}

func (d *Discord) SendLevelUpMessage(levelUpConfig *model.GuildConfigLevelupMessage, role string, levelRoleLevel int, randomTip string, uActivity *response.HandleUserActivityResponse) {
	// priority: config channel -> chat channel
	channelID := levelUpConfig.ChannelID
	if levelUpConfig.ChannelID == "" {
		channelID = uActivity.ChannelID
	}
	if channelID == "" {
		d.log.Info("Action was not performed at any channel")
		return
	}
	if role == "" {
		role = "N/A"
	}

	// TODO: get emoji from backend instead of hardcode

	dcUser, err := d.session.User(uActivity.UserID)
	if err != nil {
		d.log.Errorf(err, "SendLevelUpMessage - failed to get discord user %s", uActivity.UserID)
		return
	}

	msgEmbed := d.formatLevelUpMessage(uActivity, dcUser, role, randomTip, levelRoleLevel)
	d.log.Fields(logger.Fields{"channelId": channelID, "userID": uActivity.UserID}).Info("Sending level up message")
	_, err = d.session.ChannelMessageSendEmbed(channelID, msgEmbed)
	if err != nil {
		d.log.Errorf(err, "SendLevelUpMessage - failed to send level up msg")
		return
	}
}

func (d *Discord) formatLevelUpMessage(uActivity *response.HandleUserActivityResponse, dcUser *discordgo.User, role, randomTip string, levelRoleLevel int) *discordgo.MessageEmbed {
	starEmoji, err := d.repo.Emojis.ListEmojis([]string{"STAR"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get star emoji")
		return nil
	}

	xpEmoji, err := d.repo.Emojis.ListEmojis([]string{"XP"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get xp emoji")
		return nil
	}

	gemEmoji, err := d.repo.Emojis.ListEmojis([]string{"GEM"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get gem emoji")
		return nil
	}

	badgeEmoji, err := d.repo.Emojis.ListEmojis([]string{"BADGE"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get badge emoji")
		return nil
	}

	pointRightEmoji, err := d.repo.Emojis.ListEmojis([]string{"POINT_RIGHT"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get point right emoji")
		return nil
	}

	trophyEmoji, err := d.repo.Emojis.ListEmojis([]string{"TROPHY"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get trophy emoji")
		return nil
	}

	lineEmoji, err := d.repo.Emojis.ListEmojis([]string{"LINE"})
	if err != nil {
		d.log.Errorf(err, "formatLevelUpMessage - failed to get line emoji")
		return nil
	}

	description := []string{fmt.Sprintf("%s Congrats <@%s> on leveling up.\n", *starEmoji[0].DiscordId, uActivity.UserID)}
	description = append(description, fmt.Sprintf("%s Your current level is %d.\n", *xpEmoji[0].DiscordId, uActivity.CurrentLevel))
	description = append(description, fmt.Sprintf("%s To reach level %d, you would now need to have %d xp.\n", *gemEmoji[0].DiscordId, uActivity.CurrentLevel+1, uActivity.NextLevel.MinXP))
	description = append(description, fmt.Sprintf("%s The next level role is %s, which is at level %d.\n", *badgeEmoji[0].DiscordId, role, levelRoleLevel))
	description = append(description, strings.Repeat(*lineEmoji[0].DiscordId, 10)+"\n")
	description = append(description, fmt.Sprintf("%s Here are some things you can do to accrue xp:\n", *pointRightEmoji[0].DiscordId))
	description = append(description, "+ Chatting\n")
	description = append(description, "+ Invite your frens\n")
	description = append(description, "+ Doing quests\n")
	description = append(description, "+ Swap/tip/deposit/withdraw\n")

	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s Level up!", *trophyEmoji[0].DiscordId),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: dcUser.AvatarURL(""),
		},
		Description: strings.Join(description, ""),
		Color:       mochiLogColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    randomTip,
			IconURL: "",
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (d *Discord) NotifyStealFloorPrice(price float64, floor float64, url string, name string, image string) error {
	msgEmbed := discordgo.MessageEmbed{
		Title:       "NFT Steal Alert!",
		Description: fmt.Sprintf("%s was listed at %v, under the floor price of %v\nClick here to buy now [Market](%s)", name, price, floor, url),
		Color:       mochiLogColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: util.StandardizeUri(image),
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendEmbed(d.mochiSaleChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) NotifyStealAveragePrice(price float64, avg float64, url string, name string, image string) error {
	msgEmbed := discordgo.MessageEmbed{
		Title:       "NFT Steal Alert!",
		Description: fmt.Sprintf("%s was listed at %v, under the average price of %v\nClick here to buy now [Market](%s)", name, price, avg, url),
		Color:       mochiLogColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: util.StandardizeUri(image),
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendEmbed(d.mochiSaleChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type string) error {
	if guildID == "" || logChannelID == "" || userID == "" || roleID == "" {
		return nil
	}

	member, err := d.session.GuildMember(guildID, userID)
	if err != nil {
		d.log.Errorf(err, "[svc.Discord.SendUpdateRolesLog] - session.GuildMember failed %s", userID)
		return err
	}

	description := fmt.Sprintf("<@%s> has been assigned a new role\n<:_:1090477901725577287>**Type**: %s\n<:_:1098461538609799278>**Role**: <@&%s>", userID, _type, roleID)
	msgEmbed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/1090477916132999270.png?size=240&quality=lossless",
			Name:    "New role granted",
		},
		Description: description,
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Try /tip or /airdrop to send out money",
		},
	}

	_, err = d.session.ChannelMessageSendEmbed(logChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("[svc.Discord.SendUpdateRolesLog] - ChannelMessageSendEmbed failed to channel %s: %s", logChannelID, err.Error())
	}
	return nil
}

func (d *Discord) generateGuildInviteLink(guild *discordgo.Guild) string {
	inviteUrl := ""
	channels, err := d.session.GuildChannels(guild.ID)
	if err != nil {
		d.log.Infof("[discord.NotifyNewGuild] guild %s has no channels", guild.ID)
		return ""
	}
	for _, ch := range channels {
		if ch.Type == discordgo.ChannelTypeGuildText {
			invite, err := d.session.ChannelInviteCreate(ch.ID, discordgo.Invite{Guild: guild, Channel: ch})
			if err == nil {
				inviteUrl = fmt.Sprintf("https://discord.gg/%s", invite.Code)
				break
			}
		}
	}
	return inviteUrl
}

func (d *Discord) NotifyGuildDelete(guildID, guildName, iconURL string, guildsLeft int) error {
	msg := &discordgo.MessageEmbed{
		Title:       "Time to say goodbye!",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: iconURL},
		Description: fmt.Sprintf("Mochi just left guild `%s`\nGuilds left: `%v`", guildName, guildsLeft),
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, msg)
	if err != nil {
		d.log.Fields(logger.Fields{"msg": msg}).Error(err, "session.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (d *Discord) SendFeedback(req *request.UserFeedbackRequest, feedbackID string) error {
	title := "Feedback received."
	if req.Command != "" {
		title = fmt.Sprintf("Feedback received for command: %s", req.Command)
	}
	msg := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "In progress",
						Style:    2,
						CustomID: fmt.Sprintf("feedback_handle-set-in-progress_%s", feedbackID),
					},
				},
			},
		},
		Embed: &discordgo.MessageEmbed{
			Title:       title,
			Footer:      &discordgo.MessageEmbedFooter{Text: req.Username, IconURL: req.Avatar},
			Description: fmt.Sprintf("<@%s> has something to say\n`%s`", req.DiscordID, req.Feedback),
			Color:       mochiLogColor,
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		},
	}
	_, err := d.session.ChannelMessageSendComplex(d.mochiFeedbackChannelID, msg)
	if err != nil {
		d.log.Fields(logger.Fields{"body": req}).Error(err, "session.ChannelMessageSendEmbed() failed")
		return err
	}
	return nil
}

func (d *Discord) SendTipActivityLogs(channelID, userID string, author *discordgo.MessageEmbedAuthor, description, image string) error {
	if channelID == "" {
		return nil
	}
	dcUser, err := d.session.User(userID)
	if err != nil {
		d.log.Errorf(err, "[SendTipActivityLogs] - get discord user failed %s", userID)
		return err
	}
	msgEmbed := discordgo.MessageEmbed{
		Author:      author,
		Description: description,
		Color:       mochiTipLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: dcUser.AvatarURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Try /tip or /airdrop to send out money",
		},
	}
	if image != "" {
		msgEmbed.Image = &discordgo.MessageEmbedImage{URL: image}
	}
	_, err = d.session.ChannelMessageSendEmbed(channelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("[SendTipActivityLogs] - ChannelMessageSendEmbed failed - channel %s: %s", channelID, err.Error())
	}

	return nil
}

func (d *Discord) NotifyCompleteCollectionIntegration(guildID string, collectionName string, symbol string, chain string, image string) error {
	// get guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		d.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[discord.NotifyCompleteCollectionIntegration] d.session.Guild() failed")
		return err
	}

	msgEmbed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Collection %s integrated", collectionName),
		Description: fmt.Sprintf("**Guild: ** `%s`\n**Symbol: ** `%s`\n**Chain: ** `%s`", guild.Name, symbol, chain),
		Color:       mochiLogColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err = d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, &msgEmbed)
	if err != nil {
		d.log.Error(err, "[discord.NotifyCompleteCollectionIntegration] d.session.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (d *Discord) NotifyCompleteCollectionSync(guildID string, collectionName string, symbol string, chain string, image string) error {
	// get guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		d.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[discord.NotifyCompleteCollectionSync] d.session.Guild() failed")
		return err
	}

	msgEmbed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Collection %s synced", collectionName),
		Description: fmt.Sprintf("**Guild: ** `%s`\n**Symbol: ** `%s`\n**Chain: ** `%s`", guild.Name, symbol, chain),
		Color:       mochiLogColor,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err = d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, &msgEmbed)
	if err != nil {
		d.log.Error(err, "[discord.NotifyCompleteCollectionSync] d.session.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (d *Discord) Channel(channelID string) (*discordgo.Channel, error) {
	if channelID == "" {
		return nil, errors.ErrInvalidDiscordChannelID
	}
	channel, err := d.session.Channel(channelID)
	if err != nil {
		d.log.Error(err, "[discord.Channel] d.session.Channel() failed")
		return nil, err

	}
	return channel, nil
}

func (d *Discord) CreateChannel(guildID string, createData discordgo.GuildChannelCreateData) (*discordgo.Channel, error) {
	if guildID == "" {
		return nil, errors.ErrInvalidDiscordGuildID
	}
	channel, err := d.session.GuildChannelCreateComplex(guildID, createData)
	if err != nil {
		d.log.Fields(logger.Fields{
			"guildID":    guildID,
			"createData": createData,
		}).Error(err, "[discord.CreateChannel] d.session.GuildChannelCreateComplex() failed")
		return nil, err

	}
	return channel, nil
}

func (d *Discord) DeleteChannel(channelId string) error {
	_, err := d.session.ChannelDelete(channelId)
	if err != nil {
		d.log.Error(err, "[discord.DeleteChannel] d.session.ChannelDelete() failed")
		return err
	}
	return nil
}

func (d *Discord) SendMessage(channelID string, payload discordgo.MessageSend) error {
	if channelID == "" {
		return errors.ErrInvalidDiscordChannelID
	}
	if _, err := d.session.ChannelMessageSendComplex(channelID, &payload); err != nil {
		d.log.Fields(logger.Fields{
			"channelID": channelID,
			"payload":   payload,
		}).Error(err, "[discord.SendMessage] d.session.ChannelMessageSendComplex() failed")
		return err
	}
	return nil
}

func (d *Discord) SendDMUserPriceAlert(userID, symbol string, alertType model.AlertType, price float64) error {
	var description string
	switch alertType {
	case model.PriceReaches:
		description = fmt.Sprintf("%v reaches %v", symbol, price)
	case model.PriceDropsTo:
		description = fmt.Sprintf("<:_:1058304303888093194> %v is under %v", symbol, price)
	case model.PriceRisesAbove:
		description = fmt.Sprintf("<:_:1058304334779125780> %v rises above %v", symbol, price)
	case model.ChangeIsOver:
		description = fmt.Sprintf("<:_:1058304334779125780> %v is up by %v%%", symbol, price)
	case model.ChangeIsUnder:
		description = fmt.Sprintf("<:_:1058304303888093194> %v is down by %v%%", symbol, price)
	}
	privChan, err := d.session.UserChannelCreate(userID)
	if err != nil {
		d.log.Error(err, "[discord.SendDMUserPriceAlert] d.session.UserChannelCreate() failed")
		return err
	}
	msg := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    fmt.Sprintf("%s price has changed", symbol),
				IconURL: "https://cdn.discordapp.com/emojis/1095990150342918205.gif?size=240&quality=lossless",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://cdn.discordapp.com/attachments/984660970624409630/1098472181631045742/Mochi_Pose_14.png",
			},
			Description: description,
			Color:       mochiTipLogColor,
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Try /alert to setup new alert",
			},
		},
	}
	d.session.ChannelMessageSendComplex(privChan.ID, msg)
	return nil
}

func (d *Discord) SendDM(userID string, payload discordgo.MessageSend) error {
	privChan, err := d.session.UserChannelCreate(userID)
	if err != nil {
		d.log.Error(err, "[discord.SendDM] d.session.UserChannelCreate() failed")
		return err
	}
	if _, err := d.session.ChannelMessageSendComplex(privChan.ID, &payload); err != nil {
		d.log.Fields(logger.Fields{
			"userID":  userID,
			"payload": payload,
		}).Error(err, "[discord.SendDM] d.session.ChannelMessageSendComplex() failed")
		return err
	}
	return nil
}

func (d *Discord) GetGuildMembers(guildID string) ([]*discordgo.Member, error) {
	result := []*discordgo.Member{}

	next := true

	lastID := ""

	for next {
		members, err := d.session.GuildMembers(guildID, lastID, 1000)
		if err != nil {
			d.log.Error(err, "[discord.GetGuildMembers] d.session.GuildMembers() failed")
			break
		}

		result = append(result, members...)

		if len(members) < 1000 {
			next = false
		}

		lastID = members[len(members)-1].User.ID
	}

	return result, nil
}

func (d *Discord) GetUser(userID string) (*discordgo.User, error) {
	user, err := d.session.User(userID)
	if err != nil {
		d.log.Error(err, "[discord.GetGuild] d.session.Guild() failed")
		return nil, err
	}

	return user, nil
}

func (d *Discord) GetGuild(guildID string) (*discordgo.Guild, error) {
	guild, err := d.session.Guild(guildID)
	if err != nil {
		d.log.Error(err, "[discord.GetGuild] d.session.Guild() failed")
		return nil, err
	}

	// build guild icon url
	guild.Icon = discordgo.EndpointGuildIcon(guildID, guild.Icon)

	return guild, nil
}

func (d *Discord) GetGuildRoles(guildID string) ([]*model.DiscordGuildRole, error) {
	resp, err := d.session.Request("GET", discordgo.EndpointGuildRoles(guildID), nil)
	if err != nil {
		d.log.Error(err, "[discord.GetGuildRoles] d.session.Request() failed")
		return nil, err
	}

	var roles []*model.DiscordGuildRole
	if err := json.Unmarshal(resp, &roles); err != nil {
		d.log.Error(err, "[discord.GetGuildRoles] json.Unmarshal() failed")
		return nil, err
	}

	return roles, nil
}
