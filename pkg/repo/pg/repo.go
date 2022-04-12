package pg

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/repo"
)

// NewRepo new pg repo implimentation
func NewRepo(db *gorm.DB) *repo.Repo {
	return &repo.Repo{}
}
