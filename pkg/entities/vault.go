package entities

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/consolelabs/mochi-typeset/typeset"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/vault"
	vaulttxquery "github.com/defipod/mochi/pkg/repo/vault_transaction"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochipay"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) CreateVault(req *request.CreateVaultRequest) (*model.Vault, error) {
	// auto generate vault address when desig mode = false
	walletAddress, solanaWalletAddress := "", ""
	walletNumber := -1
	if !req.DesigMode {
		// get latest wallet number in db
		latestWalletNumber, err := e.repo.Vault.GetLatestWalletNumber()
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.repo.Vault.GetLatestWalletNumber failed")
			return nil, err
		}

		// generate evm wallet
		account, err := e.vaultwallet.GetAccountByWalletNumber(int(latestWalletNumber.Int64) + 1)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.vaultwallet.GetAccountByWalletNumber failed")
			return nil, err
		}
		walletAddress = account.Address.Hex()

		// generate solana wallet
		solanaWallet, err := e.vaultwallet.GetAccountSolanaByWalletNumber(int(latestWalletNumber.Int64) + 1)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.vaultwallet.GetAccountSolanaByWalletNumber failed")
			return nil, err
		}
		solanaWalletAddress = solanaWallet.PublicKey.ToBase58()

		walletNumber = int(latestWalletNumber.Int64) + 1
	}

	vault, err := e.repo.Vault.Create(&model.Vault{
		GuildId:             req.GuildId,
		Name:                req.Name,
		Threshold:           req.Threshold,
		WalletAddress:       walletAddress,
		WalletNumber:        int64(walletNumber),
		SolanaWalletAddress: solanaWalletAddress,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - e.repo.Vault.Create failed")
		return nil, err
	}

	userDiscordID := ""
	profile, err := e.svc.MochiProfile.GetByID(req.VaultCreator)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetVaults] svc.MochiProfile.GetByID() failed")
		return nil, err
	}

	for _, acc := range profile.AssociatedAccounts {
		if acc.Platform == mochiprofile.PlatformDiscord {
			userDiscordID = acc.PlatformIdentifier
		}
	}

	// default for vault creator will be added as treasurer
	_, err = e.repo.VaultTreasurer.Create(&model.VaultTreasurer{
		VaultId:       vault.Id,
		GuildId:       req.GuildId,
		Role:          consts.VaultCreatorRole,
		UserProfileId: req.VaultCreator,
		UserDiscordId: userDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - add treasurer failed")
		return nil, err
	}

	return vault, nil
}

func (e *Entity) GetVaults(req request.GetVaultsRequest) ([]model.Vault, error) {
	listQuery := vault.ListQuery{
		GuildID:       req.GuildID,
		EvmWallet:     req.EvmAddress,
		SolanaWallet:  req.SolanaAddress,
		Threshold:     req.Threshold,
		UserProfileID: req.ProfileID,
	}

	// query db
	vaults, err := e.repo.Vault.List(listQuery)
	if err != nil {
		e.log.Fields(logger.Fields{"query": listQuery}).Errorf(err, "[entity.GetVaults] repo.Vault.List() failed")
		return nil, err
	}

	if req.NoFetchAmount != "true" {
		for i, vault := range vaults {
			walletAssetsEVM, _, _, err := e.ListWalletAssets(request.ListWalletAssetsRequest{Type: "eth", Address: vault.WalletAddress})
			if err != nil {
				e.log.Fields(logger.Fields{"vault": vault}).Errorf(err, "[entity.GetVaults] e.ListWalletAssets() failed")
				return nil, err
			}
			vaults[i].TotalAmountEVM = fmt.Sprintf("%.4f", sumBal(walletAssetsEVM))

			walletAssetsSolana, _, _, err := e.ListWalletAssets(request.ListWalletAssetsRequest{Type: "sol", Address: vault.SolanaWalletAddress})
			if err != nil {
				e.log.Fields(logger.Fields{"vault": vault}).Errorf(err, "[entity.GetVaults] e.ListWalletAssets() failed")
				vaults[i].TotalAmountSolana = "0"
			}
			if len(walletAssetsSolana) > 0 {
				vaults[i].TotalAmountSolana = fmt.Sprintf("%.4f", sumBal(walletAssetsSolana))
			}
		}
	}

	return vaults, nil
}

