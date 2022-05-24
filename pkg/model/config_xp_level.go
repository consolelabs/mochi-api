package model

type ConfigXpLevel struct {
	Level int    `json:"level" gorm:"primaryKey"`
	MinXP string `json:"min_xp"`
}
