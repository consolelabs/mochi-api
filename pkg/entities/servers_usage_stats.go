package entities

import (
	"encoding/json"

	"github.com/defipod/mochi/pkg/model"
)

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
	return e.kafka.Produce(e.cfg.Kafka.Topic, "webhook.gitbook", b)
}
