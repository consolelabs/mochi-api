package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
)

func (e *Entity) GetGmConfig(guildID string) (*model.GuildConfigGmGn, error) {
	config, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
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

func (e *Entity) GetUpvoteTiersConfig() ([]model.UpvoteStreakTier, error) {
	tiers, err := e.repo.UpvoteStreakTier.GetAll()
	if err != nil {
		e.log.Errorf(err, "[e.GetUpvoteTiersConfig] failed to get upvote tiers")
		return nil, err
	}
	return tiers, nil
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

func (e *Entity) GetGuildTokens(guildID string) ([]model.Token, error) {
	if guildID == "" {
		return e.repo.Token.GetDefaultTokens()
	}

	gTokens, err := e.repo.GuildConfigToken.GetByGuildID(guildID)
	if err != nil {
		e.log.Error(err, "[Entity][GetGuildTokens] repo.GuildConfigToken.GetByGuildID failed")
		return nil, err
	}
	// get tokens with guild_default = TRUE
	if len(gTokens) == 0 {
		return e.repo.Token.GetDefaultTokens()
	}

	var tokens []model.Token
	for _, gToken := range gTokens {
		tokens = append(tokens, *gToken.Token)
	}
	return tokens, nil
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

func (e *Entity) checkRoleIDBeenConfig(guildID, roleID string) error {
	err := e.checkRoleIDInDefaultRole(guildID, roleID)
	if err != nil {
		return err
	}
	err = e.checkRoleIDInLevelRole(guildID, roleID)
	if err != nil {
		return err
	}
	err = e.checkRoleIDInNFTRole(guildID, roleID)
	if err != nil {
		return err
	}
	err = e.checkRoleIDInReactionRole(guildID, roleID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Entity) checkRoleIDInLevelRole(guildID, roleID string) error {
	_, err := e.repo.GuildConfigLevelRole.GetByRoleID(guildID, roleID)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		return fmt.Errorf("Role has been used for level role.")
	default:
		e.log.Error(err, "[entity.checkRoleIDInLevelRole] repo.GuildConfigLevelRole.GetByRoleID failed")
		return err
	}
}

func (e *Entity) checkRoleIDInNFTRole(guildID, roleID string) error {
	_, err := e.repo.GuildConfigGroupNFTRole.GetByRoleID(guildID, roleID)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		return fmt.Errorf("Role has been used for NFT role.")
	default:
		e.log.Error(err, "[entity.checkRoleIDInNFTRole] repo.GuildConfigNFTRole.GetByRoleID failed")
		return err
	}
}

func (e *Entity) checkRoleIDInReactionRole(guildID, roleID string) error {
	configs, err := e.ListAllReactionRoles(guildID)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		for _, cfg := range configs.Configs {
			for _, v := range cfg.Roles {
				if v.ID == roleID {
					return fmt.Errorf("Role has been used for reaction role.")
				}
			}
		}
		return nil
	default:
		e.log.Error(err, "[entity.checkRoleIDInReactionRole] e.ListAllReactionRoles failed")
		return err
	}
}

func (e *Entity) checkRoleIDInDefaultRole(guildID, roleID string) error {
	defaultRole, err := e.repo.GuildConfigDefaultRole.GetAllByGuildID(guildID)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		if roleID == defaultRole.RoleID {
			return fmt.Errorf("Role has been used for default role.")
		}
		return nil
	default:
		e.log.Error(err, "[entity.checkRoleIDInDefaultRole] repo.GuildConfigDefaultRole.GetAllByGuildID failed")
		return err
	}
}

func (e *Entity) ConfigLevelRole(req request.ConfigLevelRoleRequest) error {
	err := e.checkRoleIDBeenConfig(req.GuildID, req.RoleID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID}).Error(err, "[entity.ConfigLevelRole] check roleID config failed")
		return err
	}
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

func (e *Entity) AddGuildMemberRole(guildID, userID, roleID string) error {
	return e.discord.GuildMemberRoleAdd(guildID, userID, roleID)
}

func (e *Entity) RemoveGuildMemberRole(guildID, userID, roleID string) error {
	return e.discord.GuildMemberRoleRemove(guildID, userID, roleID)
}

func (e *Entity) ListGuildNFTRoleConfigs(guildID string) ([]model.GuildConfigGroupNFTRole, error) {
	return e.repo.GuildConfigGroupNFTRole.ListByGuildID(guildID)
}

func (e *Entity) ListMemberNFTRolesToAdd(listGroupConfigNFTRoles []model.GuildConfigGroupNFTRole, guildID string) (map[[2]string]bool, error) {
	mrs, err := e.repo.UserNFTBalance.GetUserNFTBalancesByUserInGuild(guildID)
	if err != nil {
		return nil, err
	}
	rolesToAdd := make(map[[2]string]bool)

	for _, mr := range mrs {
		rolesToAdd[[2]string{mr.UserDiscordID, mr.RoleID}] = true
	}
	return rolesToAdd, nil
}

