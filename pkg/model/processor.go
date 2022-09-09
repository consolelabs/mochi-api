package model

type CreateUserTransaction struct {
	Dapp   string     `json:"dapp"`
	Action string     `json:"action"`
	Data   UserTxData `json:"data"`
}

type UserTxData struct {
	TxHash             string  `json:"tx_hash,omitempty"`
	UserID             string  `json:"user_id,omitempty"`
	UserDiscordId      string  `json:"user_discord_id,omitempty"`
	UserWalletAddress  string  `json:"user_wallet_address,omitempty"`
	RecipientDiscordId string  `json:"recipient_discord_id,omitempty"`
	RecipientUserID    string  `json:"recipient_user_id,omitempty"`
	Guild              string  `json:"guild_id,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	Cryptocurrency     string  `json:"cryptocurrency,omitempty"`
	TokenPriceSymbol   string  `json:"token_price_symbol,omitempty"`
	PoolName           string  `json:"pool_name,omitempty"`
	PoolToken          string  `json:"pool_token,omitempty"`
	Link               string  `json:"link,omitempty"`
	TwitterAddress     string  `json:"twitter_address,omitempty"`
	NekoId             string  `json:"neko_id,omitempty"`
	NekoName           string  `json:"name,omitempty"`
	NekoNameBid        string  `json:"neko_name_bid,omitempty"`
	Time               string  `json:"time,omitempty"`
	StreakCount        int     `json:"streak_count,omitempty"`
	TotalCount         int     `json:"total_count,omitempty"`
	ContractAddress    string  `json:"contract_address,omitempty"`
}

type CreateUserTxResponse struct {
	Data UserXps `json:"data"`
}

type UserXps struct {
	FameXp             int64 `json:"fame_xp"`
	LoyaltyXp          int64 `json:"loyalty_xp"`
	NextFameXps        int64 `json:"next_fame_xps"`
	NextLoyaltyXps     int64 `json:"next_loyalty_xps"`
	NextReputationXps  int64 `json:"next_reputation_xps"`
	NobilityXp         int64 `json:"nobility_xp"`
	ReputationXp       int64 `json:"reputation_xp"`
	TotalFameXps       int64 `json:"total_fame_xps"`
	TotalLoyaltyXps    int64 `json:"total_loyalty_xps"`
	TotalNobilityXps   int64 `json:"total_nobility_xps"`
	TotalReputationXps int64 `json:"total_reputation_xps"`
}

type UserFactionXps struct {
	UserDiscordId string `json:"user_discord_id,omitempty"`
	FameXp        int64  `json:"fame_xp"`
	LoyaltyXp     int64  `json:"loyalty_xp"`
	NobilityXp    int64  `json:"nobility_xp"`
	ReputationXp  int64  `json:"reputation_xp"`
}

type UserFactionXpsMapping struct {
	ImperialXp int64 `json:"imperial_xp"`
	RebellioXp int64 `json:"rebellio_xp"`
	MerchantXp int64 `json:"merchant_xp"`
	AcademyXp  int64 `json:"academy_xp"`
}

type GetUserFactionXpsResponse struct {
	Data UserFactionXps `json:"data"`
}
