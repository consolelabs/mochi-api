package model

import "time"

type ProductTheme struct {
	Id        int64     `json:"id"`
	Image     string    `json:"image"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
