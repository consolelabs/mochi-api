package guildconfiglogchannel

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(model *model.GuildConfigLogChannel) (*model.GuildConfigLogChannel, error)
	Get(query Query) (models []model.GuildConfigLogChannel, err error)
}
