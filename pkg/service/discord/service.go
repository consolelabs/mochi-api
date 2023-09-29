package discord

import (
	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	// mochi logs
	NotifyAddNewCollection(guildID string, collectionName string, symbol string, chain string, image string) error
	NotifyStealFloorPrice(price float64, floor float64, url string, name string, image string) error
	NotifyStealAveragePrice(price float64, floor float64, url string, name string, image string) error
	NotifyCompleteCollectionIntegration(guildID string, collectionName string, symbol string, chain string, image string) error
	NotifyCompleteCollectionSync(guildID string, collectionName string, symbol string, chain string, image string) error

	// moderation logs
	NotifyNewGuild(newGuildID string, count int) error
	SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type string) error
	SendFeedback(req *request.UserFeedbackRequest, feedbackID string) error
	SendGuildActivityLogs(channelID, userID, title, description string) error
	SendLevelUpMessage(levelUpConfig *model.GuildConfigLogChannel, role string, levelRoleLevel int, randomTip string, uActivity *response.HandleUserActivityResponse)
	NotifyGmStreak(channelID string, userDiscordID string, streakCount int, podTownXps model.CreateUserTxResponse) error
	NotifyGuildDelete(guildID, guildName, iconURL string, guildsLeft int) error
	SendTipActivityLogs(channelID, userID string, author *discordgo.MessageEmbedAuthor, description, image string) error

	// channel interaction
	Channel(channelID string) (*discordgo.Channel, error)
	CreateChannel(guildID string, createData discordgo.GuildChannelCreateData) (*discordgo.Channel, error)
	DeleteChannel(channelId string) error

	// Price alert
	SendDMUserPriceAlert(userID, symbol string, alertType model.AlertType, price float64) error

	// Common func
	SendMessage(channelID string, msgSend discordgo.MessageSend) error
	SendDM(userID string, payload discordgo.MessageSend) error

	// Guild
	GetGuildMembers(guildID string) ([]*discordgo.Member, error)
	GetGuild(guildID string) (*discordgo.Guild, error)
	GetGuildRoles(guildID string) ([]*model.DiscordGuildRole, error)

	// User
	GetUser(userID string) (*discordgo.User, error)
	CreateDiscussionChannelForProposal(guildId, proposalChannelID, proposalTitle string) (string, error)

	// Emoji
	GetGuildEmojis() ([]*discordgo.Emoji, error)
}