func (e *Entity) NewGuildGroupNFTRoleConfig(req request.ConfigGroupNFTRoleRequest) (*response.ConfigGroupNFTRoleResponse, error) {
	err := e.checkRoleIDBeenConfig(req.GuildID, req.RoleID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "roleID": req.RoleID}).Error(err, "[entity.NewGuildNFTRoleConfig] check roleID config failed")
		return nil, err
	}

	// create record guild_config_group_nft_role
	groupConfig, err := e.repo.GuildConfigGroupNFTRole.Create(model.GuildConfigGroupNFTRole{
		GuildID:        req.GuildID,
		RoleID:         req.RoleID,
		GroupName:      req.GroupName,
		NumberOfTokens: req.NumberOfTokens,
	})
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID":        req.GuildID,
			"roleID":         req.RoleID,
			"collectionID":   req.CollectionAddress,
			"numberOfTokens": req.NumberOfTokens,
		}).Error(err, "[entity.NewGuildNFTRoleConfig] cannot create config group nft role")
		return nil, err
	}

	collectionConfigs := make([]response.NFTCollectionConfig, 0)
	for _, collection := range req.CollectionAddress {
		nftCollection, err := e.repo.NFTCollection.GetByAddress(collection)
		if err != nil {
			e.log.Fields(logger.Fields{"collectionAddress": collection}).Error(err, "[entity.NewGuildNFTRoleConfig] cannot get collection")
			return nil, err
		}

		// create record guild_config_nft_role
		config, err := e.repo.GuildConfigNFTRole.Create(model.GuildConfigNFTRole{
			NFTCollectionID: nftCollection.ID,
			GroupID:         groupConfig.ID,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"nftCollectionId": nftCollection.ID, "groupId": groupConfig.ID}).Error(err, "[entity.NewGuildNFTRoleConfig] cannot create config nft role")
			return nil, err
		}
		collectionConfigs = append(collectionConfigs, response.NFTCollectionConfig{
			ID:           config.ID.UUID.String(),
			CollectionID: nftCollection.ID.UUID.String(),
			Address:      nftCollection.Address,
			Name:         nftCollection.Name,
			Symbol:       nftCollection.Symbol,
			ChainID:      nftCollection.ChainID,
			ERCFormat:    nftCollection.ERCFormat,
			IsVerified:   nftCollection.IsVerified,
			CreatedAt:    nftCollection.CreatedAt,
			Image:        nftCollection.Image,
			Author:       nftCollection.Author,
		})
	}
	return &response.ConfigGroupNFTRoleResponse{
		GuildID:              req.GuildID,
		RoleID:               req.RoleID,
		GroupName:            req.GroupName,
		NFTCollectionConfigs: collectionConfigs,
		NumberOfTokens:       req.NumberOfTokens,
	}, nil

}

func (e *Entity) RemoveGuildGroupNFTRoleConfig(id string) error {
	err := e.repo.GuildConfigNFTRole.DeleteByGroupId(id)
	if err != nil {
		e.log.Fields(logger.Fields{
			"groupConfigNFTRole": id,
		}).Error(err, "[entity.RemoveGuildNFTRoleConfig] cannot delete config nft role")
		return err
	}
	err = e.repo.GuildConfigGroupNFTRole.Delete(id)
	if err != nil {
		e.log.Fields(logger.Fields{
			"groupConfigNFTRole": id,
		}).Error(err, "[entity.RemoveGuildNFTRoleConfig] cannot delete config group nft role")
		return err
	}
	return nil
}

func (e *Entity) RemoveGuildNFTRoleConfig(ids []string) error {
	err := e.repo.GuildConfigNFTRole.DeleteByIds(ids)
	if err != nil {
		e.log.Fields(logger.Fields{
			"ids": ids,
		}).Error(err, "[entity.RemoveGuildNFTRoleConfig] cannot delete config group nft role")
		return err
	}
	return nil
}

