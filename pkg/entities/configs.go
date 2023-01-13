package entities

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	guildconfigtwitterblacklist "github.com/defipod/mochi/pkg/repo/guild_config_twitter_blacklist"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetGmConfig(guildID string) (*model.GuildConfigGmGn, error) {
	config, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetGmConfig] repo.GuildConfigGmGn.GetByGuildID failed")
		return nil, err
	}

	return config, nil
}

func (e *Entity) UpsertGmConfig(req request.UpsertGmConfigRequest) error {
	if err := e.repo.GuildConfigGmGn.CreateOne(&model.GuildConfigGmGn{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		Msg:       req.Msg,
		Emoji:     req.Emoji,
		Sticker:   req.Sticker,
	}); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[Entity][UpsertGmConfig] repo.GuildConfigGmGn.UpsertOne failed")
		return err
	}

	return nil
}

func (e *Entity) GetWelcomeChannelConfig(guildID string) (*model.GuildConfigWelcomeChannel, error) {
	config, err := e.repo.GuildConfigWelcomeChannel.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetWelcomeChannelConfig] repo.GuildConfigWelcomeChannel.GetByGuildID failed")
		return nil, err
	}

	return config, nil
}

func (e *Entity) UpsertWelcomeChannelConfig(req request.UpsertWelcomeConfigRequest) (*model.GuildConfigWelcomeChannel, error) {
	if req.WelcomeMsg == "" {
		previousConfig, err := e.repo.GuildConfigWelcomeChannel.GetByGuildID(req.GuildID)
		if err == gorm.ErrRecordNotFound {
			req.WelcomeMsg = "Greetings $name :wave: Welcome to the guild! Hope you enjoy your stay."
		} else {
			req.WelcomeMsg = previousConfig.WelcomeMessage
		}
	}
	config, err := e.repo.GuildConfigWelcomeChannel.UpsertOne(&model.GuildConfigWelcomeChannel{
		GuildID:        req.GuildID,
		ChannelID:      req.ChannelID,
		WelcomeMessage: req.WelcomeMsg,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[Entity][UpsertWelcomeChannelConfig] repo.GuildConfigWelcomeChannel.UpsertOne failed")
		return nil, err
	}

	return config, nil
}

func (e *Entity) DeleteWelcomeChannelConfig(req request.DeleteWelcomeConfigRequest) error {
	if err := e.repo.GuildConfigWelcomeChannel.DeleteOne(&model.GuildConfigWelcomeChannel{
		GuildID: req.GuildID,
	}); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[Entity][DeleteWelcomeChannelConfig] repo.GuildConfigWelcomeChannel.DeleteOne failed")
		return err
	}

	return nil
}

func (e *Entity) GetVoteChannelConfig(guildID string) (*model.GuildConfigVoteChannel, error) {
	config, err := e.repo.GuildConfigVoteChannel.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetVoteChannelConfig] repo.GuildConfigVoteChannel.GetByGuildID failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) UpsertVoteChannelConfig(req request.UpsertVoteChannelConfigRequest) (*model.GuildConfigVoteChannel, error) {
	config, err := e.repo.GuildConfigVoteChannel.UpsertOne(&model.GuildConfigVoteChannel{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[Entity][UpsertVoteChannelConfig] repo.GuildConfigVoteChannel.UpsertOne failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) DeleteVoteChannelConfig(req request.DeleteVoteChannelConfigRequest) error {
	if err := e.repo.GuildConfigVoteChannel.DeleteOne(&model.GuildConfigVoteChannel{
		GuildID: req.GuildID,
	}); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[Entity][DeleteVoteChannelConfig] repo.GuildConfigVoteChannel.DeleteOne failed")
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
		return fmt.Errorf("role has been used for level role")
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
		return fmt.Errorf("role has been used for NFT role")
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
					return fmt.Errorf("role has been used for reaction role")
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
			return fmt.Errorf("role has been used for default role")
		}
		return nil
	default:
		e.log.Error(err, "[entity.checkRoleIDInDefaultRole] repo.GuildConfigDefaultRole.GetAllByGuildID failed")
		return err
	}
}

