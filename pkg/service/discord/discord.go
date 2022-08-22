package discord

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Discord struct {
	session            *discordgo.Session
	log                logger.Logger
	mochiLogChannelID  string
	mochiSaleChannelID string
}

const (
	mochiLogColor = 0xFCD3C1
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
		session:            discord,
		log:                log,
		mochiLogChannelID:  cfg.MochiLogChannelID,
		mochiSaleChannelID: cfg.MochiSaleChannelID,
	}, nil
}

func (d *Discord) NotifyNewGuild(guildID string, count int) error {
	// get new guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild info: %w", err)
	}
	postfix := "th"
	switch count % 10 {
	case 1:
		postfix = "st"
	case 2:
		postfix = "nd"
	case 3:
		postfix = "rd"
	default:
		postfix = "th"
	}
	msgEmbed := discordgo.MessageEmbed{
		Title:       "Mochi has joined new Guild!",
		Description: fmt.Sprintf("**%s**, the %v%s guild", guild.Name, count, postfix),
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err = d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, &msgEmbed)
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

func (d *Discord) SendUpdateRoleMessage(logChannelID, curRoleID, oldRoleID string, uActivity *response.HandleUserActivityResponse) {
	if uActivity.ChannelID == "" && logChannelID == "" {
		d.log.Info("Action was not performed at any channel and no log channel configured as well")
		return
	}
	channelID := logChannelID
	if channelID == "" {
		channelID = uActivity.ChannelID
	}

	dcUser, err := d.session.User(uActivity.UserID)
	if err != nil {
		d.log.Errorf(err, "SendUpdateRoleMessage - failed to get discord user %s", uActivity.UserID)
		return
	}

	curRole, err := d.session.State.Role(uActivity.GuildID, curRoleID)
	if err != nil {
		d.log.Errorf(err, "SendUpdateRoleMessage - failed to get discord roleID %s", curRoleID)
		return
	}
	var oldRole = &discordgo.Role{}
	if oldRoleID == "" {
		oldRole.Name = "N/A"
	} else {
		oldRole, err = d.session.State.Role(uActivity.GuildID, oldRoleID)
		if err != nil {
			d.log.Errorf(err, "SendUpdateRoleMessage - failed to get discord roleID %s", curRoleID)
			return
		}
	}

	description := fmt.Sprintf("<@%s> has been updated role **(%s - %s)**\n\n**XP: **%d\n**Role: **%s", uActivity.UserID, oldRole.Name, curRole.Name, uActivity.CurrentXP, curRole.Name)
	msgEmbed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Role update!",
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
		d.log.Errorf(err, "SendUpdateRoleMessage - failed to send level up msg")
		return
	}
}
