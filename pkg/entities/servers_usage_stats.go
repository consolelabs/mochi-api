package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) AddServersUsageStats(info *request.UsageInformation) error {
	return e.repo.ServersUsageStats.CreateOne(&model.UsageStat{
		GuildID: info.GuildID,
		UserID:  info.UserID,
		Command: info.Command,
		Args:    info.Args,
	})
}
