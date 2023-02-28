package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
)

type fetchDiscordUsers struct {
	entity *entities.Entity
	log    logger.Logger
}

func NewFetchDiscordUsersJob(e *entities.Entity, l logger.Logger) Job {
	return &fetchDiscordUsers{
		entity: e,
		log:    l,
	}
}

func (j *fetchDiscordUsers) Run() error {
	// get guilds from db
	guilds, err := j.entity.GetGuilds()
	if err != nil {
		j.log.Error(err, "failed to get guilds")
		return err
	}

	for _, guild := range guilds.Data {
		users, err := j.entity.FetchAndSaveGuildMembers(guild.ID)
		if err != nil {
			j.log.Fields(logger.Fields{"guild": guild.ID}).Error(err, "entity.FetchAndSaveGuildMembers() failed")
			continue
		}
		j.log.Fields(logger.Fields{"guild": guild.ID, "users": users}).Error(err, "entity.FetchAndSaveGuildMembers() done")
	}
	return nil
}
