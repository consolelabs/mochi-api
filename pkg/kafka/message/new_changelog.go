package message

import (
	"time"

	"github.com/consolelabs/mochi-typeset/typeset"
)

type NewChangelog struct {
	Type                 typeset.NotificationType `json:"type"`
	NewChangelogMetadata NewChangelogMetadata     `json:"new_changelog_metadata"`
}

type NewChangelogMetadata struct {
	Product      string    `json:"product"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	GithubUrl    string    `json:"github_url"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	FileName     string    `json:"file_name"`
	IsExpired    bool      `json:"is_expired"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
