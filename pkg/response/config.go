package response

import "github.com/defipod/mochi/pkg/model"

type GuildNFTRolesResponse struct {
	model.GuildConfigNFTRole
	RoleName      string              `json:"role_name"`
	Color         int                 `json:"color"`
	NFTCollection model.NFTCollection `json:"nft_collection"`
}
