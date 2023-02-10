package entities

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) ListGuildMixRoles(guildID string) ([]model.GuildConfigMixRole, error) {
	configs, err := e.repo.GuildConfigMixRole.ListByGuildID(guildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guildID": guildID,
		}).Error(err, "[e.ListGuildMixRoles] - repo.GuildConfigMixRole.ListByGuildID failed")
		return nil, err
	}
	return configs, nil
}

func (e *Entity) RemoveGuildMixRole(id int) error {
	if _, err := e.repo.GuildConfigMixRole.Get(id); err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[e.REmoveGuildMixRole] - repo.GuildConfigMixRole.Get failed")
		if err == gorm.ErrRecordNotFound {
			return errors.ErrRecordNotFound
		}
		return err
	}
	if err := e.repo.GuildConfigMixRole.Delete(id); err != nil {
		e.log.Fields(logger.Fields{
			"id": id,
		}).Error(err, "[e.RemoveGuildMixRole] - repo.GuildConfigMixRole.Delete failed")
		return err
	}
	return nil
}

func (e *Entity) ListMemberMixRolesToAdd(listConfigTokenRoles []model.GuildConfigMixRole, guildID string) (map[[2]string]bool, error) {
	mrs, err := e.repo.GuildConfigMixRole.GetMemberCurrentRoles(guildID)
	if err != nil {
		return nil, err
	}
	rolesToAdd := make(map[[2]string]bool)

	for _, mr := range mrs {
		rolesToAdd[[2]string{mr.UserDiscordID, mr.RoleID}] = true
	}
	return rolesToAdd, nil
}

func (e *Entity) CreateGuildMixRole(req request.CreateGuildMixRole) (*model.GuildConfigMixRole, error) {
	config, err := e.repo.GuildConfigMixRole.GetByRoleID(req.GuildID, req.RoleID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{
			"guildID": req.GuildID,
			"roleID":  req.RoleID,
		}).Error(err, "[e.CreateGuildMixRole] - repo.GuildConfigMixRole.GetByRoleID failed")
		return nil, err
	}
	// Record with role id and guild id already existed
	if err == nil {
		e.log.Fields(logger.Fields{
			"guildID": req.GuildID,
			"roleID":  req.RoleID,
		}).Error(err, "[e.CreateGuildMixRole] - Mix role config already existed")
		return nil, errors.ErrMixRoleExisted
	}
	tokenRequirement, err := e.getMixRoleTokenRequirement(req.TokenRequirement)
	if err != nil {
		return nil, err
	}
	nftRequirement, err := e.getMixRoleNFTRequirement(req.NFTRequirement)
	if err != nil {
		return nil, err
	}
	newConfig := &model.GuildConfigMixRole{
		GuildID:          req.GuildID,
		RoleID:           req.RoleID,
		RequiredLevel:    req.RequiredLevel,
		TokenRequirement: tokenRequirement,
		NFTRequirement:   nftRequirement,
	}
	if err := e.repo.GuildConfigMixRole.Create(newConfig); err != nil {
		e.log.Fields(logger.Fields{
			"config": config,
		}).Error(err, "[e.CreateGuildMixRole] - repo.GuildConfigMixRole.Create failed")
		return nil, err
	}
	return newConfig, nil
}

func (e *Entity) getMixRoleTokenRequirement(req *request.MixRoleTokenRequirement) (*model.MixRoleTokenRequirement, error) {
	if req == nil {
		return nil, nil
	}
	token, err := e.repo.Token.Get(req.TokenID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"tokenID": req.TokenID,
		}).Error(err, "[e.getMixRoleTokenRequirement] - repo.Token.Get failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrTokenNotFound
		}
		return nil, err
	}
	return &model.MixRoleTokenRequirement{
		TokenID:        token.ID,
		RequiredAmount: req.Amount,
	}, nil
}

func (e *Entity) getMixRoleNFTRequirement(req *request.MixRoleNFTRequirement) (*model.MixRoleNFTRequirement, error) {
	if req == nil {
		return nil, nil
	}
	nft, err := e.repo.NFTCollection.GetByID(req.NftID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"nftCollectionID": req.NftID,
		}).Error(err, "[e.getMixRoleNFTRequirement] - repo.NFTCollection.GetByID failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrTokenNotFound
		}
		return nil, err
	}
	return &model.MixRoleNFTRequirement{
		NFTCollectionID: nft.ID.UUID.String(),
		RequiredAmount:  req.Amount,
	}, nil
}
