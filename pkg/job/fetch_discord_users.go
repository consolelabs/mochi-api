package job

import "github.com/defipod/mochi/pkg/entities"

type fetchDiscordUsersJob struct {
	entity *entities.Entity
}

func New(entity *entities.Entity) Job {
	return &fetchDiscordUsersJob{
		entity: entity,
	}
}

func (job *fetchDiscordUsersJob) Run() error {

	return nil
}