func sumBal(walletAssets []response.WalletAssetData) (sum float64) {
	for _, asset := range walletAssets {
		sum += asset.UsdBalance
	}
	return
}

func (e *Entity) GetVaultConfigChannel(guildId string) (*model.VaultConfig, error) {
	vaultConfig, err := e.repo.VaultConfig.GetByGuildId(guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return vaultConfig, nil
}

func (e *Entity) CreateVaultConfigChannel(req *request.CreateVaultConfigChannelRequest) error {
	return e.repo.VaultConfig.Create(&model.VaultConfig{
		GuildId:   req.GuildId,
		ChannelId: req.ChannelId,
	})
}

func (e *Entity) CreateConfigThreshold(req *request.CreateConfigThresholdRequest) (*model.Vault, error) {
	vault, err := e.repo.Vault.GetByNameAndGuildId(req.Name, req.GuildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not found")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateConfigThreshold] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	_, err = e.repo.Vault.UpdateThreshold(&model.Vault{
		GuildId:   req.GuildId,
		Name:      req.Name,
		Threshold: req.Threshold,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateConfigThreshold] - e.repo.Vault.UpdateThreshold failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId:   req.GuildId,
		VaultId:   vault.Id,
		Action:    consts.TreasurerConfigThresholdType,
		Threshold: req.Threshold,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}
	return vault, nil
}

func (e *Entity) AddTreasurerToVault(req *request.AddTreasurerToVaultRequest) (*model.VaultTreasurer, error) {
	userDiscordID := ""
	profile, err := e.svc.MochiProfile.GetByID(req.UserProfileID)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetVaults] svc.MochiProfile.GetByID() failed")
		return nil, err
	}

	for _, acc := range profile.AssociatedAccounts {
		if acc.Platform == mochiprofile.PlatformDiscord {
			userDiscordID = acc.PlatformIdentifier
		}
	}

	treasurer, err := e.repo.VaultTreasurer.Create(&model.VaultTreasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserProfileId: req.UserProfileID,
		UserDiscordId: userDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTreasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerAddType,
		Target:  req.UserProfileID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}

	return treasurer, nil
}

func (e *Entity) TransferVaultToken(req *request.TransferVaultTokenRequest) error {
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Vault.GetById failed")
		return err
	}

	treasurer, err := e.repo.VaultTreasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.VaultTreasurer.GetByVaultId failed")
		return err
	}

	listNotify := []string{}
	for _, t := range treasurer {
		listNotify = append(listNotify, t.UserProfileId)
	}

	token, err := e.svc.MochiPay.GetToken(req.Token, req.Chain)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.svc.MochiPay.GetToken failed")
		return err
	}

	treasurerRequest, err := e.repo.VaultRequest.GetById(req.RequestId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.VaultRequest.GetById failed")
		return err
	}

	if !slices.Contains(listNotify, treasurerRequest.UserProfileId) {
		listNotify = append(listNotify, treasurerRequest.UserProfileId)
	}

	amountBigIntStr := util.FloatToString(req.Amount, token.Decimal)

	validateBalance := e.validateBalance(token, vault.WalletAddress, vault.SolanaWalletAddress, req.Amount)
	if !validateBalance {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - validateBalance failed")
		return fmt.Errorf("balance not enough")
	}

	recipientPay := treasurerRequest.UserProfileId
	if recipientPay == "" {
		recipientPay = treasurerRequest.RequesterProfileId
	}

	// address = "" aka destination addres = "", use mochi wallet instead
	privateKey, destination := "", ""
	if token.Chain.ChainId == "999" {
		account, err := e.vaultwallet.GetAccountSolanaByWalletNumber(int(vault.WalletNumber))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetAccountSolanaByWalletNumber failed")
			return err
		}

		privateKey, err = e.vaultwallet.GetPrivateKeyByAccountSolana(*account)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetPrivateKeyByAccountSolana failed")
			return err
		}

		destination, _ = e.vaultwallet.SolanaCentralizedWalletAddress()
	} else {
		account, err := e.vaultwallet.GetAccountByWalletNumber(int(vault.WalletNumber))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetAccountByWalletNumber failed")
			return err
		}

		privateKey, err = e.vaultwallet.GetPrivateKeyByAccount(account)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.vaultwallet.GetPrivateKeyByAccount failed")
			return err
		}

		destination = e.cfg.CentralizedWalletAddress
	}

	if req.Address != "" {
		destination = req.Address
	}

	_, err = e.svc.MochiPay.TransferVaultMochiPay(request.MochiPayVaultRequest{
		ProfileId:  treasurerRequest.RequesterProfileId,
		Amount:     amountBigIntStr,
		To:         destination,
		PrivateKey: privateKey,
		Token:      token.Symbol,
		Chain:      token.Chain.ChainId,
		Name:       vault.Name,
		VaultId:    vault.Id,
		Reciever:   recipientPay,
		Message:    treasurerRequest.Message,
		ListNotify: listNotify,
		RequestId:  treasurerRequest.Id,
		Platform:   req.Platform,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.svc.MochiPay.TransferVaultMochiPay failed")
		return err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId:   req.GuildId,
		VaultId:   req.VaultId,
		Action:    consts.TreasurerTransferType,
		ToAddress: req.Address,
		Amount:    req.Amount,
		Token:     req.Token,
		Sender:    treasurerRequest.RequesterProfileId,
		Target:    treasurerRequest.UserProfileId,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return err
	}

	return nil
}

