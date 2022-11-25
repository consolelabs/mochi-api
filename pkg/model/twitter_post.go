package model

import (
	"time"

	"gorm.io/gorm"
)

type TwitterPost struct {
	TwitterID     string    `json:"twitter_id"`
	TwitterHandle string    `json:"twitter_handle"`
	TweetID       string    `json:"tweet_id"`
	GuildID       string    `json:"guild_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
}

func (s *TwitterPost) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now().UTC()
	return nil
}
