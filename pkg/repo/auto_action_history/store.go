package auto_action_history

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id string) (*model.AutoActionHistory, error)
	Update(id string, field string, value interface{}) error
	CountByTriggerActionUserMessage(triggerId int64, actionId int64, userId string, messageId string) (int64, error)
	FirstOrCreate(autoTrigger *model.AutoActionHistory) error
	Create(autoTrigger *model.AutoActionHistory) error
}
