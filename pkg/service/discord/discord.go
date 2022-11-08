package discord

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
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
	if channelID == "" {
		return nil
	}

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

func (d *Discord) SendLevelUpMessage(logChannelID, role string, uActivity *response.HandleUserActivityResponse) {
	if !uActivity.LevelUp {
		return
	}
	if uActivity.ChannelID == "" && logChannelID == "" {
		d.log.Info("Action was not performed at any channel and no log channel configured as well")
		return
	}
	channelID := logChannelID
	if channelID == "" {
		channelID = uActivity.ChannelID
	}
	if role == "" {
		role = "N/A"
	}

	dcUser, err := d.session.User(uActivity.UserID)
	if err != nil {
		d.log.Errorf(err, "SendLevelUpMessage - failed to get discord user %s", uActivity.UserID)
		return
	}

	description := fmt.Sprintf("<@%s> has leveled up **(%d - %d)**\n\n**XP: **%d\n**Role: **%s", uActivity.UserID, uActivity.CurrentLevel-1, uActivity.CurrentLevel, uActivity.CurrentXP, role)
	msgEmbed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Level up!",
			IconURL: "https://cdn.discordapp.com/emojis/984824963112513607.png?size=240&quality=lossless",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: dcUser.AvatarURL(""),
		},
		Description: description,
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

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
						Label:    "Set in progress",
						Style:    1,
						CustomID: fmt.Sprintf("handle-feedback-set-in-progress_%s", feedbackID),
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
