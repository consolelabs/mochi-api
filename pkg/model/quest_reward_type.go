package model

import "github.com/google/uuid"

type QuestRewardType struct {
	ID   uuid.UUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	Name string    `json:"name"`
}

func (QuestRewardType) TableName() string {
	return "quests_reward_types"
}
