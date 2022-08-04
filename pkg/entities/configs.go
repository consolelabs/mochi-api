package entities

import (
	"fmt"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
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
func (e *Entity) UpsertSalesTrackerConfig(req request.UpsertSalesTrackerConfigRequest) error {
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
		gMemberRoleLog := e.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		})
		if err := e.discord.GuildMemberRoleRemove(guildID, userID, roleID); err != nil {
			gMemberRoleLog.Error(err, "[Entity][RemoveGuildMemberRoles] discord.GuildMemberRoleRemove failed")
			return err
		}
		gMemberRoleLog.Info("[Entity][RemoveGuildMemberRoles] discord.GuildMemberRoleRemove executed successfully")
	}
	return nil
}

func (e *Entity) AddGuildMemberRoles(guildID string, rolesToAdd map[string]string) error {
	for userID, roleID := range rolesToAdd {
		gMemberRoleLog := e.log.Fields(logger.Fields{
			"guildId": guildID,
			"userId":  userID,
			"roleId":  roleID,
		})
		if err := e.discord.GuildMemberRoleAdd(guildID, userID, roleID); err != nil {
			gMemberRoleLog.Error(err, "[Entity][AddGuildMemberRoles] discord.GuildMemberRoleAdd failed")
			return err
		}
		gMemberRoleLog.Info("[Entity][AddGuildMemberRoles] discord.GuildMemberRoleAdd executed successfully")
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

func (e *Entity) CreateRepostReactionEvent(req request.CreateMessageRepostHistRequest) (string, error) {
	conf, err := e.repo.GuildConfigRepostReaction.GetByReaction(req.GuildID, req.Reaction)
	if err != nil {
		return "", err
	}
	if req.ReactionCount < conf.Quantity {
		return "", nil
	}
	if isRepostable := e.IsRepostableMessage(req); !isRepostable {
		return "", fmt.Errorf("message cannot be reposted")
	}
	err = e.CreateRepostMessageHist(req, conf.RepostChannelID)
	if err != nil {
		return "", err
	}
	return conf.RepostChannelID, nil
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

func (e *Entity) GetAllTwitterConfig() ([]model.GuildConfigTwitterFeed, error) {
	configs, err := e.repo.GuildConfigTwitterFeed.GetAll()
	if err != nil {
		e.log.Errorf(err, "[e.GetAllTwitterConfig] failed to get all twitter configs")
		return nil, fmt.Errorf("failed to get twitter configs: %v", err.Error())
	}
	return configs, nil
}
func (e *Entity) CreateTwitterConfig(req *model.GuildConfigTwitterFeed) error {
	err := e.repo.GuildConfigTwitterFeed.UpsertOne(req)
	if err != nil {
		e.log.Errorf(err, "[e.CreateTwitterConfig] failed to upsert twitter configs")
		return fmt.Errorf("failed to upsert twitter configs: %v", err.Error())
	}
	return nil
}

func (e *Entity) GetTwitterHashtagConfig(guildId string) (*response.TwitterHashtag, error) {
	hashtag, err := e.repo.GuildConfigTwitterHashtag.GetByGuildID(guildId)
	if err != nil {
		e.log.Errorf(err, "[e.GetTwitterHashtagConfig] failed to get twitter hashtag configs")
		return nil, fmt.Errorf("failed to get twitter hashtags: %v", err.Error())
	}
	return &response.TwitterHashtag{
		UserID:    hashtag.UserID,
		GuildID:   hashtag.GuildID,
		ChannelID: hashtag.ChannelID,
		Hashtag:   strings.Split(hashtag.Hashtag, ","),
		CreatedAt: hashtag.CreatedAt,
	}, nil
}

func (e *Entity) DeleteTwitterHashtagConfig(guildId string) error {
	err := e.repo.GuildConfigTwitterHashtag.DeleteByGuildID(guildId)
	if err != nil {
		e.log.Errorf(err, "[e.DeleteTwitterHashtagConfig] failed to delete twitter hashtag configs")
		return fmt.Errorf("failed to delete twitter hashtags: %v", err.Error())
	}
	return nil
}

func (e *Entity) CreateTwitterHashtagConfig(req *request.TwitterHashtag) error {
	hashtags := ""
	for _, tag := range req.Hashtag {
		hashtags += tag + ","
	}
	err := e.repo.GuildConfigTwitterHashtag.UpsertOne(&model.GuildConfigTwitterHashtag{
		UserID:    req.UserID,
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		Hashtag:   strings.TrimSuffix(hashtags, ","), //save as '#abc,#bca,#abe'
	})
	if err != nil {
		e.log.Errorf(err, "[e.CreateTwitterHashtagConfig] failed to upsert twitter hashtag configs")
		return fmt.Errorf("failed to create twitter hashtag: %v", err.Error())
	}
	return nil
}
