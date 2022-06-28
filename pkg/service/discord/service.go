package discord

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	NotifyNewGuild(newGuildID string) error
	SendGuildActivityLogs(channelID, userID, title, description string) error
	SendLevelUpMessage(logChannelID, role string, uActivity *response.HandleUserActivityResponse)
}
