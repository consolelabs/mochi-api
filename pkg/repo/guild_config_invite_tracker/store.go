package guildconfiginvitetracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Upsert(cmd *model.GuildConfigInviteTracker) error
}
