package entities

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
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

func (e *Entity) TotalCommandUsage(guildId string) (*response.Metric, error) {
	totalUsage, err := e.repo.ServersUsageStats.TotalUsage()
	if err != nil {
		e.log.Error(err, "[entities.TotalCommandUsage] - failed to get total usage")
		return nil, err
	}

	serverUsage, err := e.repo.ServersUsageStats.TotalUsageByGuildId(guildId)
	if err != nil {
		e.log.Error(err, "[entities.TotalCommandUsage] - failed to get server usage")
		return nil, err
	}

	return &response.Metric{
		TotalCommandUsage:  totalUsage,
		ServerCommandUsage: serverUsage,
	}, nil
}
