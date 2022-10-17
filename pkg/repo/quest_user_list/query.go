package questuserlist

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type ListQuery struct {
	UserID      *string
	QuestID     *uuid.UUID
	StartTime   *time.Time
	Routine     *model.QuestRoutine
	Action      *model.QuestAction
	NotAction   *model.QuestAction
	IsCompleted *bool
	IsClaimed   *bool
}
