package model

type ConfigXpLevel struct {
	Level int `json:"level" gorm:"primaryKey"`
	MinXP int `json:"min_xp"`
}