func (e *Entity) GetUserRoleByLevel(guildID string, level int) (string, error) {
	config, err := e.repo.GuildConfigLevelRole.GetHighest(guildID, level)
	if err != nil {
		return "", err
	}

	return config.RoleID, nil
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
				ExplorerUrl:  util.GetCollectionExplorerUrl(collection.Address, collection.ChainID),
				ChainName:    util.ConvertChainIDToChain(collection.ChainID),
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
		ReactionType:    consts.ReactionTypeMessage,
	})
}

func (e *Entity) CreateConfigRepostReactionConversation(req request.ConfigRepostReactionStartStop) error {
	return e.repo.GuildConfigRepostReaction.UpsertOne(model.GuildConfigRepostReaction{
		GuildID:         req.GuildID,
		EmojiStart:      req.EmojiStart,
		EmojiStop:       req.EmojiStop,
		RepostChannelID: req.RepostChannelID,
		ReactionType:    consts.ReactionTypeConversation,
	})
}

func (e *Entity) GetGuildRepostReactionConfigs(guildID string, reactionType string) ([]model.GuildConfigRepostReaction, error) {
	return e.repo.GuildConfigRepostReaction.GetByGuildIDAndReactionType(guildID, reactionType)
}

func (e *Entity) CreateRepostMessageReactionEvent(req request.MessageReactionRequest) (*model.MessageRepostHistory, error) {
	blacklistChannels, err := e.repo.GuildBlacklistChannelRepostConfigs.GetByGuildID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[e.CreateRepostReactionEvent] failed to get blacklist channel repost config")
		return nil, err
	}
	for _, v := range blacklistChannels {
		if req.ChannelID == v.ChannelID {
			return nil, nil
		}
	}
	conf, err := e.repo.GuildConfigRepostReaction.GetByReaction(req.GuildID, req.Reaction)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// reaction is not match with config sb, so we check if it match with config start, stop
			confStartStop, err := e.repo.GuildConfigRepostReaction.GetByReactionStartOrStop(req.GuildID, req.Reaction)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, nil
				}
				return nil, err
			}

			err = e.CreateRepostReactionEventWithStartStop(req, confStartStop)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}

		e.log.Fields(logger.Fields{"guildID": req.GuildID, "reaction": req.Reaction}).
			Error(err, "[e.CreateRepostReactionEvent] failed to get guild config repost reaction")
		return nil, err
	}

	// && req.Reaction != conf.EmojiStart && req.Reaction != conf.EmojiStop
	if req.Reaction != conf.Emoji {
		return nil, nil
	}
	if req.ReactionCount < conf.Quantity {
		return nil, nil
	}

	var repostMsgRes model.MessageRepostHistory
	// server has start, stop config
	if conf.EmojiStart != "" && conf.EmojiStop != "" {
		// when has reaction start -> is_start = true and is_stop = false -> if different with this -> not allow to repost
		msgRepostHistory, err := e.repo.MessageRepostHistory.GetByMessageID(req.GuildID, req.MessageID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				e.log.Infof("[e.CreateRepostReactionEvent] message repost history not found")
				return nil, nil
			}
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create start repost message history")
			return nil, err
		}
		if msgRepostHistory.IsStart && !msgRepostHistory.IsStop {
			req.RepostChannelID = conf.RepostChannelID
			repostMsg, err := e.CreateRepostMessageHistory(req)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create repost message history")
				return nil, err
			}

			repostMsg.RepostChannelID = conf.RepostChannelID
			repostMsgRes = *repostMsg
		}
	} else {
		req.RepostChannelID = conf.RepostChannelID
		repostMsg, err := e.CreateRepostMessageHistory(req)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create repost message history")
			return nil, err
		}

		repostMsg.RepostChannelID = conf.RepostChannelID
		repostMsgRes = *repostMsg
	}

	return &repostMsgRes, nil
}