func (e *Entity) ListGuildGroupNFTRoles(guildID string) ([]response.ListGuildNFTRoleConfigsResponse, error) {
	groupRoles, err := e.repo.GuildConfigGroupNFTRole.ListByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[entity.ListGuildGroupNFTRoles] cannot get list guild config nft role")
		return nil, fmt.Errorf("failed to list guild nft roles: %v", err.Error())
	}

	dr, err := e.discord.GuildRoles(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[entity.ListGuildGroupNFTRoles] cannot get guild roles from discord")
		return nil, fmt.Errorf("failed to list discord guild roles: %v", err.Error())
	}

	res := make([]response.ListGuildNFTRoleConfigsResponse, 0)
	for _, groupRole := range groupRoles {
		groupRoleResp := response.ListGuildNFTRoleConfigsResponse{
			Id:             groupRole.ID.UUID.String(),
			GuildId:        groupRole.GuildID,
			GroupName:      groupRole.GroupName,
			RoleId:         groupRole.RoleID,
			NumberOfTokens: groupRole.NumberOfTokens,
		}

		// get role name + color from discord to enrich response
		for _, r := range dr {
			if groupRole.RoleID == r.ID {
				groupRoleResp.RoleName = r.Name
				groupRoleResp.Color = r.Color
			}
		}

		// get data collection from db to enrich response
		configs := make([]response.NFTCollectionConfig, 0)
		for _, role := range groupRole.GuildConfigNFTRole {
			collection, err := e.repo.NFTCollection.GetByID(role.NFTCollectionID.UUID.String())
			if err != nil {
				e.log.Fields(logger.Fields{
					"guildID":    guildID,
					"collection": role.NFTCollectionID.UUID.String(),
				}).Error(err, "[entity.ListGuildGroupNFTRoles] cannot get collection for config nft role")
				return nil, err
			}
			configs = append(configs, response.NFTCollectionConfig{
				ID:           role.ID.UUID.String(),
				CollectionID: collection.ID.UUID.String(),
				Address:      collection.Address,
				Name:         collection.Name,
				Symbol:       collection.Symbol,
				ChainID:      collection.ChainID,
				ERCFormat:    collection.ERCFormat,
				IsVerified:   collection.IsVerified,
				CreatedAt:    collection.CreatedAt,
				Image:        collection.Image,
				Author:       collection.Author,
			})
		}
		groupRoleResp.NFTCollectionConfigs = configs
		res = append(res, groupRoleResp)
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
		UserID:          hashtag.UserID,
		GuildID:         hashtag.GuildID,
		ChannelID:       hashtag.ChannelID,
		Hashtag:         strings.Split(hashtag.Hashtag, ","),
		TwitterUsername: strings.Split(hashtag.TwitterUsername, ","),
		RuleID:          hashtag.RuleID,
		FromTwitter:     strings.Split(hashtag.FromTwitter, ","),
		CreatedAt:       hashtag.CreatedAt,
		UpdatedAt:       hashtag.UpdatedAt,
	}, nil
}

func (e *Entity) GetAllTwitterHashtagConfig() ([]response.TwitterHashtag, error) {
	data, err := e.repo.GuildConfigTwitterHashtag.GetAll()
	hashtags := []response.TwitterHashtag{}
	if err != nil {
		e.log.Errorf(err, "[e.GetTwitterHashtagConfig] failed to get twitter hashtag configs")
		return nil, fmt.Errorf("failed to get twitter hashtags: %v", err.Error())
	}
	for _, tag := range data {
		hashtags = append(hashtags, response.TwitterHashtag{
			UserID:          tag.UserID,
			GuildID:         tag.GuildID,
			ChannelID:       tag.ChannelID,
			Hashtag:         strings.Split(tag.Hashtag, ","),
			TwitterUsername: strings.Split(tag.TwitterUsername, ","),
			RuleID:          tag.RuleID,
			FromTwitter:     strings.Split(tag.FromTwitter, ","),
			CreatedAt:       tag.CreatedAt,
			UpdatedAt:       tag.UpdatedAt,
		})
	}
	return hashtags, nil
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
	usernames := ""
	fromTwitter := ""
	for _, tag := range req.Hashtag {
		hashtags += tag + ","
	}
	for _, usr := range req.TwitterUsername {
		usernames += usr + ","
	}
	for _, from := range req.FromTwitter {
		fromTwitter += from + ","
	}
	err := e.repo.GuildConfigTwitterHashtag.UpsertOne(&model.GuildConfigTwitterHashtag{
		UserID:          req.UserID,
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		RuleID:          req.RuleID,
		Hashtag:         strings.TrimSuffix(hashtags, ","), //save as '#abc,#bca,#abe'
		TwitterUsername: strings.TrimSuffix(usernames, ","),
		FromTwitter:     strings.TrimSuffix(fromTwitter, ","),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		e.log.Errorf(err, "[e.CreateTwitterHashtagConfig] failed to upsert twitter hashtag configs")
		return fmt.Errorf("failed to create twitter hashtag: %v", err.Error())
	}
	return nil
}

func (e *Entity) GetDefaultCollectionSymbol(guildID string) ([]model.GuildConfigDefaultCollection, error) {
	data, err := e.repo.GuildConfigDefaultCollection.GetByGuildID(guildID)
	if err != nil {
		e.log.Errorf(err, "[e.GetDefaultCollectionSymbol] failed to get default collection: %s", err)
		return nil, err
	}
	return data, nil
}

func (e *Entity) CreateDefaultCollectionSymbol(req request.ConfigDefaultCollection) error {
	err := e.repo.GuildConfigDefaultCollection.Upsert(&model.GuildConfigDefaultCollection{
		GuildID:   req.GuildID,
		Symbol:    req.Symbol,
		Address:   req.Address,
		ChainID:   util.ConvertChainToChainId(req.ChainID),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		e.log.Errorf(err, "[e.CreateDefaultCollectionSymbol] failed to upsert default ticker: %s", err)
		return err
	}
	return nil
}
