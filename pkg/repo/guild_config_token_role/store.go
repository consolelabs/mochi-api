package guildconfigtokenrole

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(*model.GuildConfigTokenRole) error
	Get(id int) (model *model.GuildConfigTokenRole, err error)
	ListByGuildID(guildID string) ([]model.GuildConfigTokenRole, error)
	Update(*model.GuildConfigTokenRole) error
	Delete(id int) error
	ListAllTokenConfigs(guildID string) ([]model.Token, error)
	ListConfigGuildIds() ([]string, error)
}
