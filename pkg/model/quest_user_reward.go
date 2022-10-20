package model

import (
	"time"

	"github.com/google/uuid"
)

type QuestUserReward struct {
	UserID       string       `json:"user_id"`
	QuestID      *uuid.UUID   `json:"quest_id" swaggertype:"string"`
	RewardID     uuid.UUID    `json:"reward_id" swaggertype:"string"`
	RewardTypeID uuid.UUID    `json:"reward_type_id" swaggertype:"string"`
	RewardAmount int          `json:"reward_amount"`
	PassID       *uuid.UUID   `json:"pass_id" swaggertype:"string"`
	StartTime    *time.Time   `json:"start_time"`
	ClaimedAt    time.Time    `json:"claimed_at"`
	Reward       *QuestReward `json:"reward"`
}

func (QuestUserReward) TableName() string {
	return "quests_user_rewards"
}