func (e *Entity) AutoTransferVaultToken(req *model.AutoTransferVaultTokenRequest) error {
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.repo.Vault.GetById failed")
		return err
	}

	treasurer, err := e.repo.VaultTreasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.repo.VaultTreasurer.GetByVaultId failed")
		return err
	}

	listNotify := []string{}
	for _, t := range treasurer {
		profileMember, err := e.svc.MochiProfile.GetByDiscordID(t.UserDiscordId, true)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.repo.Profile.GetByDiscordId failed")
			return err
		}
		listNotify = append(listNotify, profileMember.ID)
	}

	token, err := e.svc.MochiPay.GetToken(req.Token, req.Chain)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.svc.MochiPay.GetToken failed")
		return err
	}

	receiverProfile, err := e.svc.MochiProfile.GetByDiscordID(req.Target, true)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.svc.MochiProfile.GetByDiscordId failed")
		return err
	}

	if !slices.Contains(listNotify, receiverProfile.ID) {
		listNotify = append(listNotify, receiverProfile.ID)
	}

	amountBigIntStr := util.FloatToString(req.Amount, token.Decimal)

	validateBalance := e.validateBalance(token, vault.WalletAddress, vault.SolanaWalletAddress, req.Amount)
	if !validateBalance {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - validateBalance failed")
		return fmt.Errorf("balance not enough")
	}

	// address = "" aka destination addres = "", use mochi wallet instead
	privateKey, destination := "", ""
	if token.Chain.ChainId == "999" {
		account, err := e.vaultwallet.GetAccountSolanaByWalletNumber(int(vault.WalletNumber))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.vaultwallet.GetAccountSolanaByWalletNumber failed")
			return err
		}

		privateKey, err = e.vaultwallet.GetPrivateKeyByAccountSolana(*account)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.vaultwallet.GetPrivateKeyByAccountSolana failed")
			return err
		}

		destination, _ = e.vaultwallet.SolanaCentralizedWalletAddress()
	} else {
		account, err := e.vaultwallet.GetAccountByWalletNumber(int(vault.WalletNumber))
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.vaultwallet.GetAccountByWalletNumber failed")
			return err
		}

		privateKey, err = e.vaultwallet.GetPrivateKeyByAccount(account)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.vaultwallet.GetPrivateKeyByAccount failed")
			return err
		}

		destination = e.cfg.CentralizedWalletAddress
	}

	if req.Address != "" {
		destination = req.Address
	}

	_, err = e.svc.MochiPay.TransferVaultMochiPay(request.MochiPayVaultRequest{
		ProfileId:  receiverProfile.ID,
		Amount:     amountBigIntStr,
		To:         destination,
		PrivateKey: privateKey,
		Token:      token.Symbol,
		Chain:      token.Chain.ChainId,
		Name:       vault.Name,
		Message:    req.Message,
		ListNotify: listNotify,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.svc.MochiPay.TransferVaultMochiPay failed")
		return err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId:   req.GuildId,
		VaultId:   req.VaultId,
		Action:    consts.TreasurerTransferType,
		ToAddress: req.Address,
		Amount:    req.Amount,
		Token:     req.Token,
		Sender:    receiverProfile.ID, //TODO thanhpn change to id of auto bot
		Target:    req.Target,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AutoTransferVaultToken] - e.repo.VaultTransaction.Create failed")
		return err
	}

	return nil
}

