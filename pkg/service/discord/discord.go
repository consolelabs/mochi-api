package discord

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Discord struct {
	session                *discordgo.Session
	log                    logger.Logger
	mochiGuildID           string
	mochiLogChannelID      string
	mochiSaleChannelID     string
	mochiActivityChannelID string
	mochiFeedbackChannelID string
}

const (
	mochiLogColor       = 0xFCD3C1
	mochiUpvoteMsgColor = 0x47ffc2
	mochiErrorColor     = 0xD94F50
	mochiSuccessColor   = 0x5cd97d
)

func NewService(
	cfg config.Config,
	log logger.Logger,
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
			approveIcon += "<:approve:933341948402618378>"
		}
	} else {
		for i := 0; i < 100; i++ {
			approveIcon += "<:approve:933341948402618378>"
		}
		approveIcon += "(+" + strconv.Itoa(int(streakCount-100)) + "<:approve:933341948402618378>)"
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

func (d *Discord) SendLevelUpMessage(levelUpConfig *model.GuildConfigLevelupMessage, role string, uActivity *response.HandleUserActivityResponse) {
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

	description := fmt.Sprintf("<:pumpeet:930840081554624632> **You are leveled up to level %d, <@%s>**", uActivity.CurrentLevel, uActivity.UserID)
	if levelUpConfig.Message != "" {
		description = description + "\n\n" + strings.Replace(levelUpConfig.Message, "$name", fmt.Sprintf("<@%s>", uActivity.UserID), -1)
		description = strings.Replace(description, `\n`, "\n", -1)
	} else {
		description = description + "\n\nThe results of hard work and dedication always look like luck to some. But you know you've earned every ounce of your success. <:mooning:930840083278487562> "
	}

	dcUser, err := d.session.User(uActivity.UserID)
	if err != nil {
		d.log.Errorf(err, "SendLevelUpMessage - failed to get discord user %s", uActivity.UserID)
		return
	}
	image := ""
	if levelUpConfig.ImageURL != "" {
		image = levelUpConfig.ImageURL
	}

	msgEmbed := discordgo.MessageEmbed{
		Title: "<:trophy:1060414870895464478> Level up!",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: dcUser.AvatarURL(""),
		},
		Image: &discordgo.MessageEmbedImage{
			URL: image,
		},
		Description: description,
		Color:       mochiLogColor,
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("XP achieved: %d", uActivity.CurrentXP),
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	d.log.Fields(logger.Fields{"channelId": channelID, "userID": uActivity.UserID}).Info("Sending level up message")
	_, err = d.session.ChannelMessageSendEmbed(channelID, &msgEmbed)
	if err != nil {
		d.log.Errorf(err, "SendLevelUpMessage - failed to send level up msg")
		return
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

	description := fmt.Sprintf("<@%s> has been assigned a new role\n**Type**: %s\n**Role**: <@&%s>", userID, _type, roleID)
	msgEmbed := discordgo.MessageEmbed{
		Title:       "Role updated",
		Description: description,
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: member.AvatarURL(""),
		},
	}

	_, err = d.session.ChannelMessageSendEmbed(logChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("[svc.Discord.SendUpdateRolesLog] - ChannelMessageSendEmbed failed to channel %s: %s", logChannelID, err.Error())
	}
	return nil
}

func (d *Discord) SendUpvoteMessage(discordID, source string, isStranger bool) error {
	if discordID == "" || source == "" {
		return nil
	}

	voteRemindStr := "\n\nCheck your progress and vote for Mochi with `$vote`"
	msgEmbed := discordgo.MessageEmbed{}
	if isStranger {
		// user can upvote without being in a guild
		sourceName, sourceUrl := util.UpvoteSourceNameAndUrl(source)
		msgEmbed = discordgo.MessageEmbed{
			Title:       "Thank you stranger.",
			Description: fmt.Sprintf("A mysterious person just upvoted Mochi on [%s](%s). Thank you, whoever you are", sourceName, sourceUrl) + voteRemindStr,
			Color:       mochiUpvoteMsgColor,
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Image: &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/986854719999864863/1019481825804029972/unknown.png",
			},
		}
	} else {
		embed := util.GenerateUpvoteMessage(discordID, source)
		msgEmbed = discordgo.MessageEmbed{
			Title:       embed.Title,
			Description: embed.Description + voteRemindStr,
			Color:       mochiUpvoteMsgColor,
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
			Image: &discordgo.MessageEmbedImage{
				URL: embed.Image,
			},
		}
	}
	_, err := d.session.ChannelMessageSendEmbed(d.mochiActivityChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("[svc.Discord.SendUpvoteMessage] - ChannelMessageSendEmbed failed to channel %s: %s", d.mochiActivityChannelID, err.Error())
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

// Reply to the lastest $vote message
func (d *Discord) ReplyUpvoteMessage(msg *response.SetUpvoteMessageCacheResponse, source string) error {
	embed := util.GenerateUpvoteMessage(msg.UserID, source)
	voteRemindStr := "\n\nCheck your progress and vote for Mochi with `$vote`"
	if msg.GuildID != d.mochiGuildID {
		voteRemindStr += "\n[Join Mochi server](https://discord.gg/FUDwZ2GqnN) <:threat:1019815998116859965>"
	}
	msgEmbed := discordgo.MessageEmbed{
		Title:       embed.Title,
		Description: embed.Description + voteRemindStr,
		Color:       mochiUpvoteMsgColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
		Image: &discordgo.MessageEmbedImage{
			URL: embed.Image,
		},
	}
	_, err := d.session.ChannelMessageSendEmbedReply(msg.ChannelID, &msgEmbed, &discordgo.MessageReference{
		MessageID: msg.MessageID,
		ChannelID: msg.ChannelID,
		GuildID:   msg.GuildID,
	})
	if err != nil {
		d.log.Error(err, "[discord.ReplyUpvoteMessage] failed to reply")
		return err
	}
	return nil
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

func (d *Discord) NotifyMemberLeave(req *request.MemberRemoveWebhookRequest, jlChannelID string) error {
	msg := &discordgo.MessageEmbed{
		Title:       "Say goodbye :wave:",
		Description: fmt.Sprintf("%s has left the server :wave:", req.Username),
		Color:       mochiErrorColor,
		Footer:      &discordgo.MessageEmbedFooter{Text: "Leaving", IconURL: req.Avatar},
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendEmbed(jlChannelID, msg)
	if err != nil {
		d.log.Fields(logger.Fields{"req": req}).Error(err, "session.ChannelMessageSendEmbed() failed")
		return err
	}
	return nil
}

func (d *Discord) NotifyMemberJoin(discordID, avatar, jlChannelID string, userCount int64) error {
	postfix := util.NumberPostfix(int(userCount))
	msg := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Welcome the %v%s member of your server :tada:", userCount, postfix),
		Footer:      &discordgo.MessageEmbedFooter{Text: "Onboarding", IconURL: avatar},
		Description: fmt.Sprintf("<@%s> has just joined your server. Give a heartwarming welcome :wave:", discordID),
		Color:       mochiSuccessColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendEmbed(jlChannelID, msg)
	if err != nil {
		d.log.Fields(logger.Fields{"discordID": discordID, "JLChannelID": jlChannelID}).Error(err, "session.ChannelMessageSendEmbed() failed")
		return err
	}
	return nil
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

func (d *Discord) SendTipActivityLogs(channelID, userID, title, description, image string) error {
	if channelID == "" {
		return nil
	}
	dcUser, err := d.session.User(userID)
	if err != nil {
		d.log.Errorf(err, "[SendTipActivityLogs] - get discord user failed %s", userID)
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
		Footer: &discordgo.MessageEmbedFooter{
			Text: "👉 You can say thank to your friend by $tip",
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

func (d *Discord) CreateDiscussionChannelForProposal(guildId, proposalChannelID, proposalTitle string) (string, error) {
	proposalChannel, err := d.Channel(proposalChannelID)
	if err != nil {
		d.log.Fields(logger.Fields{
			"proposalChannelID": proposalChannelID,
		}).Error(err, "[discord.CreateDiscussionChannelForProposal] get channel failed")
		return "", errors.ErrInvalidDiscordChannelID
	}
	discussChannelCreateData := discordgo.GuildChannelCreateData{
		Name:                 proposalTitle,
		Type:                 discordgo.ChannelTypeGuildText,
		PermissionOverwrites: proposalChannel.PermissionOverwrites,
		ParentID:             proposalChannel.ParentID,
	}
	discussionChannel, err := d.CreateChannel(guildId, discussChannelCreateData)
	if err != nil {
		d.log.Fields(logger.Fields{"guildId": guildId}).Error(err, "CreateDiscussionChannelForProposal - GuildChannelCreate failed")
		return "", err
	}
	return discussionChannel.ID, nil
}

func (d *Discord) NotifyNewProposal(channelID string, proposal response.SnapshotProposalDataResponse) error {
	body := proposal.Proposal.Body
	if len(body) > 250 {
		body = body[0:249] + "..."
	}
	title := proposal.Proposal.Title
	if len(title) > 70 {
		title = title[0:69] + "..."
	}
	// remove image file name
	reg := regexp.MustCompile(`[a-zA-Z]*(\.png|\.jpeg|\.jpg)`)
	res := reg.ReplaceAllString(body, " ")
	msgEmbed := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("<:mail:1058304339237666866> %s", title),
		Description: fmt.Sprintf("%s\n\n<:social:933281365586227210> Vote [here](https://snapshot.org/#/%s/proposal/%s)\n<:transaction:933341692667506718> Voting will close at: <t:%d>", res, proposal.Proposal.Space.ID, proposal.Proposal.ID, proposal.Proposal.End),
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	_, err := d.session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content: "> @everyone",
		Embed:   &msgEmbed,
	})
	if err != nil {
		d.log.Error(err, "[discord.NotifyNewProposal] d.session.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (d *Discord) NotifyNewCommonwealthDiscussion(req request.NewCommonwealthDiscussionRequest) error {
	body := req.Discussion.Plaintext
	if len(body) > 250 {
		body = body[0:249] + "..."
	}
	// remove image file name from body
	reg := regexp.MustCompile(`[a-zA-Z]*(\.png|\.jpeg|\.jpg)`)
	description := reg.ReplaceAllString(body, " ")
	// remove - from title
	regT := regexp.MustCompile(`%20`)
	title := regT.ReplaceAllString(body, " ")
	if len(title) > 70 {
		title = title[0:69] + "..."
	}
	discussionLink := fmt.Sprintf("https://commonwealth.im/%s/%s/%d", req.Discussion.Chain, req.Discussion.Kind, req.Discussion.ID)
	msgSend := discordgo.MessageSend{
		Content: "> @everyone",
		Embeds: []*discordgo.MessageEmbed{
			{
				Author: &discordgo.MessageEmbedAuthor{
					Name:    req.Community.Name,
					URL:     req.Community.Website,
					IconURL: req.Community.IconURL,
				},
				Title:       title,
				Description: description,
				Color:       mochiLogColor,
				Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
				Footer: &discordgo.MessageEmbedFooter{
					Text:    "commonwealth.im",
					IconURL: "https://pbs.twimg.com/profile_images/1562880197376020480/6R_gefq8_400x400.jpg",
				},
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Join Thread",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: fmt.Sprintf("proposal_join_thread_commonwealth-%s-%d", req.Community.CommunityID, req.Discussion.ID),
						Emoji: discordgo.ComponentEmoji{
							Name: "conversation",
							ID:   "1078633892493410334",
						},
					},
					discordgo.Button{
						Label:    "Go To The Discussion",
						Style:    discordgo.LinkButton,
						Disabled: false,
						URL:      discussionLink,
					},
				},
			},
		},
	}
	_, err := d.session.ChannelMessageSendComplex(req.ChannelID, &msgSend)
	if err != nil {
		d.log.Error(err, "[discord.NotifyNewProposal] d.session.ChannelMessageSendEmbed() failed")
	}
	return err
}

func (d *Discord) SendDMUserPriceAlert(userID, symbol string, alertType model.AlertType, price float64) error {
	var description string
	switch alertType {
	case model.PriceReaches:
		description = fmt.Sprintf("%v reaches %v", symbol, price)
	case model.PriceDropsTo:
		description = fmt.Sprintf("%v is under %v", symbol, price)
	case model.PriceRisesAbove:
		description = fmt.Sprintf("%v rises above %v", symbol, price)
	case model.ChangeIsOver:
		description = fmt.Sprintf("%v is up by %v%%", symbol, price)
	case model.ChangeIsUnder:
		description = fmt.Sprintf("%v is down by %v%%", symbol, price)
	}
	privChan, err := d.session.UserChannelCreate(userID)
	if err != nil {
		d.log.Error(err, "[discord.SendDMUserPriceAlert] d.session.UserChannelCreate() failed")
		return err
	}
	msg := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Your price alert is triggered",
			Description: description,
			Color:       mochiLogColor,
			Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
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
