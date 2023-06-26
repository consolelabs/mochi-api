package auto_trigger

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetAutoTriggers(guildID string) ([]model.AutoTrigger, error)
	Update(trigger string, field string, value interface{}) error
	CountByGuildUser(guildId, userId string) (int64, error)
	FirstOrCreate(autoTrigger *model.AutoTrigger) error
	Create(autoTrigger *model.AutoTrigger) error
}
