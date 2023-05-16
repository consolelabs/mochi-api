package model

import "github.com/google/uuid"

type QuestReward struct {
	ID           uuid.UUID        `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	QuestID      *uuid.UUID       `json:"quest_id" swaggertype:"string"`
	RewardTypeID uuid.UUID        `json:"reward_type_id" swaggertype:"string"`
	RewardAmount int              `json:"reward_amount"`
	PassID       *uuid.UUID       `json:"pass_id" swaggertype:"string"`
	Quest        *Quest           `json:"quest" swaggerignore:"true"`
	RewardType   *QuestRewardType `json:"reward_type"`
}

func (QuestReward) TableName() string {
	return "quests_rewards"
}
