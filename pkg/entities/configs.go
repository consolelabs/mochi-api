package entities

import (
	"fmt"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
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
func (e *Entity) GetSalesTrackerConfig(guildID string) (*model.GuildConfigSalesTracker, error) {
	config, err := e.repo.GuildConfigSalesTracker.GetByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return config, nil
}
func (e *Entity) UpsertSalesTrackerConfig(req request.UpsertSTConfigRequest) error {
	tmp := &model.GuildConfigSalesTracker{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	}

	if err := e.repo.GuildConfigSalesTracker.UpsertOne(tmp); err != nil {
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

func (e *Entity) AddGuildMemberRole(guildID, userID, roleID string) error {
	return e.discord.GuildMemberRoleAdd(guildID, userID, roleID)
}

func (e *Entity) RemoveGuildMemberRole(guildID, userID, roleID string) error {
	return e.discord.GuildMemberRoleRemove(guildID, userID, roleID)
}

func (e *Entity) ListGuildNFTRoleConfigs(guildID string) ([]model.GuildConfigNFTRole, error) {
	return e.repo.GuildConfigNFTRole.ListByGuildID(guildID)
}

func (e *Entity) ListMemberNFTRolesToAdd(guildID string) (map[[2]string]bool, error) {
	mrs, err := e.repo.GuildConfigNFTRole.GetMemberCurrentRoles(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get member current roles: %v", err.Error())
	}

	rolesToAdd := make(map[[2]string]bool)

	for _, mr := range mrs {
		rolesToAdd[[2]string{mr.UserID, mr.RoleID}] = true
	}

	return rolesToAdd, nil
}

func (e *Entity) NewGuildNFTRoleConfig(req request.ConfigNFTRoleRequest) (*model.GuildConfigNFTRole, error) {

	nftcollection, err := e.repo.NFTCollection.GetByID(req.NFTCollectionID.UUID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get nft collection: %v", err.Error())
	}

	if nftcollection.ERCFormat == "1155" && req.TokenID == "" {
		return nil, fmt.Errorf("token id is required for erc1155 nft collections")
	}

	err = e.repo.GuildConfigNFTRole.UpsertOne(&req.GuildConfigNFTRole)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert guild config nft role: %v", err.Error())
	}

	return &req.GuildConfigNFTRole, nil
}

func (e *Entity) EditGuildNFTRoleConfig(req request.ConfigNFTRoleRequest) (*model.GuildConfigNFTRole, error) {

	nftcollection, err := e.repo.NFTCollection.GetByID(req.NFTCollectionID.UUID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get nft collection: %v", err.Error())
	}

	if nftcollection.ERCFormat == "1155" && req.TokenID == "" {
		return nil, fmt.Errorf("token id is required for erc1155 nft collections")
	}

	err = e.repo.GuildConfigNFTRole.Update(&req.GuildConfigNFTRole)
	if err != nil {
		return nil, fmt.Errorf("failed to update guild config nft role: %v", err.Error())
	}

	return &req.GuildConfigNFTRole, nil
}

func (e *Entity) RemoveGuildNFTRoleConfig(id string) error {
	err := e.repo.GuildConfigNFTRole.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to remove guild nft role config")
	}
	return nil
}

func (e *Entity) ListGuildNFTRoles(guildID string) ([]response.GuildNFTRolesResponse, error) {
	roles, err := e.repo.GuildConfigNFTRole.ListByGuildID(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to list guild nft roles: %v", err.Error())
	}

	dr, err := e.discord.GuildRoles(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to list discord guild roles: %v", err.Error())
	}

	nftCollections, err := e.repo.NFTCollection.ListByGuildID(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to list nft collections: %v", err.Error())
	}

	res := make([]response.GuildNFTRolesResponse, len(roles))

	for i, role := range roles {
		roleResp := response.GuildNFTRolesResponse{
			GuildConfigNFTRole: role,
		}

		for _, r := range dr {
			if role.RoleID == r.ID {
				roleResp.RoleName = r.Name
				roleResp.Color = r.Color
			}
		}

		for _, nft := range nftCollections {
			if nft.ID == role.NFTCollectionID {
				roleResp.NFTCollection = nft
			}
		}
		res[i] = roleResp
	}

	return res, nil
}

func (e *Entity) ConfigRepostReaction(req request.ConfigRepostRequest) error {
	return e.repo.GuildConfigRepostReaction.UpsertOne(model.GuildConfigRepostReaction{
		GuildID:         req.GuildID,
		Emoji:           req.Emoji,
		Quantity:        req.Quantity,
		RepostChannelID: req.RepostChannelID,
	})
}

func (e *Entity) GetGuildRepostReactionConfigs(guildID string) ([]model.GuildConfigRepostReaction, error) {
	return e.repo.GuildConfigRepostReaction.GetByGuildID(guildID)
}

func (e *Entity) GetGuildRepostReactionConfigByReaction(guildID string, reaction string) (model.GuildConfigRepostReaction, error) {
	return e.repo.GuildConfigRepostReaction.GetByReaction(guildID, reaction)
}

func (e *Entity) RemoveGuildRepostReactionConfig(guildID string, emoji string) error {
	return e.repo.GuildConfigRepostReaction.DeleteOne(guildID, emoji)
}

func (e *Entity) ListActivityConfigsByName(activityName string) ([]model.GuildConfigActivity, error) {
	activities, err := e.repo.GuildConfigActivity.ListByActivity(activityName)
	if err != nil {
		return nil, fmt.Errorf("failed to list activity configs: %v", err.Error())
	}
	return activities, nil
}

func (e *Entity) ToggleActivityConfig(guildID, activityName string) (*model.GuildConfigActivity, error) {
	activity, err := e.repo.Activity.GetByName(activityName)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity: %v", err.Error())
	}

	config := model.GuildConfigActivity{
		GuildID:    guildID,
		ActivityID: activity.ID,
		Active:     true,
	}

	err = e.repo.GuildConfigActivity.UpsertToggleActive(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert guild config activity: %v", err.Error())
	}

	return &config, nil
}