func (e *Entity) CreateRepostConversationReactionEvent(req request.MessageReactionRequest) (*model.ConversationRepostHistories, error) {
	blacklistChannels, err := e.repo.GuildBlacklistChannelRepostConfigs.GetByGuildID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[e.CreateRepostConversationReactionEvent] failed to get blacklist channel repost config")
		return nil, err
	}
	for _, v := range blacklistChannels {
		if req.ChannelID == v.ChannelID {
			return nil, nil
		}
	}
	configConversation, err := e.repo.GuildConfigRepostReaction.GetByReactionConversationStartOrStop(req.GuildID, req.Reaction)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "reaction": req.Reaction}).Error(err, "[e.CreateRepostReactionEvent] failed to get guild config repost reaction")
		return nil, err
	}

	if req.Reaction == configConversation.EmojiStart {
		// start conversation repost by create record with origin_start_message_id
		err := e.repo.ConversationRepostHistories.Upsert(model.ConversationRepostHistories{
			GuildID:              req.GuildID,
			OriginChannelID:      req.ChannelID,
			OriginStartMessageID: req.MessageID,
			RepostChannelID:      configConversation.RepostChannelID,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create conversation repost history")
			return nil, err
		}
		return nil, nil
	}

	if req.Reaction == configConversation.EmojiStop {
		// stop conversation repost by update record with origin_stop_message_id
		// first get with channel and guild id to see if theres any started conversation
		// - yes: get and update with stop message id
		// - no: return nil
		conversationRepostHistory, err := e.repo.ConversationRepostHistories.GetByGuildAndChannel(req.GuildID, req.ChannelID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to get conversation repost history")
			return nil, err
		}

		err = e.repo.ConversationRepostHistories.Update(&model.ConversationRepostHistories{
			ID:                   conversationRepostHistory.ID,
			GuildID:              req.GuildID,
			OriginChannelID:      req.ChannelID,
			OriginStartMessageID: conversationRepostHistory.OriginStartMessageID,
			OriginStopMessageID:  req.MessageID,
			RepostChannelID:      configConversation.RepostChannelID,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to stop conversation repost")
			return nil, err
		}
		conversationRepostHistory.OriginStopMessageID = req.MessageID

		return conversationRepostHistory, nil
	}

	return nil, nil
}

func (e *Entity) CreateRepostReactionEventWithStartStop(req request.MessageReactionRequest, conf model.GuildConfigRepostReaction) error {
	// check emoji == start -> start to allow user react
	if req.Reaction == conf.EmojiStart && req.ReactionCount == 1 {
		reqStartMessageRepost := &model.MessageRepostHistory{
			GuildID:         req.GuildID,
			OriginChannelID: req.ChannelID,
			OriginMessageID: req.MessageID,
			RepostChannelID: conf.RepostChannelID,
			ReactionCount:   0,
			IsStart:         true,
			IsStop:          false,
		}

		err := e.repo.MessageRepostHistory.Upsert(*reqStartMessageRepost)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create start repost message history")
			return err
		}
		return nil
	}

	// check emoji == stop -> stop to allow user react
	if req.Reaction == conf.EmojiStop && req.ReactionCount >= 1 {
		reqStopMessageRepost := &model.MessageRepostHistory{
			GuildID:         req.GuildID,
			OriginChannelID: req.ChannelID,
			OriginMessageID: req.MessageID,
			RepostChannelID: conf.RepostChannelID,
			IsStart:         true,
			IsStop:          true,
		}
		err := e.repo.MessageRepostHistory.Upsert(*reqStopMessageRepost)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[e.CreateRepostReactionEvent] failed to create stop repost message history")
			return err
		}
		return nil
	}
	return nil
}

func (e *Entity) RemoveGuildRepostReactionConfig(guildID, emoji string) error {
	_, err := e.repo.GuildConfigRepostReaction.GetByReaction(guildID, emoji)
	if err == gorm.ErrRecordNotFound {
		return baseerrs.ErrRecordNotFound
	}
	return e.repo.GuildConfigRepostReaction.DeleteConfigMessage(guildID, emoji)
}

func (e *Entity) RemoveConfigRepostReactionConversation(guildID, emojiStart, emojiStop string) error {
	return e.repo.GuildConfigRepostReaction.DeleteConfigConversation(guildID, emojiStart, emojiStop)
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
		return nil, err
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
		return nil, err
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

func (e *Entity) GetGuildPruneExclude(guildID string) (*response.GuildPruneExcludeList, error) {
	configs, err := e.repo.GuildConfigPruneExclude.GetByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.GetGuildPruneExclude] failed to get prune excluded roles")
		return nil, err
	}
	if len(configs) == 0 {
		return nil, nil
	}

	roles := []string{}
	for _, cfg := range configs {
		roles = append(roles, cfg.RoleID)
	}

	return &response.GuildPruneExcludeList{
		GuildID: configs[0].GuildID,
		Roles:   roles,
	}, nil
}

