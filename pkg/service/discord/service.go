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
	NotifyMemberJoin(discordID, avatar, jlChannelID string, userCount int64) error
	NotifyMemberLeave(req *request.MemberRemoveWebhookRequest, jlChannelId string) error
	SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type string) error
	SendFeedback(req *request.UserFeedbackRequest, feedbackID string) error
	SendGuildActivityLogs(channelID, userID, title, description string) error
	SendLevelUpMessage(levelUpConfig *model.GuildConfigLevelupMessage, role string, uActivity *response.HandleUserActivityResponse)
	NotifyGmStreak(channelID string, userDiscordID string, streakCount int, podTownXps model.CreateUserTxResponse) error
	SendUpvoteMessage(discordID, source string, isStranger bool) error
	ReplyUpvoteMessage(msg *response.SetUpvoteMessageCacheResponse, source string) error
	NotifyGuildDelete(guildID, guildName, iconURL string, guildsLeft int) error
	SendTipActivityLogs(channelID, userID, title, description, image string) error

	// channel interaction
	Channel(channelID string) (*discordgo.Channel, error)
	CreateChannel(guildID string, createData discordgo.GuildChannelCreateData) (*discordgo.Channel, error)
	DeleteChannel(channelId string) error

	// DAO voting
	SendMessage(channelID string, msgSend discordgo.MessageSend) error
	NotifyNewProposal(channelID string, proposal response.SnapshotProposalDataResponse) error
	CreateDiscussionChannelForProposal(guildId, proposalChannelID, proposalTitle string) (string, error)
}
