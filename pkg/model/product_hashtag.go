package model

import "time"

type ProductHashtag struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Title         string    `json:"title"`
	Image         string    `json:"image"`
	Color         string    `json:"color"`
	TelegramTitle string    `json:"telegram_title"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
