package guildconfigtwitterblacklist

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(ListQuery) ([]model.GuildConfigTwitterBlacklist, error)
	Upsert(*model.GuildConfigTwitterBlacklist) error
	Delete(guildID, twitterID string) error
}
