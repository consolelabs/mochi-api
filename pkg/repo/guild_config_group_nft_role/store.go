package guildconfiggroupnftrole

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(config model.GuildConfigGroupNFTRole) (*model.GuildConfigGroupNFTRole, error)
	ListByGuildID(guildID string) ([]model.GuildConfigGroupNFTRole, error)
	Delete(id string) error
	GetByRoleID(guildID, roleID string) (*model.GuildConfigGroupNFTRole, error)
}
