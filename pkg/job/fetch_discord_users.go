package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
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

	createUserRequests := make([]request.CreateUserRequest, 0)

	// fetch users from discord
	for _, guild := range guilds.Data {
		guildUsers, err := j.entity.GetGuildUsersFromDiscord(guild.ID)
		if err != nil {
			j.log.Error(err, "failed to get guild users")
			continue
		}

		j.log.Fields(logger.Fields{"guild": guild.ID, "users": len(guildUsers)}).Infof("fetched guild users")

		for _, user := range guildUsers {
			createUserRequests = append(createUserRequests, request.CreateUserRequest{
				ID:       user.User.ID,
				Username: user.User.Username,
				Nickname: user.Nickname,
				GuildID:  user.GuildID,
			})
		}
	}

	j.log.Fields(logger.Fields{"users": len(createUserRequests)}).Infof("creating users")

	// create users
	for _, req := range createUserRequests {
		if _, err := j.entity.GetOneOrUpsertUser(req.ID); err != nil {
			j.log.Fields(logger.Fields{"user": req}).Error(err, "failed to create user")
			continue
		}

		if err := j.entity.CreateGuildUserIfNotExists(req.GuildID, req.ID, req.Nickname); err != nil {
			j.log.Fields(logger.Fields{"user": req}).Error(err, "failed to create guild user")
			continue
		}
	}
	return nil
}
