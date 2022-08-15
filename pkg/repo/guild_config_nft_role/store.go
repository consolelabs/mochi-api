package guild_config_nft_role

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetMemberCurrentRoles(guildID string) ([]model.MemberNFTRole, error)
	ListByGuildID(guildID string) ([]model.GuildConfigNFTRole, error)
	GetByRoleID(guildID, roleID string) (*model.GuildConfigNFTRole, error)
	UpsertOne(config *model.GuildConfigNFTRole) error
	Update(config *model.GuildConfigNFTRole) error
	Delete(id string) error
}
