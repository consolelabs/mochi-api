package entities

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common/math"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/chain"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
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

	var nftOptId int64 = 1
	var tokenOptId int64 = 2
	var optionId *int64
	if config.Type != nil && *config.Type == model.NFT {
		optionId = &nftOptId
	}
	if config.Type != nil && *config.Type == model.CryptoToken {
		optionId = &tokenOptId
	}

	token, err := e.repo.Token.GetByAddress(config.Address, int(config.ChainID))
	if err != nil {
		e.log.Fields(logger.Fields{
			"walletAddress": config.Address,
			"chainId":       config.ChainID,
		}).Error(err, "[entities.CreateDaoProposal] - repo.Token.GetByAddress failed")
		return nil, err
	}
	requiredAmtBig := big.NewInt(1).Mul(big.NewInt(1), math.BigPow(10, int64(token.Decimals))).Text(10)
	proposalVoteOption := model.DaoProposalVoteOption{
		ProposalId:     daoProposal.Id,
		Address:        config.Address,
		ChainId:        config.ChainID,
		Symbol:         config.Symbol,
		VoteOptionId:   optionId,
		RequiredAmount: requiredAmtBig,
	}

	_, err = e.repo.DaoProposalVoteOption.Create(&proposalVoteOption)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.CreateDaoProposal][repo.DaoProposalVoteOption.Create()] - failed to create dao proposal vote option")
		return nil, err
	}

	discussionChannel, err := e.svc.Discord.CreateDiscussionChannelForProposal(req.GuildId, req.VotingChannelId, req.Title)
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
	// userWallet, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(query.UserID, query.GuildID)
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return &response.TokenHolderStatus{
	// 			Data: &response.TokenHolderStatusData{IsWalletConnected: false},
	// 		}, nil
	// 	}
	// 	e.log.Fields(logger.Fields{
	// 		"userID":  query.UserID,
	// 		"guildID": query.GuildID,
	// 	}).Error(err, "[entities.TokenHolderStatus] - repo.UserWallet.GetOneByDiscordIDAndGuildID failed")
	// 	return nil, err
	// }
	// Not connect to wallet
	// if userWallet.Address == "" {
	// 	return &response.TokenHolderStatus{
	// 		Data: &response.TokenHolderStatusData{IsWalletConnected: false},
	// 	}, nil
	// }
	userWalletAddress := ""

	// Connected to wallet, check another criteria for the action
	switch query.Action {
	case request.CreateProposal:
		return e.tokenHolderStatusForCreatingProposal(userWalletAddress, query)
	case request.Vote:
		return e.tokenHolderStatusForVoting(userWalletAddress, query)
	default:
		// invalid action or not supported
		e.log.Fields(logger.Fields{
			"action": query.Action,
		}).Error(errs.ErrBadRequest, "[entities.TokenHolderStatus] - invalid action")
		return nil, errs.ErrBadRequest
	}
}

