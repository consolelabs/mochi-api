package entities

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) ListAllReactionRoles(guildID string) (*response.ListRoleReactionResponse, error) {
	configs, err := e.repo.GuildConfigReactionRole.ListAllByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	var roleConfigs = make([]response.RoleReactionByMessage, 0)
	for _, c := range configs {
		var roles []response.Role
		err = json.Unmarshal([]byte(c.ReactionRoles), &roles)
		if err != nil {
			return nil, err
		}
		roleConfigs = append(roleConfigs, response.RoleReactionByMessage{
			MessageID: c.MessageID,
			ChannelID: c.ChannelID,
			Roles:     roles,
		})
	}

	var res = response.ListRoleReactionResponse{
		GuildID: guildID,
		Configs: roleConfigs,
		Success: true,
	}
	return &res, nil
}

func (e *Entity) GetReactionRoleByMessageID(guildID, messageID, reaction string) (*response.RoleReactionResponse, error) {
	config, err := e.repo.GuildConfigReactionRole.GetByMessageID(guildID, messageID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var roles []response.Role
	err = json.Unmarshal([]byte(config.ReactionRoles), &roles)
	if err != nil {
		return nil, err
	}

	var filteredRole response.Role
	for _, r := range roles {
		if r.Reaction == reaction {
			filteredRole = r
		}
	}

	var res = response.RoleReactionResponse{
		GuildID:   config.GuildID,
		MessageID: config.MessageID,
		ChannelID: config.ChannelID,
		Role:      filteredRole,
	}
	return &res, nil
}

func (e *Entity) UpdateConfigByMessageID(req request.RoleReactionUpdateRequest) (*response.RoleReactionConfigResponse, error) {
	err := e.checkRoleIDBeenConfig(req.GuildID, req.RoleID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID}).Error(err, "[entity.UpdateConfigByMessageID] check roleID config failed")
		return nil, err
	}
	var roles []response.Role

	config, err := e.repo.GuildConfigReactionRole.GetByMessageID(req.GuildID, req.MessageID)

	// Create new config if not exists
	if errors.Is(err, gorm.ErrRecordNotFound) {
		roles = append(roles, response.Role{
			ID:       req.RoleID,
			Reaction: req.Reaction,
		})
		data, err := json.Marshal(roles)
		if err != nil {
			return nil, err
		}

		err = e.repo.GuildConfigReactionRole.CreateRoleConfig(req, string(data))
		if err != nil {
			return nil, err
		}

		return &response.RoleReactionConfigResponse{
			MessageID: req.MessageID,
			GuildID:   req.GuildID,
			ChannelID: req.ChannelID,
			Roles:     roles,
			Success:   true,
		}, err
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(config.ReactionRoles), &roles)
	if err != nil {
		return nil, err
	}

	if IsRoleAlreadyExist(roles, req.RoleID) {
		return nil, errors.New("role exists")
	}

	if IsEmojiAlreadyExist(roles, req.Reaction) {
		return nil, errors.New("reaction exists")
	}

	roles = append(roles, response.Role{
		ID:       req.RoleID,
		Reaction: req.Reaction,
	})
	data, err := json.Marshal(roles)
	if err != nil {
		return nil, err
	}

	err = e.repo.GuildConfigReactionRole.UpdateRoleConfig(req, string(data))
	if err != nil {
		return nil, err
	}

	var res = response.RoleReactionConfigResponse{
		MessageID: config.MessageID,
		GuildID:   config.GuildID,
		ChannelID: config.ChannelID,
		Roles:     roles,
		Success:   true,
	}
	return &res, nil
}

func (e *Entity) ClearReactionMessageConfig(req request.RoleReactionUpdateRequest) error {
	return e.repo.GuildConfigReactionRole.ClearMessageConfig(req.GuildID, req.MessageID)
}

func (e *Entity) RemoveSpecificRoleReaction(req request.RoleReactionUpdateRequest) error {
	var roles []response.Role

	config, err := e.repo.GuildConfigReactionRole.GetByMessageID(req.GuildID, req.MessageID)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(config.ReactionRoles), &roles)
	if err != nil {
		return err
	}

	var updatedRoles []response.Role
	for _, r := range roles {
		if r.ID != req.RoleID {
			updatedRoles = append(updatedRoles, r)
		}
	}

	if len(updatedRoles) == 0 {
		err = e.repo.GuildConfigReactionRole.ClearMessageConfig(req.GuildID, req.MessageID)
		if err != nil {
			return err
		}
	} else {
		data, err := json.Marshal(updatedRoles)
		if err != nil {
			return err
		}

		err = e.repo.GuildConfigReactionRole.UpdateRoleConfig(req, string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func IsRoleAlreadyExist(roles []response.Role, roleID string) bool {
	existed := false
	for _, r := range roles {
		if r.ID == roleID {
			existed = true
		}
	}
	return existed
}

func IsEmojiAlreadyExist(roles []response.Role, reaction string) bool {
	existed := false
	for _, r := range roles {
		if r.Reaction == reaction {
			existed = true
		}
	}
	return existed
}
