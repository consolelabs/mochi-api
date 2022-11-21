package userfeedback

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateOne(*model.UserFeedback) (*model.UserFeedback, error)
	GetAll(page int, size int) ([]model.UserFeedback, int64, error)
	GetAllByStatus(status string, page int, size int) ([]model.UserFeedback, int64, error)
	GetAllByCommand(command string, page int, size int) ([]model.UserFeedback, int64, error)
	GetAllByDiscordID(id string, page int, size int) ([]model.UserFeedback, int64, error)
	UpdateStatusByID(id string, status string) (*model.UserFeedback, error)
}