func (e *Entity) UpsertGuildPruneExclude(req request.UpsertGuildPruneExcludeRequest) error {
	err := e.repo.GuildConfigPruneExclude.UpsertOne(&model.GuildConfigWhitelistPrune{
		GuildID: req.GuildID,
		RoleID:  req.RoleID,
	})
	if err != nil {
		e.log.Errorf(err, "[e.UpsertGuildPruneExclude] failed to upsert prune excluded roles: %s", err)
		return err
	}
	return nil
}

func (e *Entity) DeleteGuildPruneExclude(req request.UpsertGuildPruneExcludeRequest) error {
	err := e.repo.GuildConfigPruneExclude.DeleteOne(&model.GuildConfigWhitelistPrune{
		GuildID: req.GuildID,
		RoleID:  req.RoleID,
	})
	if err != nil {
		e.log.Errorf(err, "[e.DeleteGuildPruneExclude] failed to delete prune excluded roles: %s", err)
		return err
	}
	return nil
}

func (e *Entity) EditMessageRepost(req *request.EditMessageRepostRequest) error {
	err := e.repo.MessageRepostHistory.EditMessageRepost(req)
	if err != nil {
		e.log.Errorf(err, "[e.EditMessageRepost] failed to edit message repost: %s", err)
		return err
	}
	return nil
}

func (e *Entity) GetJoinLeaveChannelConfig(guildID string) (*model.GuildConfigJoinLeaveChannel, error) {
	config, err := e.repo.GuildConfigJoinLeaveChannel.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[Entity][GetJoinLeaveChannelConfig] repo.GuildConfigJoinLeaveChannel.GetByGuildID failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) UpsertJoinLeaveChannelConfig(req request.UpsertJoinLeaveChannelConfigRequest) (*model.GuildConfigJoinLeaveChannel, error) {
	config, err := e.repo.GuildConfigJoinLeaveChannel.UpsertOne(&model.GuildConfigJoinLeaveChannel{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[Entity][UpsertJoinLeaveChannelConfig] repo.GuildConfigJoinLeaveChannel.UpsertOne failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) DeleteJoinLeaveChannelConfig(req request.DeleteJoinLeaveChannelConfigRequest) error {
	if err := e.repo.GuildConfigJoinLeaveChannel.DeleteOne(&model.GuildConfigJoinLeaveChannel{
		GuildID: req.GuildID,
	}); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID}).Error(err, "[Entity][DeleteJoinLeaveChannelConfig] repo.GuildConfigJoinLeaveChannel.DeleteOne failed")
		return err
	}

	return nil
}

func (e *Entity) CreateBlacklistChannelRepostConfig(req request.BalcklistChannelRepostConfigRequest) error {
	if err := e.repo.GuildBlacklistChannelRepostConfigs.UpsertOne(model.GuildBlacklistChannelRepostConfig{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	}); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[entities.DeleteJoinLeaveChannelConfig] - failed to create blacklist channel repost config")
		return err
	}
	return nil
}

func (e *Entity) GetGuildBlacklistChannelRepostConfig(guildID string) ([]model.GuildBlacklistChannelRepostConfig, error) {
	configs, err := e.repo.GuildBlacklistChannelRepostConfigs.GetByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entities.DeleteJoinLeaveChannelConfig] - failed to get blacklist channel repost config")
		return nil, err
	}
	return configs, nil
}

func (e *Entity) DeleteBlacklistChannelRepostConfig(req request.BalcklistChannelRepostConfigRequest) error {
	if err := e.repo.GuildBlacklistChannelRepostConfigs.DeleteOne(req.GuildID, req.ChannelID); err != nil {
		e.log.Fields(logger.Fields{"guildID": req.GuildID, "channelID": req.ChannelID}).Error(err, "[entities.DeleteJoinLeaveChannelConfig] - failed to create blacklist channel repost config")
		return err
	}
	return nil
}

func (e *Entity) AddToTwitterBlackList(req request.AddToTwitterBlackListRequest) error {
	return e.repo.GuildConfigTwitterBlacklist.Upsert(&model.GuildConfigTwitterBlacklist{
		GuildID:         req.GuildID,
		TwitterUsername: req.TwitterUsername,
		TwitterID:       req.TwitterID,
		CreatedBy:       req.CreatedBy,
	})
}

