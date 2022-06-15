package model

type UserBalance struct {
	UserID  string  `json:"user_id"`
	TokenID int     `json:"token_id"`
	Balance float64 `json:"balance"`
	Token   *Token  `json:"token" gorm:"foreignKey:TokenID"`
}
