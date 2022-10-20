package quest

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type ListQuery struct {
	ID         *uuid.UUID
	Action     string
	NotActions []model.QuestAction
	Routine    string
	Sort       string
}
