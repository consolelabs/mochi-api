package userfeedback

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	CreateOne(*model.UserFeedback) (*model.UserFeedback, error)
	GetAll() ([]model.UserFeedback, error)
	GetAllByStatus(status string) ([]model.UserFeedback, error)
	GetAllByCommand(command string) ([]model.UserFeedback, error)
	GetAllByDiscordID(id string) ([]model.UserFeedback, error)
	UpdateStatusByID(id string, status string) (*model.UserFeedback, error)
}
