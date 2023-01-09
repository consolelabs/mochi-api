package entities

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) CreateDaoProposal(req *request.CreateDaoProposalRequest) (*model.DaoProposal, error) {
	config, err := e.repo.GuildConfigDaoProposal.GetByGuildId(req.GuildId)
	if err != nil {
		e.log.Fields(logger.Fields{"GuildId": req.GuildId}).Error(err, "[entities.CreateDaoProposal][repo.GuildConfigDaoProposal.GetByGuildID] - failed to get guild config")
		return nil, err
	}

	daoProposal, err := e.repo.DaoProposal.Create(&model.DaoProposal{
		GuildId:                  req.GuildId,
		GuildConfigDaoProposalId: config.Id,
		VotingChannelId:          req.VotingChannelId,
		CreatorId:                req.CreatorId,
		Title:                    req.Title,
		Description:              req.Description,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposal.Create()] - failed to create dao proposal")
		return nil, err
	}

	proposalVoteOption := model.DaoProposalVoteOption{
		ProposalId:     daoProposal.Id,
		Address:        config.Address,
		ChainId:        config.ChainID,
		Symbol:         config.Symbol,
		RequiredAmount: config.RequiredAmount,
	}
	if req.VoteOption != nil {
		proposalVoteOption.VoteOptionId = &req.VoteOption.Id
		proposalVoteOption.Address = req.VoteOption.Address
		proposalVoteOption.ChainId = req.VoteOption.ChainId
		proposalVoteOption.Symbol = req.VoteOption.Symbol
		proposalVoteOption.RequiredAmount = strconv.FormatInt(req.VoteOption.RequiredAmount, 10)
	}

	_, err = e.repo.DaoProposalVoteOption.Create(&proposalVoteOption)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposalVoteOption.Create()] - failed to create dao proposal vote option")
		return nil, err
	}

	discussionChannel, err := e.svc.Discord.CreateDiscussionChannelForProposal(req.GuildId, req.Title)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][svc.Discord.CreateDiscussionChannelForProposal()] - failed to create discussion channel")
		return nil, err
	}

	err = e.repo.DaoProposal.UpdateDiscussionChannel(daoProposal.Id, discussionChannel)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposal.UpdateDiscussionChannel()] - failed to update discussion channel")
		return nil, err
	}

	daoProposal.DiscussionChannelId = discussionChannel

	return daoProposal, nil
}

func (e *Entity) DeleteDaoProposal(proposalId string) error {
	proposalIdNumber, err := strconv.ParseInt(proposalId, 10, 64)
	if err != nil {
		e.log.Error(err, "[Entity][DeleteDaoProposal] proposal_id is invalid")
		return errs.ErrInvalidProposalID
	}

	proposal, err := e.repo.DaoProposal.GetById(proposalIdNumber)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposal.GetById] - failed to get DAO proposal")
		return err
	}

	err = e.repo.DaoProposalVoteOption.DeleteAllByProposalID(proposal.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposalVoteOption.DeleteAllByProposalID()] - failed to clear dao proposal vote options")
		return err
	}

	err = e.repo.DaoProposal.DeleteById(proposal.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][repo.DaoProposal.DeleteById()] - failed to delete dao proposal")
		return err
	}

	err = e.svc.Discord.DeleteChannel(proposal.DiscussionChannelId)
	if err != nil {
		e.log.Fields(logger.Fields{"proposalId": proposalId}).Error(err, "[entities.DeleteDaoProposal][svc.Discord.DeleteChannel()] - failed to delete discussion channel")
	}

	return nil
}

func (e *Entity) TokenHolderStatus(query request.TokenHolderStatusRequest) (*response.TokenHolderStatus, error) {
	userWallet, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(query.UserID, query.GuildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &response.TokenHolderStatus{
				Data: &response.TokenHolderStatusData{IsWalletConnected: false},
			}, nil
		}
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(err, "[entities.TokenHolderStatus] - repo.UserWallet.GetOneByDiscordIDAndGuildID failed")
		return nil, err
	}
	// Not connect to wallet
	if userWallet.Address == "" {
		return &response.TokenHolderStatus{
			Data: &response.TokenHolderStatusData{IsWalletConnected: false},
		}, nil
	}

	// Connected to wallet, check another criteria for the action
	switch query.Action {
	case request.CreateProposal:
		return e.tokenHolderStatusForCreatingProposal(userWallet.Address, query)
	case request.Vote:
		return e.tokenHolderStatusForVoting(userWallet.Address, query)
	default:
		// invalid action or not supported
		e.log.Fields(logger.Fields{
			"action": query.Action,
		}).Error(errs.ErrBadRequest, "[entities.TokenHolderStatus] - invalid action")
		return nil, errs.ErrBadRequest
	}
}

