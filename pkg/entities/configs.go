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
	query "github.com/defipod/mochi/pkg/repo/guild_config_log_channel"
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
	config, err := e.repo.GuildConfigLevelRole.GetCurrentLevelRole(guildID, level)
	if err != nil {
		return "", err
	}

	return config.RoleID, nil
}

func (e *Entity) ListGuildNFTRoleConfigs(guildID string) ([]model.GuildConfigGroupNFTRole, error) {
	return e.repo.GuildConfigGroupNFTRole.ListByGuildID(guildID)
}

func (e *Entity) ListMemberNFTRolesToAdd(guildID string) (map[[2]string]bool, error) {
	// 1. get nft balance of user address
	userAddressNFTBalance, err := e.repo.UserNFTBalance.GetUserNFTBalancesByUserInGuild(guildID)
	if err != nil {
		return nil, err
	}

	// 2. get balance of user discord
	userDiscordNFTBalance := make(map[string]int64)

	evmAccount, err := e.svc.MochiProfile.GetAllEvmAccount()
	if err != nil {
		return nil, err
	}

	for _, evmAcc := range evmAccount {
		for _, user := range userAddressNFTBalance {
			if evmAcc.PlatformIdentifier == user.UserAddress {
				profileId, err := e.svc.MochiProfile.GetByID(evmAcc.ProfileID)
				if err != nil {
					continue
				}

				for _, acc := range profileId.AssociatedAccounts {
					if acc.Platform == consts.PlatformDiscord {
						userDiscordNFTBalance[acc.PlatformIdentifier] += user.TotalBalance + user.StakingNeko
					}
				}

			}
		}
	}

	// 3. map role Id -> user discord
	rolesToAdd := make(map[[2]string]bool)

	guildConfigGroupNFTRole, err := e.ListGuildNFTRoleConfigs(guildID)
	if err != nil {
		return nil, err
	}

	for _, config := range guildConfigGroupNFTRole {
		for key, value := range userDiscordNFTBalance {
			if int(value) >= config.NumberOfTokens {
				rolesToAdd[[2]string{key, config.RoleID}] = true
			}
		}
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

func (e *Entity) EditMessageRepost(req *request.EditMessageRepostRequest) error {
	err := e.repo.MessageRepostHistory.EditMessageRepost(req)
	if err != nil {
		e.log.Errorf(err, "[e.EditMessageRepost] failed to edit message repost: %s", err)
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
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
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

func (e *Entity) GetTokenRole(id int) (*model.GuildConfigTokenRole, error) {
	return e.repo.GuildConfigTokenRole.Get(id)
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

func (e *Entity) ListAllConfigTokens(guildID string) ([]model.Token, error) {
	tokens, err := e.repo.GuildConfigTokenRole.ListAllTokenConfigs(guildID)
	if err != nil {
		e.log.Error(err, "[e.ListAllConfigTokens] - repo.GuildConfigTokenRole.ListAllTokenConfigs failed")
		return nil, err
	}
	return tokens, nil
}

func (e *Entity) ListTokenRoleConfigGuildIds() ([]string, error) {
	guildIds, err := e.repo.GuildConfigTokenRole.ListConfigGuildIds()
	if err != nil {
		e.log.Error(err, "[e.ListTokenRoleConfigGuildIds] - repo.GuildConfigTokenRole.ListConfigGuildIds failed")
		return nil, err
	}
	return guildIds, nil
}

func (e *Entity) GetGuildConfigTipRange(guildID string) (*response.GuildConfigTipRangeResponse, error) {
	data, err := e.repo.GuildConfigTipRange.GetByGuildID(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"guildID": guildID}).Info("[entities.GetGuildConfigTipRange] - repo.GuildConfigTipRange.GetByGuildID() not found")
			return nil, nil
		}
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entities.GetGuildConfigTipRange] - repo.GuildConfigTipRange.GetByGuildID() failed")
		return nil, err
	}

	return &response.GuildConfigTipRangeResponse{
		GuildID:   data.GuildID,
		Min:       data.Min,
		Max:       data.Max,
		UpdatedAt: data.UpdatedAt,
	}, err
}

func (e *Entity) UpsertGuildConfigTipRange(req request.UpsertGuildConfigTipRangeRequest) (*response.GuildConfigTipRangeResponse, error) {
	config, err := e.repo.GuildConfigTipRange.GetByGuildID(req.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.UpdateGuildConfigTipRange] - repo.GuildConfigTipRange.GetByGuildID() failed")
		return nil, err
	}
	//create if not exists
	if err == gorm.ErrRecordNotFound {
		config = &model.GuildConfigTipRange{
			GuildID: req.GuildID,
		}
	}
	if req.Max != nil {
		config.Max = req.Max
	}
	if req.Min != nil {
		config.Min = req.Min
	}

	if (config.Max != nil && config.Min != nil) && *config.Min > *config.Max {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.UpdateGuildConfigTipRange] - invalid amount")
		return nil, errors.ErrBadRequest
	}
	config.UpdatedAt = time.Now()

	data, err := e.repo.GuildConfigTipRange.UpsertOne(config)
	if err != nil {
		e.log.Fields(logger.Fields{"request": req}).Error(err, "[entities.UpdateGuildConfigTipRange] - repo.GuildConfigTipRange.UpsertOne() failed")
		return nil, err
	}

	return &response.GuildConfigTipRangeResponse{
		GuildID:   data.GuildID,
		Min:       data.Min,
		Max:       data.Max,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func (e *Entity) RemoveGuildConfigTipRange(guildID string) error {
	if err := e.repo.GuildConfigTipRange.Remove(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][RemoveGuildConfigTipRange] GuildConfigTipRange.Remove() failed")
		return err
	}

	return nil
}

func (e *Entity) CreateGuildConfigLogChannel(req request.CreateConfigLogChannelRequest) (*model.GuildConfigLogChannel, error) {
	return e.repo.GuildConfigLogChannel.Upsert(&model.GuildConfigLogChannel{
		GuildId:   req.GuildID,
		ChannelId: req.ChannelID,
		LogType:   req.LogType,
	})
}

func (e *Entity) GetGuildConfigLogChannel(req request.QueryConfigLogChannel) ([]model.GuildConfigLogChannel, error) {
	return e.repo.GuildConfigLogChannel.Get(query.Query{
		GuildId: req.GuildID,
		LogType: req.LogType,
	})
}
