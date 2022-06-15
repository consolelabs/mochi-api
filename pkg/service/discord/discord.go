package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/config"
)

type Discord struct {
	session           *discordgo.Session
	mochiLogChannelID string
}

const (
	mochiLogColor = 0xFCD3C1
)

func NewService(
	cfg config.Config,
) (Service, error) {
	// *** discord ***
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}
	return &Discord{
		session:           discord,
		mochiLogChannelID: cfg.MochiLogChannelID,
	}, nil
}

func (d *Discord) NotifyNewGuild(guildID string) error {
	// get new guild info
	guild, err := d.session.Guild(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild info: %w", err)
	}

	msgEmbed := discordgo.MessageEmbed{
		Title:       "Mochi has joined new Guild!",
		Description: fmt.Sprintf("**%s** (%s)", guild.Name, guild.ID),
		Color:       mochiLogColor,
	}

	_, err = d.session.ChannelMessageSendEmbed(d.mochiLogChannelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (d *Discord) SendGuildActivityLogs(channelID, title, description string) error {
	if channelID == "" {
		return nil
	}

	msgEmbed := discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       mochiLogColor,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	_, err := d.session.ChannelMessageSendEmbed(channelID, &msgEmbed)
	if err != nil {
		return fmt.Errorf("failed to send activity logs to channel %s: %w", channelID, err)
	}

	return nil
}
