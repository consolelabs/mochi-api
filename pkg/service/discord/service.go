package discord

type Service interface {
	NotifyNewGuild(newGuildID string) error
}