func (e *Entity) tokenHolderStatusForCreatingProposal(walletAddress string, query request.TokenHolderStatusRequest) (*response.TokenHolderStatus, error) {
	if query.GuildID == "" {
		return nil, errs.ErrInvalidDiscordGuildID
	}
	if query.GuidelineChannelID == "" {
		return nil, errs.ErrInvalidDiscordChannelID
	}

	config, err := e.repo.GuildConfigDaoProposal.
		GetByGuildIDAndGuideLineChannelID(query.GuildID, query.GuidelineChannelID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(err, "[entities.TokenHolderStatus] - repo.GuildConfigDaoProposal.GetByGuildIDAndGuidelineChannelID failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	if config.Authority != model.TokenHolder {
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(errs.ErrInvalidAuthorityType, "[entities.TokenHolderStatus] - authority is not token holder")
		return nil, errs.ErrInvalidAuthorityType
	}

	if config.Type == nil {
		e.log.Fields(logger.Fields{
			"configID": config.Id,
		}).Error(errs.ErrInternalError, "[entities.TokenHolderStatus] - proposal voting type is nil")
		return nil, fmt.Errorf("config type data is mismatch")
	}

	userBalance, err := e.calculateUserBalance(*config.Type, walletAddress, config.Address, config.ChainID)
	if err != nil {
		return nil, err
	}
	requiredAmountBigInt, ok := new(big.Int).SetString(config.RequiredAmount, 10)
	if !ok {
		err = fmt.Errorf("cannot convert big int from string")
		e.log.Fields(logger.Fields{
			"requiredAmount": config.RequiredAmount,
			"base":           10,
		}).Error(err, "[entities.TokenHolderStatus] - new(big.Int).SetString failed")
		return nil, err
	}
	isQualified := userBalance.Cmp(requiredAmountBigInt) != -1
	userHoldingAmount := userBalance.Text(10)
	return &response.TokenHolderStatus{
		Data: &response.TokenHolderStatusData{
			IsWalletConnected: true,
			IsQualified:       &isQualified,
			UserHoldingAmount: &userHoldingAmount,
			GuildConfig:       config,
		},
	}, nil
}

func (e *Entity) tokenHolderStatusForVoting(walletAddress string, query request.TokenHolderStatusRequest) (*response.TokenHolderStatus, error) {
	if query.ProposalID == "" {
		return nil, errs.ErrInvalidProposalID
	}
	proposalIDNum, err := strconv.ParseInt(query.ProposalID, 10, 64)
	if err != nil {
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(errs.ErrInvalidProposalID, "[entities.TokenHolderStatus] - repo.GuildConfigDaoProposal.GetByGuildIDAndGuidelineChannelID failed")
		return nil, errs.ErrInvalidProposalID
	}
	config, err := e.repo.DaoProposalVoteOption.GetOneByProposalID(proposalIDNum)
	if err != nil {
		e.log.Fields(logger.Fields{
			"proposalID": query.ProposalID,
		}).Error(errs.ErrInvalidProposalID, "[entities.TokenHolderStatus] - repo.DaoProposalVoteOption.GetOneByProposalID failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}
	voteOption := config.VoteOption
	if voteOption == nil {
		e.log.Fields(logger.Fields{
			"voteOptionID": config.Id,
		}).Error(errs.ErrInternalError, "[entities.TokenHolderStatus] - vote option is nil")
		return nil, fmt.Errorf("vote option of id %v is nil", config.Id)
	}

	userBalance, err := e.calculateUserBalance(voteOption.Type, walletAddress, config.Address, config.ChainId)
	if err != nil {
		return nil, err
	}
	requiredAmountBigInt, ok := new(big.Int).SetString(config.RequiredAmount, 10)
	if !ok {
		err = fmt.Errorf("cannot convert big int from string")
		e.log.Fields(logger.Fields{
			"requiredAmount": config.RequiredAmount,
			"base":           10,
		}).Error(err, "[entities.TokenHolderStatus] - new(big.Int).SetString failed")
		return nil, err
	}
	isQualified := userBalance.Cmp(requiredAmountBigInt) != -1
	userHoldingAmount := userBalance.Text(10)
	return &response.TokenHolderStatus{
		Data: &response.TokenHolderStatusData{
			IsWalletConnected: true,
			IsQualified:       &isQualified,
			UserHoldingAmount: &userHoldingAmount,
			VoteConfig:        config,
		},
	}, nil
}

func (e *Entity) calculateUserBalance(votingType model.ProposalVotingType, walletAddress, tokenAddress string, chainId int64) (*big.Int, error) {
	chainIdStr := strconv.FormatInt(chainId, 10)
	var balanceOf func(string) (*big.Int, error)
	var err error
	switch votingType {
	case model.NFT:
		nftConfig := model.NFTCollectionConfig{
			ChainID:   chainIdStr,
			Address:   tokenAddress,
			ERCFormat: "erc721",
		}
		balanceOf, err = e.GetNFTBalanceFunc(nftConfig)
		if err != nil {
			e.log.Fields(logger.Fields{
				"config": nftConfig,
			}).Error(err, "[entities.TokenHolderStatus] - e.GetNFTBalanceFunc failed")
			return nil, err
		}
	case model.CryptoToken:
		balanceOf, err = e.GetTokenBalanceFunc(chainIdStr, tokenAddress)
		if err != nil {
			e.log.Fields(logger.Fields{
				"chainID": chainIdStr,
				"address": tokenAddress,
			}).Error(err, "[entities.TokenHolderStatus] - e.GetTokenBalanceFunc failed")
			return nil, err
		}
	}
	balance, err := balanceOf(walletAddress)
	if err != nil {
		e.log.Fields(logger.Fields{
			"wallletAddress": walletAddress,
		}).Error(err, "[entities.TokenHolderStatus] - get user balance failed")
		return nil, err
	}
	return balance, err
}
