package mochiprofile

type Service interface {
	GetByDiscordID(discordID string, noFetchAmount bool) (*GetProfileResponse, error)
	GetApiKeyByProfileID(profileID string) (*ProfileApiKeyResponse, error)
	CreateProfileApiKey(profileAccessToken string) (*ProfileApiKeyResponse, error)
	GetByID(profileID, authorization string) (*GetProfileResponse, error)
	GetByIds(profileIds []string) ([]GetProfileResponse, error)
	GetByDiscordIds(discordIds []string) ([]GetProfileResponse, error)
	GetAllEvmAccount() ([]*EvmAssociatedAccount, error)
	AssociateDex(profileId, platform, apiKey, apiSecret string) error
	UnlinkDex(profileId, platform string) error
	GetOnboardingStatus(profileId string) (res *OnboardingStatusResponse, err error)
	MarkUserDidOnboarding(profileId string) error
	GetByTelegramID(telegramID string, noFetchAmount bool) (*GetProfileResponse, error)
	GetProfileActivities(profileID string) (any, error)
	GetAllBinanceAccount() (*GetAllBinanceAssociatedAccountData, error)
	GetTopActiveUsers(top int) ([]GetProfileResponse, error)
}
