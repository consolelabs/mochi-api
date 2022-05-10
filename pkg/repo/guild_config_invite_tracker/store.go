package guildconfiginvitetracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(guildID string) (*model.GuildConfigInviteTracker, error)
	Upsert(cmd *model.GuildConfigInviteTracker) error
}
