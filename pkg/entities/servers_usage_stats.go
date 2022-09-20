package entities

import "github.com/defipod/mochi/pkg/request"

func (e *Entity) AddServersUsageStats(info *request.UsageInformation) error {
	return e.repo.ServersUsageStats.CreateOne(info)
}
