package entities

import (
	"encoding/json"

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
	kafkaMsg := model.KafkaMessage{
		Platform: "mochi-api",
		Gitbook: model.GitbookClick{
			Command:    cmd,
			Subcommand: action,
		},
	}
	b, err := json.Marshal(kafkaMsg)
	if err != nil {
		e.log.Error(err, "[entities.AddGitbookClick] - failed to marshal kafka message")
		return err
	}
	return e.kafka.Produce(e.cfg.Kafka.Topic, e.cfg.Kafka.Key, b)
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