func (e *Entity) CreateTreasurerRequest(req *request.CreateTreasurerRequest) (*response.CreateTreasurerRequestResponse, error) {
	// get vault from name and guild id
	vault, err := e.repo.Vault.GetByNameAndGuildId(req.VaultName, req.GuildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not exist")
		}
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	treasurers, err := e.repo.VaultTreasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTreasurer.GetByVaultId failed")
		return nil, err
	}

	if req.Type == "transfer" {
		token, err := e.svc.MochiPay.GetToken(req.Token, req.Chain)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerRequest] - e.svc.MochiPay.GetToken failed")
			return nil, err
		}

		validateBal := e.validateBalance(token, vault.WalletAddress, vault.SolanaWalletAddress, req.Amount)
		if !validateBal {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - validateBalance failed")
			return nil, fmt.Errorf("balance not enough")
		}
	}

	if req.Type == "remove" {
		if !e.validateTreasurer(treasurers, req.UserProfileId) {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerRequest] - user not in list treasurers")
			return nil, fmt.Errorf("user not in list treasurers")
		}
	}

	// create treasurer request
	treasurerReq, err := e.repo.VaultRequest.Create(&model.VaultRequest{
		GuildId:            req.GuildId,
		VaultId:            vault.Id,
		UserProfileId:      req.UserProfileId,
		Message:            req.Message,
		RequesterProfileId: req.RequesterProfileId,
		Type:               req.Type,
		Amount:             req.Amount,
		Chain:              req.Chain,
		Token:              req.Token,
		Address:            req.Address,
		MessageUrl:         req.MessageUrl,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTreasurer.Create failed")
		return nil, err
	}

	// add submission with status pending for all treasurer in vaul
	treasurerSubmission := make([]model.VaultSubmission, 0)

	for _, treasurer := range treasurers {
		status := consts.TreasurerSubmissionStatusPending
		if treasurer.UserProfileId == req.RequesterProfileId {
			status = consts.TreasurerSubmissionStatusApproved
		}

		treasurerSubmission = append(treasurerSubmission, model.VaultSubmission{
			VaultId:            vault.Id,
			GuildId:            req.GuildId,
			RequestId:          treasurerReq.Id,
			Status:             status,
			SubmitterProfileId: treasurer.UserProfileId,
			MessageUrl:         req.MessageUrl,
		})
	}

	err = e.repo.VaultSubmission.Create(treasurerSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultSubmission.Create failed")
		return nil, err
	}

	// there's 2 case here
	// - after the requester default approve the request, number of approved will pass the threshold -> execute action now
	// - or not pass the threshold -> send DM to treasurer about approve / reject button
	isDecidedAndExecuted, err := e.PostCreateTreasurerRequest(req, treasurerReq, vault, treasurers)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.PostCreateTreasurerRequest failed")
		return nil, err
	}

	return &response.CreateTreasurerRequestResponse{
		Request:              *treasurerReq,
		VaultTreasurer:       treasurers,
		IsDecidedAndExecuted: isDecidedAndExecuted,
	}, nil
}