func (e *Entity) DeleteFromTwitterBlackList(req request.DeleteFromTwitterBlackListRequest) error {
	return e.repo.GuildConfigTwitterBlacklist.Delete(req.GuildID, req.TwitterID)
}

func (e *Entity) GetTwitterBlackList(guildID string) ([]model.GuildConfigTwitterBlacklist, error) {
	return e.repo.GuildConfigTwitterBlacklist.List(guildconfigtwitterblacklist.ListQuery{GuildID: guildID})
}

func (e *Entity) GetUserTokenAlert(discordID string) (*response.DiscordUserTokenAlertResponse, error) {
	data, err := e.repo.DiscordUserTokenAlert.GetByDiscordID(discordID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"discordID": discordID}).Error(err, "[entities.GetUserTokenAlert] - failed to get user token alerts")
		return nil, err
	}
	return &response.DiscordUserTokenAlertResponse{
		Data: data,
	}, err
}

func (e *Entity) GetUserTokenAlertByID(id string) (*model.DiscordUserTokenAlert, error) {
	data, err := e.repo.DiscordUserTokenAlert.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Fields(logger.Fields{"id": id}).Error(err, "[entities.GetUserTokenAlert] - failed to get user token alerts by id")
		return nil, err
	}
	return data, err
}

func (e *Entity) GetAllUserTokenAlert() (*response.DiscordUserTokenAlertResponse, error) {
	data, err := e.repo.DiscordUserTokenAlert.GetAllActive()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		e.log.Error(err, "[entities.GetAllUserTokenAlert] - failed to get all user token alerts")
		return nil, err
	}
	return &response.DiscordUserTokenAlertResponse{
		Data: data,
	}, err
}

func (e *Entity) UpsertUserTokenAlert(req *request.UpsertDiscordUserAlertRequest) error {
	alert, err := e.repo.DiscordUserTokenAlert.UpsertOne(&model.UpsertDiscordUserTokenAlert{
		ID:        util.GetNullUUID(req.ID),
		TokenID:   req.TokenID,
		DiscordID: req.DiscordID,
		Symbol:    strings.ToLower(req.Symbol),
		PriceSet:  req.PriceSet,
		IsEnable:  req.IsEnable,
		Trend:     req.Trend,
		DeviceID:  req.DeviceID,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		e.log.Error(err, "[entities.UpsertUserTokenAlert] - failed to create user token alert")
		return err
	}

	// alert_<tokenID>_<up/down> : <alertID> - <price>
	err = e.cache.ZSet(fmt.Sprintf("alert_%s_%s", strings.ToLower(req.Symbol), req.Trend), alert.ID.UUID.String(), req.PriceSet)
	if err != nil {
		e.log.Error(err, "[entities.UpsertUserTokenAlert] - failed to cache alert")
		return err
	}
	return nil
}
func (e *Entity) DeleteUserTokenAlert(req *request.DeleteDiscordUserAlertRequest) error {
	alert, err := e.repo.DiscordUserTokenAlert.RemoveOne(&model.DiscordUserTokenAlert{
		ID: util.GetNullUUID(req.ID),
	})
	if err != nil {
		e.log.Error(err, "[entities.DeleteUserTokenAlert] - failed to delete user token alert")
		return err
	}
	err = e.cache.ZRemove(fmt.Sprintf("alert_%s_%s", strings.ToLower(alert.Symbol), alert.Trend), alert.ID.UUID.String())
	if err != nil {
		e.log.Error(err, "[entities.DeleteUserTokenAlert] - failed to remove cache")
		return err
	}
	return nil
}

func (e *Entity) UpsertMonikerConfig(req request.UpsertMonikerConfigRequest) error {
	token, err := e.repo.OffchainTipBotTokens.GetBySymbol(req.Token)
	if err != nil {
		e.log.Fields(logger.Fields{"token": req.Token}).Error(err, "[entities.UpsertMonikerConfig] - failed to get user token")
		return err
	}
	return e.repo.MonikerConfig.UpsertOne(model.MonikerConfig{
		Moniker: req.Moniker,
		Plural:  req.Plural,
		TokenID: token.ID,
		Amount:  req.Amount,
		GuildID: req.GuildID,
	})
}

func (e *Entity) GetMonikerByGuildID(guildID string) ([]response.MonikerConfigData, error) {
	configs, err := e.repo.MonikerConfig.GetByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entities.GetMonikerByGuildID] - failed to get moniker configs")
		return nil, err
	}
	tokenLst := []string{}
	checkMap := make(map[string]bool)
	for _, item := range configs {
		if _, value := checkMap[item.Token.CoinGeckoID]; !value {
			checkMap[item.Token.TokenSymbol] = true
			tokenLst = append(tokenLst, item.Token.CoinGeckoID)
		}
	}
	prices, err := e.svc.CoinGecko.GetCoinPrice(tokenLst, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": tokenLst}).Error(err, "[entities.GetMonikerByGuildID] - failed to get coin price")
		return nil, err
	}
	res := []response.MonikerConfigData{}
	for _, config := range configs {
		var configData response.MonikerConfigData
		configData.Moniker = config
		configData.Value = config.Amount * prices[config.Token.CoinGeckoID]
		res = append(res, configData)
	}
	return res, nil
}

