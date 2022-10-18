package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) AddServersUsageStats(info *request.UsageInformation) error {
	return e.repo.ServersUsageStats.CreateOne(&model.UsageStat{
		GuildID:         info.GuildID,
		UserID:          info.UserID,
		Command:         info.Command,
		Success:         info.Success,
		ExecutionTimeMs: info.ExecutionTimeMs,
		Args:            info.Args,
	})
}

func (e *Entity) AddGitbookClick(url, cmd, action string) error {
	info, err := e.repo.GitbookClickCollector.GetByCommandAndAction(cmd, action)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Error(err, "[entities.AddGitbookClick] - failed to get gitbook click info")
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return e.repo.GitbookClickCollector.UpsertOne(model.GitbookClickCollector{
			Command:     cmd,
			Action:      action,
			CountClicks: 1,
		})
	}
	info.CountClicks = info.CountClicks + 1
	return e.repo.GitbookClickCollector.UpsertOne(info)
}
