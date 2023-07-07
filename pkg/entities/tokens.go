package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/chain"
	chainpkg "github.com/defipod/mochi/pkg/chain"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo/token"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) CreateCustomToken(req request.UpsertCustomTokenConfigRequest) error {
	chain, err := e.repo.Chain.GetByShortName(req.Chain)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.Chain.GetByShortName failed")
		return fmt.Errorf("error getting chain: %v", err)
	}

	coins, err := e.SearchCoins(req.Symbol, "")
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] svc.CoinGecko.SearchCoins failed")
		return fmt.Errorf("error seaching coin: %v", err)
	}
	if len(coins) == 0 {
		return fmt.Errorf("cannot find token by symbol")
	}

	coin, err, _ := e.svc.CoinGecko.GetCoin(coins[0].ID)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] svc.CoinGecko.GetCoin failed")
		return fmt.Errorf("error getting coin: %v", err)
	}
	if coin.AssetPlatformID != chain.CoinGeckoID {
		e.log.
			Fields(logger.Fields{"asset_platform_id": coin.AssetPlatformID, "coin_gecko_id": chain.CoinGeckoID}).
			Error(nil, "[Entity][CreateCustomToken] token is not supported on this chain")
		return fmt.Errorf("token is not supported on this chain")
	}

	historicalTokenPrices, err, statusCode := e.svc.Covalent.GetHistoricalTokenPrices(chain.ID, chain.Currency, req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"statusCode": statusCode}).Error(err, "[Entity][CreateCustomToken] svc.Covalent.GetHistoricalTokenPrices failed")
		return err
	}

	token, err := e.repo.Token.GetOneBySymbol(req.Symbol)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.Token.GetOneBySymbol failed")
	}

	if token == nil {
		token = &model.Token{
			ChainID:             chain.ID,
			Address:             req.Address,
			CoinGeckoID:         coins[0].ID,
			Name:                coins[0].Name,
			Symbol:              strings.ToUpper(coins[0].Symbol),
			Decimals:            historicalTokenPrices.Data[0].Decimals,
			DiscordBotSupported: true,
			Chain:               chain,
		}
		if err = e.repo.Token.CreateOne(token); err != nil {
			e.log.Error(err, "[Entity][CreateCustomToken] repo.Token.CreateOne failed")
			return err
		}
	}

	if !token.DiscordBotSupported {
		return errors.New("token is not supported to add to your server")
	}

	gct, err := e.repo.GuildConfigToken.GetByGuildIDAndTokenID(req.GuildID, token.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.GuildConfigToken.GetByGuildIDAndTokenID failed")
		return err
	}

	if gct == nil {
		if err := e.repo.GuildConfigToken.CreateOne(model.GuildConfigToken{
			GuildID: req.GuildID,
			TokenID: token.ID,
			Active:  true,
		}); err != nil {
			e.log.Error(err, "[Entity][CreateCustomToken] repo.GuildConfigToken.CreateOne failed")
			return err
		}
	}

	return nil
}

func (e *Entity) GetAllSupportedToken(guildID string) (returnToken []model.Token, err error) {
	returnToken, err = e.repo.Token.GetSupportedTokenByGuildId(guildID)
	if err != nil {
		return returnToken, err
	}

	return returnToken, nil
}

func (e *Entity) GetSupportedToken(address, chain string) (*model.Token, error) {
	chainId, err := strconv.Atoi(util.ConvertInputToChainId(chain))
	if err != nil {
		e.log.Fields(logger.Fields{"chain": chain}).Error(err, "[Entity][GetSupportedToken] strconv.Atoi failed")
		return nil, baseerrs.ErrInvalidChain
	}

	token, err := e.repo.Token.GetByAddress(address, chainId)
	if err != nil {
		e.log.Fields(logger.Fields{"address": address}).Error(err, "[Entity][GetSupportedToken] repo.Token.GetByAddress failed")
		if err == gorm.ErrRecordNotFound {
			return nil, baseerrs.ErrRecordNotFound
		}
		return nil, err
	}
	return token, nil
}

func (e *Entity) GetDefaultToken(guildID string) (*model.Token, error) {
	if _, err := e.repo.DiscordGuilds.GetByID(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][GetDefaultToken] repo.DiscordGuilds.GetByID failed")
		return nil, err
	}

	token, err := e.repo.Token.GetDefaultTokenByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][GetDefaultToken] repo.Token.GetDefaultTokenByGuildID failed")
		return nil, err
	}

	return &token, nil
}

