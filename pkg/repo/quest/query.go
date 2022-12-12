package quest

import (
	"github.com/google/uuid"

	"github.com/defipod/mochi/pkg/model"
)

type ListQuery struct {
	ID         *uuid.UUID
	Action     string
	NotActions []model.QuestAction
	Routine    string
	Sort       string
}
