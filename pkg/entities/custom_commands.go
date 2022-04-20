package entities

import (
	"github.com/defipod/mochi/pkg/model"
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
)

func (e *Entity) CreateCustomCommand(customCommand model.GuildCustomCommand) error {

	_, err := e.repo.DiscordGuilds.GetByID(customCommand.GuildID)
	if err != nil {
		return err
	}

	return e.repo.GuildCustomCommand.UpsertOne(customCommand)
}

func (e *Entity) UpdateCustomCommand(ID, guildID string, customCommand model.GuildCustomCommand) error {

	_, err := e.repo.GuildCustomCommand.GetByIDAndGuildID(ID, guildID)
	if err != nil {
		return err
	}

	return e.repo.GuildCustomCommand.Update(ID, guildID, customCommand)
}

func (e *Entity) ListCustomCommands(guildID string, enabled *bool) ([]model.GuildCustomCommand, error) {
	return e.repo.GuildCustomCommand.GetAll(guildcustomcommand.GetAllQuery{GuildID: guildID, Enabled: enabled})
}

func (e *Entity) GetCustomCommand(ID, guildID string) (*model.GuildCustomCommand, error) {
	return e.repo.GuildCustomCommand.GetByIDAndGuildID(ID, guildID)
}

func (e *Entity) DeleteCustomCommand(ID, guildID string) error {

	_, err := e.repo.GuildCustomCommand.GetByIDAndGuildID(ID, guildID)
	if err != nil {
		return err
	}

	return e.repo.GuildCustomCommand.Delete(model.GuildCustomCommand{
		ID:      ID,
		GuildID: guildID,
	})
}
