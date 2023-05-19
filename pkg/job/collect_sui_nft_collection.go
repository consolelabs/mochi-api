package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
)

type collectSuiNftCollection struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewCollectSuiNftCollectionJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &collectSuiNftCollection{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *collectSuiNftCollection) Run() error {
	return job.entity.CreateBluemoveNFTCollectionBatch("9996")
}