func (e *Entity) DeleteMonikerConfig(req request.DeleteMonikerConfigRequest) error {
	return e.repo.MonikerConfig.DeleteOne(req.GuildID, req.Moniker)
}

func (e *Entity) GetDefaultMoniker() ([]response.MonikerConfigData, error) {
	configs, err := e.repo.MonikerConfig.GetDefaultMoniker()
	if err != nil {
		e.log.Error(err, "[entities.GetDefaultMoniker] - failed to get moniker default configs")
		return nil, err
	}

	tokenLst := []string{}
	checkMap := make(map[string]bool)
	for _, item := range configs {
		if _, value := checkMap[item.Token.CoinGeckoID]; !value {
			checkMap[item.Token.TokenSymbol] = true
			tokenLst = append(tokenLst, item.Token.CoinGeckoID)
		}
	}
	prices, err := e.svc.CoinGecko.GetCoinPrice(tokenLst, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"token": tokenLst}).Error(err, "[entities.GetDefaultMoniker] - failed to get coin price")
		return nil, err
	}

	res := []response.MonikerConfigData{}
	for _, config := range configs {
		var configData response.MonikerConfigData
		configData.Moniker = config
		configData.Value = config.Amount * prices[config.Token.CoinGeckoID]
		res = append(res, configData)
	}
	return res, nil
}

func (e *Entity) GetGuildDefaultCurrency(guildID string) (*response.GuildConfigDefaultCurrencyResponse, error) {
	data, err := e.repo.GuildConfigDefaultCurrency.GetByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entities.GetGuildDefaultCurrency] - failed to get default currency")
		return nil, err
	}
	return &response.GuildConfigDefaultCurrencyResponse{
		GuildID:     data.GuildID,
		TipBotToken: &data.TipBotToken,
		UpdatedAt:   data.UpdatedAt,
		CreatedAt:   data.CreatedAt,
	}, err
}

func (e *Entity) UpsertGuildDefaultCurrency(req request.UpsertGuildDefaultCurrencyRequest) error {
	token, err := e.repo.OffchainTipBotTokens.GetBySymbol(strings.ToUpper(req.Symbol))
	if err != nil {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.UpsertGuildDefaultCurrency] - failed to get tip bot token")
		return fmt.Errorf("token symbol not found")
	}

	err = e.repo.GuildConfigDefaultCurrency.Upsert(&model.UpsertGuildConfigDefaultCurrency{
		GuildID:       req.GuildID,
		TipBotTokenID: token.ID.String(),
		UpdatedAt:     time.Now().UTC(),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.UpsertGuildDefaultCurrency] - failed to upsert default currency")
		return err
	}
	return nil
}

func (e *Entity) DeleteGuildDefaultCurrency(guildID string) error {
	err := e.repo.GuildConfigDefaultCurrency.DeleteByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entities.DeleteGuildDefaultCurrency] - failed to upsert default currency")
		return err
	}
	return nil
}

