package model

type UserEarn struct {
	Id         int64  `json:"id"`
	UserId     string `json:"user_id"`
	EarnId     int    `json:"earn_id"`
	Status     string `json:"status"`
	IsFavorite bool   `json:"is_favorite"`

	Earn *EarnInfo `json:"earn,omitempty" gorm:"ForeignKey:EarnId;references:Id"`
}
