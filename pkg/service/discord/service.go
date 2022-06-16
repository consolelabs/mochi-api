package discord

type Service interface {
	NotifyNewGuild(newGuildID string) error
	SendGuildActivityLogs(channelID, title, description string) error
}
