package mochiprofile

type Service interface {
	GetByDiscordID(discordID string) (*GetProfileResponse, error)
	GetApiKeyByProfileID(profileID string) (*ProfileApiKeyResponse, error)
	CreateProfileApiKey(profileAccessToken string) (*ProfileApiKeyResponse, error)
	GetByID(profileID string) (*GetProfileResponse, error)
}
