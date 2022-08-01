package discord

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	NotifyNewGuild(newGuildID string, count int) error
	NotifyAddNewCollection(guildID string, collectionName string, symbol string, chain string, image string) error
	SendGuildActivityLogs(channelID, userID, title, description string) error
	SendLevelUpMessage(logChannelID, role string, uActivity *response.HandleUserActivityResponse)
	NotifyStealFloorPrice(price float64, floor float64, url string, name string, image string) error
}