func (e *Entity) CreateGuildTokenRole(req request.CreateGuildTokenRole) (*model.GuildConfigTokenRole, error) {
	chainIDStr := util.ConvertInputToChainId(req.Chain)
	if chainIDStr == "" {
		e.log.Fields(logger.Fields{
			"chain": req.Chain,
		}).Error(errors.ErrInvalidChain, "[e.CreateGuildTokenRole] - util.ConvertInputToChainId failed")
		return nil, errors.ErrInvalidChain
	}
	chainID, err := strconv.Atoi(chainIDStr)
	if err != nil {
		e.log.Fields(logger.Fields{
			"chainID": chainIDStr,
		}).Error(err, "[e.CreateGuildTokenRole] - strconv.Atoi failed")
		return nil, err
	}
	token, err := e.repo.Token.GetByAddress(req.Address, chainID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"address": req.Address,
			"chainID": chainID,
		}).Error(err, "[e.CreateGuildTokenRole] - repo.Token.GetByAddress failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrTokenNotFound
		}
		return nil, err
	}
	config := &model.GuildConfigTokenRole{
		GuildID:        req.GuildID,
		RoleID:         req.RoleID,
		RequiredAmount: req.Amount,
		TokenID:        token.ID,
	}
	if err := e.repo.GuildConfigTokenRole.Create(config); err != nil {
		e.log.Fields(logger.Fields{
			"config": config,
		}).Error(err, "[e.CreateGuildTokenRole] - repo.GuildConfigTokenRole.Create failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) ListGuildTokenRoles(guildID string) ([]model.GuildConfigTokenRole, error) {
	configs, err := e.repo.GuildConfigTokenRole.ListByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[e.ListGuildTokenRoles] - repo.GuildConfigTokenRole.ListByGuildID failed")
		return nil, err
	}
	return configs, nil
}

func (e *Entity) UpdateGuildTokenRole(id int, req request.UpdateGuildTokenRole) (*model.GuildConfigTokenRole, error) {
	config, err := e.repo.GuildConfigTokenRole.Get(id)
	if err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[e.UpdateGuildTokenRole] - repo.GuildConfigTokenRole.Get failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}
	if req.RoleID != nil {
		config.RoleID = *req.RoleID
	}
	if req.Amount != nil {
		config.RequiredAmount = *req.Amount
	}
	if req.Address != nil && req.Chain != nil {
		chainIDStr := util.ConvertInputToChainId(*req.Chain)
		if chainIDStr == "" {
			e.log.Fields(logger.Fields{
				"chain": req.Chain,
			}).Error(errors.ErrInvalidChain, "[e.UpdateGuildTokenRole] - util.ConvertInputToChainId failed")
			return nil, errors.ErrInvalidChain
		}
		chainID, err := strconv.Atoi(chainIDStr)
		if err != nil {
			e.log.Fields(logger.Fields{
				"chainID": chainIDStr,
			}).Error(err, "[e.UpdateGuildTokenRole] - strconv.Atoi failed")
			return nil, err
		}
		token, err := e.repo.Token.GetByAddress(*req.Address, chainID)
		if err != nil {
			e.log.Fields(logger.Fields{
				"address": req.Address,
				"chainID": chainID,
			}).Error(err, "[e.UpdateGuildTokenRole] - repo.Token.GetByAddress failed")
			if err == gorm.ErrRecordNotFound {
				return nil, errors.ErrTokenNotFound
			}
			return nil, err
		}
		config.TokenID = token.ID
	}
	if err := e.repo.GuildConfigTokenRole.Update(config); err != nil {
		e.log.Fields(logger.Fields{
			"config": config,
		}).Error(err, "[e.UpdateGuildTokenRole] - repo.GuildConfigTokenRole.Update failed")
		return nil, err
	}
	return config, nil
}

func (e *Entity) RemoveGuildTokenRole(id int) error {
	if _, err := e.repo.GuildConfigTokenRole.Get(id); err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[e.RemoveGuildTokenRole] - repo.GuildConfigTokenRole.Get failed")
		if err == gorm.ErrRecordNotFound {
			return errors.ErrRecordNotFound
		}
		return err
	}
	if err := e.repo.GuildConfigTokenRole.Delete(id); err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[e.RemoveGuildTokenRole] - repo.GuildConfigTokenRole.Delete failed")
		return err
	}
	return nil
}
