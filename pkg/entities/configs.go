package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetGmConfig(guildID string) (*model.GuildConfigGmGn, error) {
	config, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (e *Entity) UpsertGmConfig(req request.UpsertGmConfigRequest) error {
	if err := e.repo.GuildConfigGmGn.UpsertOne(&model.GuildConfigGmGn{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	}); err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetGuildTokens(guildID string) ([]model.GuildConfigToken, error) {
	guildTokens, err := e.repo.GuildConfigToken.GetByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return guildTokens, nil
}

func (e *Entity) UpsertGuildTokenConfig(req request.UpsertGuildTokenConfigRequest) error {
	token, err := e.repo.Token.GetBySymbol(req.Symbol, true)
	if err != nil {
		return err
	}

	if err := e.repo.GuildConfigToken.UpsertMany([]model.GuildConfigToken{{
		GuildID: req.GuildID,
		TokenID: token.ID,
		Active:  req.Active,
	}}); err != nil {
		return err
	}

	return nil
}

func (e *Entity) ConfigLevelRole(req request.ConfigLevelRoleRequest) error {
	return e.repo.GuildConfigLevelRole.UpsertOne(model.GuildConfigLevelRole{
		GuildID: req.GuildID,
		RoleID:  req.RoleID,
		Level:   req.Level,
	})
}

func (e *Entity) GetGuildLevelRoleConfigs(guildID string) ([]model.GuildConfigLevelRole, error) {
	return e.repo.GuildConfigLevelRole.GetByGuildID(guildID)
}

func (e *Entity) RemoveGuildLevelRoleConfig(guildID string, level int) error {
	return e.repo.GuildConfigLevelRole.DeleteOne(guildID, level)
}

func (e *Entity) GetUserRoleByLevel(guildID string, level int) (string, error) {
	config, err := e.repo.GuildConfigLevelRole.GetHighest(guildID, level)
	if err != nil {
		return "", err
	}

	return config.RoleID, nil
}

func (e *Entity) RemoveGuildMemberRoles(guildID string, rolesToRemove map[string]string) error {
	for userID, roleID := range rolesToRemove {
		if err := e.discord.GuildMemberRoleRemove(guildID, userID, roleID); err != nil {
			return err
		}
	}

	return nil
}

func (e *Entity) AddGuildMemberRoles(guildID string, rolesToAdd map[string]string) error {
	for userID, roleID := range rolesToAdd {
		if err := e.discord.GuildMemberRoleAdd(guildID, userID, roleID); err != nil {
			return err
		}
	}

	return nil
}