func (e *Entity) PostCreateTreasurerRequest(req *request.CreateTreasurerRequest, treasurerRequest *model.VaultRequest, vault *model.Vault, treasurers []model.VaultTreasurer) (bool, error) {
	threshold, _ := strconv.ParseFloat(vault.Threshold, 64)
	percentage := float64(1) / float64(len(treasurers)) * 100

	if percentage >= threshold {
		// execute action
		switch req.Type {
		case "add":
			_, err := e.AddTreasurerToVault(&request.AddTreasurerToVaultRequest{
				GuildId:       req.GuildId,
				VaultId:       vault.Id,
				UserProfileID: req.UserProfileId,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.AddTreasurerToVault failed")
				return false, err
			}
		case "remove":
			_, err := e.RemoveTreasurerFromVault(&request.RemoveTreasurerToVaultRequest{
				GuildId:       req.GuildId,
				VaultId:       vault.Id,
				UserProfileID: req.UserProfileId,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.RemoveTreasurerFromVault failed")
				return false, err
			}
		case "transfer":
			err := e.TransferVaultToken(&request.TransferVaultTokenRequest{
				GuildId:   req.GuildId,
				VaultId:   vault.Id,
				RequestId: treasurerRequest.Id,
				Address:   req.Address,
				Amount:    req.Amount,
				Token:     req.Token,
				Chain:     req.Chain,
				Target:    req.UserProfileId,
				Platform:  req.Platform,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.TransferVaultToken failed")
				return false, err
			}
		}

		_, err := e.CreateTreasurerSubmission(&request.CreateTreasurerSubmission{
			Type:              req.Type,
			VaultId:           vault.Id,
			SumitterProfileId: req.RequesterProfileId,
			Choice:            consts.TreasurerSubmissionStatusApproved,
			RequestId:         treasurerRequest.Id,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.CreateTreasurerSubmission failed")
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (e *Entity) validateTreasurer(treasurers []model.VaultTreasurer, userProfileId string) bool {
	for _, treasurer := range treasurers {
		if treasurer.UserProfileId == userProfileId {
			return true
		}
	}
	return false
}

func (e *Entity) validateBalance(token *mochipay.Token, address, solanaAddress, amount string) bool {
	var balance *big.Int
	var err error
	// validate balance token base
	if token.Chain.ChainId == "999" {
		balance, err = e.vaultwallet.BalanceSolana(token, solanaAddress)
		if err != nil {
			e.log.Fields(logger.Fields{"address": address, "amount": amount}).Errorf(err, "[entity.validateBalance] - e.vaultwallet.BalanceSolana failed")
			return false
		}
	} else {
		balance, err = e.vaultwallet.Balance(token, address)
		if err != nil {
			e.log.Fields(logger.Fields{"address": address, "amount": amount}).Errorf(err, "[entity.validateBalance] - e.vaultwallet.Balance failed")
			return false
		}
	}

	// check and validate balances
	amountBigIntStr := util.FloatToString(amount, token.Decimal)
	amountBigInt, err := util.StringToBigInt(amountBigIntStr)
	if err != nil {
		e.log.Fields(logger.Fields{"address": address, "amount": amount}).Errorf(err, "[entity.TransferVaultToken] - util.StringToBigInt failed")
		return false
	}

	cmp, err := util.CmpBigInt(balance, amountBigInt)
	if err != nil {
		e.log.Fields(logger.Fields{"address": address, "amount": amount}).Errorf(err, "[entity.TransferVaultToken] - util.CmpBigInt failed")
		return false
	}

	if cmp == -1 {
		e.log.Fields(logger.Fields{"address": address, "amount": amount}).Errorf(err, "[entity.TransferVaultToken] - balance not enough")
		return false
	}

	return true
}

func (e *Entity) CreateTreasurerSubmission(req *request.CreateTreasurerSubmission) (resp *response.CreateTreasurerSubmissionResponse, err error) {
	treasurerReq, err := e.repo.VaultRequest.GetById(req.RequestId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.Vault.GetById failed")
		return nil, err
	}

	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.Vault.GetById failed")
		return nil, err
	}

	modelSubmission := model.VaultSubmission{
		VaultId:            req.VaultId,
		RequestId:          req.RequestId,
		SubmitterProfileId: req.SumitterProfileId,
		Status:             req.Choice,
	}

	// get pending submission
	_, err = e.repo.VaultSubmission.GetPendingSubmission(&modelSubmission)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.VaultSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// update pending submission
	submission, err := e.repo.VaultSubmission.UpdatePendingSubmission(&modelSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.VaultSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// check if total submission >= threshold
	// get all submission of this vault
	submissions, err := e.repo.VaultSubmission.GetByRequestId(req.RequestId, req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.VaultSubmission.GetByRequestId failed")
		return nil, err
	}

	// check vault_request to see column is_approved = true, this mean it is already executed
	if treasurerReq.IsApproved {
		return &response.CreateTreasurerSubmissionResponse{
			Submission:       *submission,
			TotalSubmissions: submissions,
			VoteResult: response.VoteResult{
				IsApproved: false, // this mean we not allow this request to execute once more
			},
		}, nil
	}

	totalApprovedSubmission := 0
	totalRejectedSubmisison := 0
	for _, submission := range submissions {
		if submission.Status == consts.TreasurerSubmissionStatusApproved {
			totalApprovedSubmission++
		}
		if submission.Status == consts.TreasurerSubmissionStatusRejected {
			totalRejectedSubmisison++
		}
	}

	submission.GuildId = submissions[0].GuildId
	submission.Vault = submissions[0].Vault
	threshold, _ := strconv.ParseFloat(submissions[0].Vault.Threshold, 64)
	percentage := float64(totalApprovedSubmission) / float64(len(submissions)) * 100
	allowedRejectVote := int64(len(submissions)) - int64(math.Ceil(float64(len(submissions))*threshold/100))

	resp = &response.CreateTreasurerSubmissionResponse{
		Submission: *submission,
		VoteResult: response.VoteResult{
			IsApproved:                false,
			TotalApprovedSubmission:   int64(totalApprovedSubmission),
			TotalRejectedSubmisison:   int64(totalRejectedSubmisison),
			AllowedRejectedSubmisison: allowedRejectVote,
			TotalVote:                 int64(totalApprovedSubmission + totalRejectedSubmisison),
			TotalSubmission:           int64(len(submissions)),
			Percentage:                fmt.Sprintf("%.2f", percentage),
			Threshold:                 fmt.Sprintf("%.2f", threshold),
			ThresholdNumber:           threshold,
		},
		TotalSubmissions: submissions,
	}

	if percentage >= threshold {
		resp.VoteResult.IsApproved = true
	}

	// noti for this submission of treasurer
	voteMessage, daoVaultTotalTreasurerProposal := e.formatVoteVaultMessage(req, resp, req.SumitterProfileId, treasurerReq.UserProfileId, vault, submissions, treasurerReq)
	byteNotification, _ := json.Marshal(voteMessage)

	err = e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteNotification)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.kafka.Produce failed")
		return nil, err
	}

	// noti result of this request to user
	time.Sleep(3 * time.Second)
	voteMessage.Type = typeset.NOTIFICATION_VAULT_PROPOSAL
	voteMessage.VaultVoteMetadata.DaoVaultTotalTreasurer = daoVaultTotalTreasurerProposal
	if resp.VoteResult.IsApproved {
		err = e.repo.VaultRequest.UpdateStatus(submission.RequestId, consts.TreasurerRequestStatusApproved)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.VaultRequest.UpdateStatus failed")
			return nil, err
		}

		voteMessage.VaultVoteMetadata.TreasurerVote = "pass"
		byteResultNotificationSuccess, _ := json.Marshal(voteMessage)
		err = e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteResultNotificationSuccess)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.kafka.Produce failed")
			return nil, err
		}

	} else {
		if int64(totalRejectedSubmisison) > allowedRejectVote {
			voteMessage.VaultVoteMetadata.TreasurerVote = "failed"
			byteResultNotificationFail, _ := json.Marshal(voteMessage)
			err = e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteResultNotificationFail)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.kafka.Produce failed")
				return nil, err
			}
		}
	}

	return resp, nil
}

