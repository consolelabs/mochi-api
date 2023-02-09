package request

type CreateGuildMixRole struct {
	GuildID          string                   `json:"guild_id" binding:"required"`
	RoleID           string                   `json:"role_id" binding:"required"`
	NFTRequirement   *MixRoleNFTRequirement   `json:"nft_requirement"`
	TokenRequirement *MixRoleTokenRequirement `json:"token_requirement"`
	RequiredLevel    int                      `json:"required_level"`
}

type MixRoleNFTRequirement struct {
	NftID  string `json:"nft_id" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type MixRoleTokenRequirement struct {
	TokenID int     `json:"token_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}
