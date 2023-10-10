package model

import "time"

type TonoCommandPermission struct {
	ID                    int       `json:"id"`
	Code                  string    `json:"code"`
	DiscordPermissionFlag string    `json:"discord_permission_flag"`
	Description           string    `json:"description"`
	NeedDm                bool      `json:"need_dm"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
