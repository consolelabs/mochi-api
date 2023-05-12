package mochiprofile

type Service interface {
	GetByDiscordID(discordID string) (*GetProfileByDiscordResponse, error)
	GetApiKeyByProfileID(profileID string) (*ProfileApiKeyResponse, error)
	CreateProfileApiKey(profileAccessToken string) (*ProfileApiKeyResponse, error)
}
