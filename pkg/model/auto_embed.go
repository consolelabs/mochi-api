package model

import (
	"time"
)

type AutoEmbed struct {
	Id          int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	AuthorId    string    `json:"type_id"`
	ActionId    string    `json:"action_id"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Fields      string    `json:"fields"`
	Thumbnail   string    `json:"thumbnail"`
	Type        string    `json:"type"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`

	Image  *AutoEmbedImage  `json:"auto_embed_image" gorm:"foreignKey:EmbedId"`
	Video  *AutoEmbedVideo  `json:"auto_embed_video" gorm:"foreignKey:EmbedId"`
	Footer *AutoEmbedFooter `json:"auto_embed_footer" gorm:"foreignKey:EmbedId"`
}

type AutoEmbedImage struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	EmbedId   string    `json:"embed_id"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	ProxyUrl  string    `json:"icon_url"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type AutoEmbedVideo struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	EmbedId   string    `json:"embed_id"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type AutoEmbedFooter struct {
	Id        int64     `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	EmbedId   string    `json:"embed_id"`
	Content   string    `json:"content"`
	Url       string    `json:"url"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}
