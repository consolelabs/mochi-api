package job

import (
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service"
)

type cleanUpExpiredAssignContract struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewCleanUpExpiredAssignContractJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &cleanUpExpiredAssignContract{
		entity:  e,
		service: svc,
		log:     l,
	}
}

func (job *cleanUpExpiredAssignContract) Run() error {
	return job.entity.OffchainTipBotDeleteExpiredAssignContract()
}
