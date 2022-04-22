package guildcustomcommand

import "github.com/defipod/mochi/pkg/model"

type GetAllQuery struct {
	GuildID string
	Enabled *bool
}

type Store interface {
	GetAll(q GetAllQuery) ([]model.GuildCustomCommand, error)
	GetByIDAndGuildID(ID, guildID string) (*model.GuildCustomCommand, error)
	UpsertOne(command model.GuildCustomCommand) error
	Update(ID, guilID string, command model.GuildCustomCommand) error
	Delete(command model.GuildCustomCommand) error
}
