package request

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type GetUserQuestListRequest struct {
	UserID  string             `json:"user_id" form:"user_id" binding:"required"`
	Routine model.QuestRoutine `json:"routine" form:"routine,default=daily"`
}

type GenerateUserQuestListRequest struct {
	UserID    string             `json:"user_id"`
	Routine   model.QuestRoutine `json:"routine"`
	Quantity  int                `json:"quantity"` // number of quests needs to be generated
	StartTime time.Time          `json:"start_time"`
}

type ClaimQuestsRewardsRequest struct {
	UserID  string             `json:"user_id"`
	Routine model.QuestRoutine `json:"routine"`
}

type UpdateQuestProgressRequest struct {
	GuildID *string           `json:"guild_id"`
	UserID  string            `json:"user_id"`
	Action  model.QuestAction `json:"action"`
}
