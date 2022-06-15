package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type updateScheduledEventStatus struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewUpdateScheduledEventStatusJob(e *entities.Entity, l logger.Logger) Job {
	return &updateScheduledEventStatus{
		entity: e,
		log:    l,
	}
}

func (c *updateScheduledEventStatus) Run() error {
	activityConfigs, err := c.entity.ListActivityConfigsByName("event_host")
	if err != nil {
		return err
	}

	if len(activityConfigs) == 0 {
		c.log.Infof("no guilds enable event_host activity")
		return nil
	}

	for _, config := range activityConfigs {
		err := c.entity.NewGuildScheduledEvents(config.GuildID)
		if err != nil {
			c.log.Errorf(err, "failed to get new scheduled events")
		}

		err = c.entity.UpdateGuildScheduledEventStatus(config.GuildID)
		if err != nil {
			c.log.Errorf(err, "failed to update scheduled event status of guild %s", config.GuildID)
		}
	}

	return nil
}