func (e *Entity) SetDefaultToken(req request.ConfigDefaultTokenRequest) error {
	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[Entity][SetDefaultToken] repo.DiscordGuilds.GetByID failed")
		return err
	}

	token, err := e.repo.Token.GetBySymbol(req.Symbol, true)
	if err == gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[Entity][SetDefaultToken] repo.Token.GetBySymbol failed")
		return baseerrs.ErrRecordNotFound
	}
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[Entity][SetDefaultToken] repo.Token.GetBySymbol failed")
		return err
	}

	if err = e.repo.GuildConfigToken.UpsertOne(model.GuildConfigToken{
		GuildID:   req.GuildID,
		TokenID:   token.ID,
		Active:    true,
		IsDefault: true,
	}); err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "token_id": token.ID}).Error(err, "[Entity][SetDefaultToken] repo.GuildConfigToken.UpsertOne failed")
		return err
	}

	if err := e.repo.GuildConfigToken.UnsetOldDefaultToken(req.GuildID, token.ID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "token_id": token.ID}).Error(err, "[Entity][SetDefaultToken] repo.GuildConfigToken.SetDefaultToken failed")
		return err
	}

	return nil
}

func (e *Entity) RemoveDefaultToken(guildID string) error {
	if _, err := e.repo.DiscordGuilds.GetByID(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][RemoveDefaultToken] repo.DiscordGuilds.GetByID failed")
		return err
	}

	if err := e.repo.GuildConfigToken.RemoveDefaultToken(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][RemoveDefaultToken] repo.GuildConfigToken.RemoveDefaultToken failed")
		return err
	}

	return nil
}

func (e *Entity) TotalSupportedTokens(guildId string) (*response.Metric, error) {
	supportedTokens, supportedTokensCount, err := e.repo.Token.GetAllSupported(token.ListQuery{})
	if err != nil {
		e.log.Error(err, "[Entity][TotalSupportedTokens] repo.Token.GetAllSupported failed")
		return nil, err
	}

	serverSupportedTokens, err := e.repo.Token.GetSupportedTokenByGuildId(guildId)
	if err != nil {
		e.log.Error(err, "[Entity][TotalSupportedTokens] repo.Token.GetSupportedTokenByGuildId failed")
		return nil, err
	}

	serverToken := make([]string, 0)
	for _, token := range serverSupportedTokens {
		serverToken = append(serverToken, token.Symbol)
	}

	totalToken := make([]string, 0)
	for _, token := range supportedTokens {
		totalToken = append(totalToken, token.Symbol)
	}

	return &response.Metric{
		ServerTokenSupported: int64(len(serverSupportedTokens)),
		TotalTokenSupported:  supportedTokensCount,
		ServerToken:          serverToken,
		TotalToken:           totalToken,
	}, nil
}

func (e *Entity) GetTokensByChainID(chainID int) ([]model.Token, error) {
	return e.repo.Token.GetByChainID(chainID)
}

func (e *Entity) GetTokenBalanceFunc(chainID string, token model.Token) (func(address string) (*big.Int, error), error) {
	chainIDNum, err := strconv.Atoi(chainID)
	if err != nil {
		e.log.Errorf(err, "[strconv.Atoi] failed to convert chain %s to number", chainID)
		return nil, fmt.Errorf("failed to convert string to int: %v", err)
	}
	chain, err := e.repo.Chain.GetByID(chainIDNum)
	if err != nil {
		e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain by id %s", chainID)
		return nil, fmt.Errorf("failed to get chain by id %s: %v", chainID, err)
	}
	wallet, err := hdwallet.NewFromMnemonic(e.cfg.InDiscordWalletMnemonic)
	if err != nil {
		return nil, err
	}
	chainClient, err := chainpkg.NewClient(&e.cfg, wallet, e.log, chain.RPC, chain.APIKey, chain.APIBaseURL)
	if err != nil {
		return nil, err
	}

	balanceOf := func(address string) (*big.Int, error) {
		b, err := chainClient.RawErc20TokenBalance(address, token)
		if err != nil {
			e.log.Errorf(err, "[contractERC20.BalanceOf] failed to get balance of %s in chain %s", address, chainID)
			return nil, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, chainID, err.Error())
		}
		return b, nil
	}
	return balanceOf, nil
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
