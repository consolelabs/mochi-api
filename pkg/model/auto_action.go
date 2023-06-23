package model

import (
	"time"
)

type AutoAction struct {
	Id           int64     `json:"id" swaggertype:"string"`
	TriggerId    int64     `json:"trigger_id" gorm:"not null"`
	TypeId       string    `json:"type_id" gorm:"not null"`
	ChannelId    string    `json:"channel_id"`
	ThenActionId string    `json:"then_action_id"`
	Index        int       `json:"index"`
	Platform     string    `json:"platform"`
	Content      string    `json:"content"`
	IsPrimary    bool      `json:"is_primary"`
	CreatedAt    time.Time `json:"created_at"`
	ActionData   string    `json:"action_data"`
	LimitPerUser int       `json:"limit_per_user"`

	Type       AutoType    `json:"auto_type" gorm:"foreignKey:TypeId;references:Id"`
	Embed      AutoEmbed   `json:"auto_embed" gorm:"foreignKey:Id;references:ActionId"`
	ThenAction *AutoAction `json:"then_action" gorm:"foreignKey:ThenActionId;references:Id"`
}

type AutoActionHistory struct {
	Id        int64     `json:"id" swaggertype:"string"`
	TriggerId int64     `json:"trigger_id" gorm:"not null"`
	ActionId  int64     `json:"action_id" gorm:"not null"`
	UserId    string    `json:"user_id" gorm:"not null"`
	MessageId string    `json:"message_id" gorm:"not null"`
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"created_at"`

	Action AutoAction `json:"auto_action" gorm:"foreignKey:ActionId;references:Id"`
}
type AutoTransferVaultTokenRequest struct {
	GuildId string `json:"guild_id" binding:"required"`
	VaultId int64  `json:"vault_id" binding:"required"`
	Address string `json:"address"`
	Amount  string `json:"amount" binding:"required"`
	Token   string `json:"token" binding:"required"`
	Chain   string `json:"chain" binding:"required"`
	Target  string `json:"target"`
	Message string `json:"message"`
}
