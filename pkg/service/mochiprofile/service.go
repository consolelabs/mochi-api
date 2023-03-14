package mochiprofile

type Service interface {
	GetByDiscordID(discordID string) (*GetProfileByDiscordResponse, error)
}
