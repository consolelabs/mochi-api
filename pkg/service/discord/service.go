package discord

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type Service interface {
	NotifyAddNewCollection(guildID string, collectionName string, symbol string, chain string, image string) error
	NotifyStealFloorPrice(price float64, floor float64, url string, name string, image string) error
	NotifyStealAveragePrice(price float64, floor float64, url string, name string, image string) error

	// moderation logs
	NotifyNewGuild(newGuildID string, count int) error
	NotifyMemberJoin(discordID, avatar, jlChannelID string, userCount int64) error
	NotifyMemberLeave(req *request.MemberRemoveWebhookRequest, jlChannelId string) error
	SendUpdateRolesLog(guildID, logChannelID, userID, roleID, _type string) error
	SendFeedback(req *request.UserFeedbackRequest, feedbackID string) error
	SendGuildActivityLogs(channelID, userID, title, description string) error
	SendLevelUpMessage(logChannelID, role string, uActivity *response.HandleUserActivityResponse)
	NotifyGmStreak(channelID string, userDiscordID string, streakCount int, podTownXps model.CreateUserTxResponse) error
	SendUpvoteMessage(discordID, source string, isStranger bool) error
	ReplyUpvoteMessage(msg *response.SetUpvoteMessageCacheResponse, source string) error
	NotifyGuildDelete(guildID, guildName, iconURL string, guildsLeft int) error
	SendTipActivityLogs(channelID, userID, title, description, image string) error
}
