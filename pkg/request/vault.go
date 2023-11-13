package request

type CreateVaultRequest struct {
	VaultCreator string `json:"vault_creator" binding:"required"`
	GuildId      string `json:"guild_id"`
	Name         string `json:"name"`
	Threshold    string `json:"threshold"`
	DesigMode    bool   `json:"desig_mode" form:"size,default=false"`
}

type CreateVaultConfigChannelRequest struct {
	GuildId   string `json:"guild_id" binding:"required"`
	ChannelId string `json:"channel_id" binding:"required"`
}

type CreateConfigThresholdRequest struct {
	GuildId   string `json:"guild_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Threshold string `json:"threshold" binding:"required"`
}

type AddTreasurerToVaultRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	VaultId       int64  `json:"vault_id" binding:"required"`
	UserProfileID string `json:"user_profile_id" binding:"required"`
	ChannelId     string `json:"channel_id" binding:"required"`
}

type RemoveTreasurerToVaultRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	VaultId       int64  `json:"vault_id" binding:"required"`
	UserProfileID string `json:"user_profile_id" binding:"required"`
	ChannelId     string `json:"channel_id" binding:"required"`
}

type CreateTreasurerResultRequest struct {
	GuildId       string `json:"guild_id" binding:"required"`
	VaultId       int64  `json:"vault_id" binding:"required"`
	UserProfileID string `json:"user_profile_id"`
	ChannelId     string `json:"channel_id" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Status        string `json:"status" binding:"required"`
	Amount        string `json:"amount"`
	Chain         string `json:"chain"`
	Token         string `json:"token"`
	Address       string `json:"address"`
}

type CreateTreasurerRequest struct {
	GuildId            string `json:"guild_id" binding:"required"`
	RequesterProfileId string `json:"requester_profile_id" binding:"required"`
	VaultName          string `json:"vault_name" binding:"required"`
	UserProfileId      string `json:"user_profile_id"`
	Message            string `json:"message"`
	Type               string `json:"type" binding:"required"`
	Amount             string `json:"amount"`
	Chain              string `json:"chain"`
	Token              string `json:"token"`
	Address            string `json:"address"`
	MessageUrl         string `json:"message_url"`
	Platform           string `json:"platform"`
}

type TransferVaultTokenRequest struct {
	GuildId   string `json:"guild_id" binding:"required"`
	VaultId   int64  `json:"vault_id" binding:"required"`
	RequestId int64  `json:"request_id" binding:"required"`
	Address   string `json:"address"`
	Amount    string `json:"amount" binding:"required"`
	Token     string `json:"token" binding:"required"`
	Chain     string `json:"chain" binding:"required"`
	Target    string `json:"target"`
	Platform  string `json:"platform"`
}

type CreateTreasurerSubmission struct {
	Type              string `json:"type" binding:"required"`
	VaultId           int64  `json:"vault_id" binding:"required"`
	RequestId         int64  `json:"request_id" binding:"required"`
	SumitterProfileId string `json:"submitter_profile_id" binding:"required"`
	Choice            string `json:"choice" binding:"required"`
}

type MochiPayVaultRequest struct {
	ProfileId   string   `json:"profile_id"`
	PrivateKey  string   `json:"private_key"`
	To          string   `json:"to"`
	Amount      string   `json:"amount"`
	Token       string   `json:"token"`
	Chain       string   `json:"chain"`
	Name        string   `json:"name"`
	VaultId     int64    `json:"vault_id"`
	Reciever    string   `json:"receiver"`
	Message     string   `json:"message"`
	ListNotify  []string `json:"list_notify"`
	RequestId   int64    `json:"request_id"`
	Platform    string   `json:"platform"`
	MesssageUrl string   `json:"message_url"`
}

type GetVaultsRequest struct {
	GuildID       string `form:"guild_id"`
	ProfileID     string `form:"profile_id"`
	EvmAddress    string `form:"evm_address"`
	SolanaAddress string `form:"solana_address"`
	Threshold     string `form:"threshold"`
	NoFetchAmount string `form:"no_fetch_amount" default:"false"`
}

type GetVaultRequest struct {
	VaultId       string `uri:"vault_id"`
	NoFetchAmount string `form:"no_fetch_amount" default:"false"`
}