func (e *Entity) RemoveTreasurerFromVault(req *request.RemoveTreasurerToVaultRequest) (*model.VaultTreasurer, error) {
	treasurer, err := e.repo.VaultTreasurer.Delete(&model.VaultTreasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserProfileId: req.UserProfileID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.RemoveTreasurerFromVault] - e.repo.VaultTreasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerRemoveType,
		Target:  req.UserProfileID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.RemoveTreasurerFromVault] - e.repo.VaultTransaction.Create failed")
		return nil, err
	}

	return treasurer, nil
}

func (e *Entity) GetVaultDetail(vaultName, guildId string) (*response.VaultDetailResponse, error) {
	vault, err := e.repo.Vault.GetByNameAndGuildId(vaultName, guildId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("vault not exist")
		}
		e.log.Fields(logger.Fields{"vaultName": vaultName}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetByNameAndGuildId failed")
		return nil, err
	}

	// get balance
	bal := make([]response.Balance, 0)
	bal, err = balanceVaultDetail(vault, bal)
	if err != nil {
		e.log.Fields(logger.Fields{"vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - balanceVaultDetail failed")
		return nil, err
	}

	// get treasurers
	treasurers, err := e.repo.VaultTreasurer.GetByGuildIdAndVaultId(guildId, vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.VaultTreasurer.GetByGuildIdAndVaultName failed")
		return nil, err
	}

	// get current request
	currentRequest, err := e.repo.VaultRequest.GetCurrentRequest(vault.Id, guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.VaultRequest.GetCurrentRequest failed")
		return nil, err
	}

	currentRequestResponse := make([]response.CurrentRequest, 0)
	for _, req := range currentRequest {
		totalApprovedSubmisison := 0
		for _, sub := range req.VaultSubmission {
			if sub.Status == consts.TreasurerSubmissionStatusApproved {
				totalApprovedSubmisison++
			}
		}
		currentRequestResponse = append(currentRequestResponse, response.CurrentRequest{
			Target:                  req.UserProfileId,
			Action:                  util.Capitalize(req.Type),
			Token:                   req.Token,
			Amount:                  req.Amount,
			Address:                 req.Address,
			TotalSubmission:         int64(len(req.VaultSubmission)),
			TotalApprovedSubmission: int64(totalApprovedSubmisison),
		})
	}

	// get recent transaction
	recentTransactions, err := e.repo.VaultTransaction.GetRecentTx(vault.Id, guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.VaultTransaction.GetRecentTx failed")
		return nil, err
	}
	recentTxResponse := make([]response.VaultTransaction, 0)
	for _, tx := range recentTransactions {
		recentTxResponse = append(recentTxResponse, response.VaultTransaction{
			Action:    util.Capitalize(strings.Replace(tx.Action, "_", " ", -1)),
			Target:    tx.Target,
			Date:      tx.CreatedAt,
			Amount:    tx.Amount,
			Token:     tx.Token,
			ToAddress: tx.ToAddress,
			Threshold: tx.Threshold,
		})
	}

	return &response.VaultDetailResponse{
		WalletAddress:       vault.WalletAddress,
		SolanaWalletAddress: vault.SolanaWalletAddress,
		EstimatedTotal:      "",
		Balance:             bal,
		MyNft:               []response.MyNft{},
		VaultTreasurer:      treasurers,
		RecentTransaction:   recentTxResponse,
		CurrentRequest:      currentRequestResponse,
		Threshold:           vault.Threshold,
	}, nil
}

func balanceVaultDetail(vault *model.Vault, bal []response.Balance) ([]response.Balance, error) {
	listAssetEvm, _, _, err := e.ListWalletAssets(request.ListWalletAssetsRequest{Address: vault.WalletAddress, Type: "eth"})
	if err != nil {
		e.log.Fields(logger.Fields{"vault": vault}).Errorf(err, "[entity.balanceVaultDetail] - e.ListWalletAssets failed")
		return nil, err
	}

	listAssetSol, _, _, err := e.ListWalletAssets(request.ListWalletAssetsRequest{Address: vault.SolanaWalletAddress, Type: "sol"})
	if err != nil {
		e.log.Fields(logger.Fields{"vault": vault}).Errorf(err, "[entity.balanceVaultDetail] - e.ListWalletAssets failed")
		return nil, err
	}

	list := append(listAssetEvm, listAssetSol...)
	for idx, itm := range list {
		if itm.Token.Symbol == "ICY" && itm.Token.Chain.Name == "matic" {
			itm.Token.Price = 1.5
			list[idx] = itm
		}
		bal = append(bal, response.Balance{
			Token:  itm.Token,
			Amount: itm.Amount,
		})
	}

	return bal, nil
}

func (e *Entity) GetTreasurerRequest(requestId string) (*model.VaultRequest, error) {
	requestIdInt, err := strconv.Atoi(requestId)
	if err != nil {
		e.log.Fields(logger.Fields{"requestId": requestId}).Errorf(err, "[entity.GetTreasurerRequest] - strconv.Atoi failed")
		return nil, err
	}

	return e.repo.VaultRequest.GetById(int64(requestIdInt))
}

func (e *Entity) GetVaultTransactions(query vaulttxquery.VaultTransactionQuery) ([]model.VaultTransaction, error) {
	vault, err := e.repo.Vault.GetById(query.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"query": query}).Errorf(err, "[entity.GetVaultTransactions] - e.repo.Vault.GetById failed")
		return nil, err
	}

	vaultTxs, err := e.repo.VaultTransaction.GetTransactionByVaultId(query)
	if err != nil {
		e.log.Fields(logger.Fields{"query": query}).Errorf(err, "[entity.GetVaultTransactions] - e.repo.VaultTransaction.GetTransactionByVaultId failed")
		return nil, err
	}

	for i, vaultTx := range vaultTxs {
		vaultTx.CreatedAt = vaultTx.CreatedAt.Truncate(time.Second)
		vaultTx.VaultName = vault.Name
		vaultTxs[i] = vaultTx
	}

	return vaultTxs, nil
}
