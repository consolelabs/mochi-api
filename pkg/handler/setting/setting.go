package setting

import (
	"github.com/sirupsen/logrus"

	"github.com/defipod/mochi/pkg/entities"
	mlogger "github.com/defipod/mochi/pkg/logger"
)

type handler struct {
	entities *entities.Entity
	logger   mlogger.Logger
}

func New(e *entities.Entity, l mlogger.Logger) IHandler {
	return &handler{
		entities: e,
		logger:   l,
	}
}

var (
	logger = logrus.WithFields(logrus.Fields{
		"component": "handler.setting",
	})
)