// We still keep using user_wallet here to help keep backward compatibility
// In the future when all user data is migrated to MochiProfile service, we can remove the user_wallet
func (e *Entity) tokenHolderStatusForCreatingProposal(walletAddress string, query request.TokenHolderStatusRequest) (*response.TokenHolderStatus, error) {
	if query.GuildID == "" {
		return nil, errs.ErrInvalidDiscordGuildID
	}

	config, err := e.repo.GuildConfigDaoProposal.GetByGuildId(query.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(err, "[entities.TokenHolderStatus] - repo.GuildConfigDaoProposal.GetByGuildID failed")
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	if config.Authority != model.TokenHolder && config.Authority != model.Admin {
		e.log.Fields(logger.Fields{
			"userID":  query.UserID,
			"guildID": query.GuildID,
		}).Error(errs.ErrInvalidAuthorityType, "[entities.TokenHolderStatus] - authority is not token holder")
		return nil, errs.ErrInvalidAuthorityType
	}

	isQualified := true
	if config.Type == nil {
		return &response.TokenHolderStatus{
			Data: &response.TokenHolderStatusData{
				IsWalletConnected: true,
				IsQualified:       &isQualified,
				GuildConfig:       config,
			},
		}, nil
	}

	userBalance, err := e.calculateUserBalance(*config.Type, walletAddress, query.UserID, config.Address, config.ChainID)
	if err != nil {
		return nil, err
	}
	e.log.Infof("USER BALANCE %s", userBalance.Text(10))
	requiredAmountBigInt, ok := new(big.Int).SetString(config.RequiredAmount, 10)
	if !ok {
		err = fmt.Errorf("cannot convert big int from string")
		e.log.Fields(logger.Fields{
			"requiredAmount": config.RequiredAmount,
			"base":           10,
		}).Error(err, "[entities.TokenHolderStatus] - new(big.Int).SetString failed")
		return nil, err
	}
	isQualified = userBalance.Cmp(requiredAmountBigInt) != -1
	userHoldingAmount := userBalance.Text(10)
	token, err := e.repo.Token.GetByAddress(config.Address, int(config.ChainID))
	if err != nil {
		e.log.Fields(logger.Fields{
			"walletAddress": config.Address,
			"chainId":       config.ChainID,
		}).Error(err, "[entities.TokenHolderStatus] - repo.Token.GetByAddress failed")
		return nil, err
	}
	// Convert to floating type
	requiredAmtFloat := new(big.Float).SetInt(requiredAmountBigInt)
	decimalsFloat := new(big.Float).SetInt(math.BigPow(10, int64(token.Decimals)))
	requiredAmtText, _ := requiredAmtFloat.Quo(requiredAmtFloat, decimalsFloat).Float64()
	config.RequiredAmount = fmt.Sprintf("%.2f", requiredAmtText)
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
	if query.ProposalID == nil {
		return nil, errs.ErrInvalidProposalID
	}
	config, err := e.repo.DaoProposalVoteOption.GetOneByProposalID(*query.ProposalID)
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
	isQualified := true
	if voteOption == nil {
		return &response.TokenHolderStatus{
			Data: &response.TokenHolderStatusData{
				IsWalletConnected: true,
				IsQualified:       &isQualified,
				VoteConfig:        config,
			},
		}, nil
	}

	userBalance, err := e.calculateUserBalance(voteOption.Type, walletAddress, query.UserID, config.Address, config.ChainId)
	if err != nil {
		return nil, err
	}
	// Just need user balance > 0
	isQualified = userBalance.Cmp(big.NewInt(0)) == 1

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

func (e *Entity) calculateUserBalance(votingType model.ProposalVotingType, walletAddress, discordID, tokenAddress string, chainId int64) (*big.Int, error) {
	switch votingType {
	case model.NFT:
		return e.calculateNFTBalance(chainId, tokenAddress, walletAddress)
	case model.CryptoToken:
		return e.CalculateTokenBalance(chainId, tokenAddress, discordID)
	default:
		return nil, fmt.Errorf("invalid voting type")
	}
}

func (e *Entity) calculateNFTBalance(chainId int64, tokenAddress, walletAddress string) (*big.Int, error) {
	chainIdStr := strconv.FormatInt(chainId, 10)
	nftConfig := model.NFTCollectionConfig{
		ChainID:   chainIdStr,
		Address:   tokenAddress,
		ERCFormat: "erc721",
	}
	balanceOf, err := e.GetNFTBalanceFunc(nftConfig)
	if err != nil {
		e.log.Fields(logger.Fields{
			"config": nftConfig,
		}).Error(err, "[entities.TokenHolderStatus] - e.GetNFTBalanceFunc failed")
		return nil, err
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

func (e *Entity) CalculateTokenBalance(chainId int64, tokenAddress, discordID string) (*big.Int, error) {
	profiles, err := e.svc.MochiProfile.GetByDiscordID(discordID, true)
	if err != nil {
		e.log.Fields(logger.Fields{"discordID": discordID}).Error(err, "cannot get mochi profile")
		return nil, err
	}
	includedPlatform := mochiprofile.PlatformEVM
	if chainId == 999 {
		includedPlatform = mochiprofile.PlatformSol
	}
	var walletAddrs []string
	for _, p := range profiles.AssociatedAccounts {
		if p.Platform == includedPlatform {
			walletAddrs = append(walletAddrs, p.PlatformIdentifier)
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(walletAddrs))
	bals := make(chan *big.Int)
	go func() {
		wg.Wait()
		close(bals)
	}()

	// Fetch balance concurrently
	for _, addr := range walletAddrs {
		go func(chainId int64, tokenAddress, currentWallet string) {
			defer wg.Done()
			bal, err := e.fetchTokenBalanceByChain(chainId, tokenAddress, currentWallet)
			if err == nil {
				bals <- bal
			} else {
				e.log.Fields(logger.Fields{"discordID": discordID, "tokenAddr": tokenAddress}).Error(err, "fetchTokenBalanceByChain() failed")
			}
		}(chainId, tokenAddress, addr)
	}

	counter := 0
	totalBalances := big.NewInt(0)
	for b := range bals {
		totalBalances = big.NewInt(0).Add(totalBalances, b)
		counter += 1
	}

	if counter < len(walletAddrs) {
		return nil, fmt.Errorf("error while fetching balance - user %s", discordID)
	}

	return totalBalances, nil
}

func (e *Entity) fetchTokenBalanceByChain(chainId int64, tokenAddress, walletAddress string) (*big.Int, error) {
	log := e.log.Fields(logger.Fields{"chainID": chainId, "tokenAddress": tokenAddress, "walletAddress": walletAddress})
	switch chainId {
	case 999: // SOL
		client := chain.NewSolanaClient(&e.cfg, e.log, nil)
		bal, err := client.GetTokenBalance(walletAddress, tokenAddress)
		if err != nil {
			log.Error(err, "[e.fetchTokenbalanceByChain] solClient.GetTokenBalance failed")
			return nil, err
		}
		return bal, nil
	case 1, 10, 56, 137, 250, 42161: //EVM
		token := model.Token{
			Address: tokenAddress,
			ChainID: int(chainId),
		}
		balanceOf, err := e.GetTokenBalanceFunc(strconv.FormatInt(chainId, 10), token)
		if err != nil {
			log.Error(err, "[e.fetchTokenBalanceByChain] - e.GetTokenBalanceFunc failed")
			return nil, err
		}
		balance, err := balanceOf(walletAddress)
		if err != nil {
			log.Error(err, "[e.fetchTokenBalanceByChain] - get user balance failed")
			return nil, err
		}
		return balance, err
	default:
		return nil, fmt.Errorf("chain is not supported")
	}
}

// Get all threads in a community
func (e *Entity) GetAllCommonwealthThreads(communityId string) (*response.CommonwealthThreadResponse, error) {
	return e.svc.Commonwealth.GetThreads(communityId)
}

func (e *Entity) GetAllCommonwealthData() ([]model.CommonwealthLatestData, error) {
	return e.repo.CommonwealthLatestData.GetAll()
}

func (e *Entity) UpdateCommonwealthData(model model.CommonwealthLatestData) error {
	return e.repo.CommonwealthLatestData.UpsertOne(model)
}

func (e *Entity) GetProposalUsage(page string, size string) (*response.GuildProposalUsageResponse, error) {
	res := []response.GuildProposalUsageData{}
	p, _ := strconv.Atoi(page)
	sz, _ := strconv.Atoi(size)
	ppsCount, total, err := e.repo.DaoProposal.GetUsageStatsWithPaging(p, sz)
	if err != nil {
		e.log.Errorf(err, "[e.GetProposalUsage] failed to get proposal count")
		return nil, err
	}
	for _, pps := range *ppsCount {
		active := true
		cfg, err := e.repo.GuildConfigDaoProposal.GetByGuildId(pps.GuildId)
		// has proposals but no config found -> deleted config
		if err == gorm.ErrRecordNotFound || cfg == nil {
			active = false
		}
		res = append(res, response.GuildProposalUsageData{
			GuildId:       pps.GuildId,
			GuildName:     pps.GuildName,
			ProposalCount: pps.Count,
			IsActive:      active,
		})
	}
	return &response.GuildProposalUsageResponse{
		Pagination: response.PaginationResponse{
			Pagination: model.Pagination{
				Page: int64(p),
				Size: int64(sz),
			},
			Total: total,
		},
		Data: &res,
	}, nil
}

func (e *Entity) GetDaoTrackerMetric(page string, size string) (*response.DaoTrackerSpaceCountResponse, error) {
	p, _ := strconv.Atoi(page)
	sz, _ := strconv.Atoi(size)
	spaceCount, total, err := e.repo.GuildConfigDaoTracker.GetUsageStatsWithPaging(p, sz)
	if err != nil {
		e.log.Errorf(err, "[e.GetProposalUsage] failed to get proposal count")
		return nil, err
	}
	return &response.DaoTrackerSpaceCountResponse{
		Pagination: response.PaginationResponse{
			Pagination: model.Pagination{
				Page: int64(p),
				Size: int64(sz),
			},
			Total: total,
		},
		Data: &spaceCount,
	}, nil
}
