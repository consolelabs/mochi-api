package quest

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type ListQuery struct {
	ID        *uuid.UUID
	Action    *model.QuestAction
	NotAction *model.QuestAction
	Routine   *model.QuestRoutine
}
