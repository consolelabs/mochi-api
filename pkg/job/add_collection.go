package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
)

type addCollection struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewAddCollectionJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &addCollection{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *addCollection) Run() error {
	return job.entity.AddCollection()
}
