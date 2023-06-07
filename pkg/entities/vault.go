package entities

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
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

	// default for vault creator will be added as treasurer
	_, err = e.repo.Treasurer.Create(&model.Treasurer{
		VaultId:       vault.Id,
		GuildId:       req.GuildId,
		UserDiscordId: req.VaultCreator,
		Role:          consts.VaultCreatorRole,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateVault] - add treasurer failed")
		return nil, err
	}

	return vault, nil
}

func (e *Entity) GetVaults(req request.GetVaultsRequest) ([]model.Vault, error) {
	listQuery := vault.ListQuery{
		GuildID:      req.GuildID,
		EvmWallet:    req.EvmAddress,
		SolanaWallet: req.SolanaAddress,
		Threshold:    req.Threshold,
	}

	// find discord ID by given profile ID
	if req.ProfileID != "" {
		profile, err := e.svc.MochiProfile.GetByID(req.ProfileID)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetVaults] svc.MochiProfile.GetByID() failed")
			return nil, err
		}

		for _, acc := range profile.AssociatedAccounts {
			if acc.Platform == mochiprofile.PlatformDiscord {
				listQuery.UserDiscordID = acc.PlatformIdentifier
			}
		}
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
				return nil, err
			}
			vaults[i].TotalAmountSolana = fmt.Sprintf("%.4f", sumBal(walletAssetsSolana))
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

func (e *Entity) GetVaultInfo() (*model.VaultInfo, error) {
	return e.repo.VaultInfo.Get()
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

func (e *Entity) AddTreasurerToVault(req *request.AddTreasurerToVaultRequest) (*model.Treasurer, error) {
	treasurer, err := e.repo.Treasurer.Create(&model.Treasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserDiscordId: req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerAddType,
		Target:  req.UserDiscordID,
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

	treasurer, err := e.repo.Treasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Treasurer.GetByVaultId failed")
		return err
	}

	listNotify := []string{}
	for _, t := range treasurer {
		profileMember, err := e.svc.MochiProfile.GetByDiscordID(t.UserDiscordId, true)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Profile.GetByDiscordId failed")
			return err
		}
		listNotify = append(listNotify, profileMember.ID)
	}

	token, err := e.svc.MochiPay.GetToken(req.Token, req.Chain)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.svc.MochiPay.GetToken failed")
		return err
	}

	treasurerRequest, err := e.repo.TreasurerRequest.GetById(req.RequestId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.TreasurerRequest.GetById failed")
		return err
	}

	profile, err := e.svc.MochiProfile.GetByDiscordID(treasurerRequest.UserDiscordId, true)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - e.repo.Profile.GetByDiscordId failed")
		return err
	}

	if !slices.Contains(listNotify, profile.ID) {
		listNotify = append(listNotify, profile.ID)
	}

	amountBigIntStr := util.FloatToString(req.Amount, token.Decimal)

	validateBalance := e.validateBalance(token, vault.WalletAddress, vault.SolanaWalletAddress, req.Amount)
	if !validateBalance {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.TransferVaultToken] - validateBalance failed")
		return fmt.Errorf("balance not enough")
	}

	recipientPay := profile.ID
	if recipientPay == "" {
		recipientPay = treasurerRequest.Requester
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
		ProfileId:  profile.ID,
		Amount:     amountBigIntStr,
		To:         destination,
		PrivateKey: privateKey,
		Token:      token.Symbol,
		Chain:      token.Chain.ChainId,
		Name:       vault.Name,
		Requester:  recipientPay,
		Message:    treasurerRequest.Message,
		ListNotify: listNotify,
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
		Sender:    recipientPay,
		Target:    req.Target,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.VaultTransaction.Create failed")
		return err
	}

	return nil
}

func (e *Entity) CreateTreasurerResult(req *request.CreateTreasurerResultRequest) error {
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Vault.GetById failed")
		return err
	}

	action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)

	msg := prepareMessageNotifyTreasurerResult(req, action, vault.Name, thumbnail)

	err = sendNotifyTreasurerResult(msg, req.ChannelId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - sendNotifyTreasurerResult failed")
		return err
	}

	return nil
}

func prepareParamNotifyTreasurerResult(notifyType string) (action, thumbnail string) {
	action = consts.TreasurerAddedAction
	thumbnail = "https://cdn.discordapp.com/attachments/1090195482506174474/1092703907911847976/image.png"
	if notifyType == consts.TreasurerRemoveType {
		action = consts.TreasurerRemovedAction
		thumbnail = "https://cdn.discordapp.com/attachments/1090195482506174474/1092755046556516394/image.png"
	} else if notifyType == consts.TreasurerTransferType {
		action = consts.TreasurerTransferType
		thumbnail = "https://cdn.discordapp.com/attachments/1003381172178530494/1105400697836556368/vault_open.gif"
	}
	return action, thumbnail
}

func prepareMessageNotifyTreasurerResult(req *request.CreateTreasurerResultRequest, action, vaultName, thumbnail string) (msg discordgo.MessageSend) {
	destination := fmt.Sprintf("`%s`", util.ShortenAddress(req.Address))
	if req.Address == "" {
		destination = fmt.Sprintf("<@%s>", req.UserDiscordID)
	}

	if req.Status == consts.TreasurerStatusSuccess {
		description := fmt.Sprintf("<@%s> has been %s to **%s vault**", req.UserDiscordID, action, vaultName)
		title := fmt.Sprintf("<:check:1077631110047080478> Treasurer was successfully %s", action)
		if action == consts.TreasurerTransferType {
			description = fmt.Sprintf("%s %s %s has been sent to %s\nWe will notify you when all done.", util.TokenEmoji(strings.ToUpper(req.Token)), req.Amount, strings.ToUpper(req.Token), destination)
			title = "<:check:1077631110047080478> Transfer was successfullly submitted"
		}

		msg = discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       title,
					Description: description,
					Color:       0x5CD97D,
					Timestamp:   time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
	} else {
		description := fmt.Sprintf("<@%s> has not been %s to **%s vault**", req.UserDiscordID, action, vaultName)
		if action == consts.TreasurerTransferType {
			description = fmt.Sprintf("%s %s %s has not been sent to %s", util.TokenEmoji(strings.ToUpper(req.Token)), req.Amount, strings.ToUpper(req.Token), destination)
		}
		msg = discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("<:revoke:1077631119073230970> Treasurer was not %s", action),
					Description: description,
					Color:       0xD94F4F,
					Timestamp:   time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
	}
	return msg
}

func sendNotifyTreasurerResult(msg discordgo.MessageSend, channelId string) error {
	err := e.svc.Discord.SendMessage(channelId, msg)
	if err != nil {
		e.log.Fields(logger.Fields{"msg": msg, "channelId": channelId}).Errorf(err, "[entity.AddTreasurerToVault] - e.svc.Discord.SendMessage failed")
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

	treasurers, err := e.repo.Treasurer.GetByVaultId(vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.GetByVaultId failed")
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
		if !e.validateTreasurer(treasurers, req.UserDiscordId) {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerRequest] - user not in list treasurers")
			return nil, fmt.Errorf("user not in list treasurers")
		}
	}

	// create treasurer request
	treasurerReq, err := e.repo.TreasurerRequest.Create(&model.TreasurerRequest{
		GuildId:       req.GuildId,
		VaultId:       vault.Id,
		UserDiscordId: req.UserDiscordId,
		Message:       req.Message,
		Requester:     req.Requester,
		Type:          req.Type,
		Amount:        req.Amount,
		Chain:         req.Chain,
		Token:         req.Token,
		Address:       req.Address,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	// add submission with status pending for all treasurer in vaul
	treasurerSubmission := make([]model.TreasurerSubmission, 0)

	for _, treasurer := range treasurers {
		status := consts.TreasurerSubmissionStatusPending
		if treasurer.UserDiscordId == req.Requester {
			status = consts.TreasurerSubmissionStatusApproved
		}

		treasurerSubmission = append(treasurerSubmission, model.TreasurerSubmission{
			VaultId:   vault.Id,
			GuildId:   req.GuildId,
			RequestId: treasurerReq.Id,
			Status:    status,
			Submitter: treasurer.UserDiscordId,
		})
	}

	err = e.repo.TreasurerSubmission.Create(treasurerSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.AddTreasurerToVault] - e.repo.TreasurerSubmission.Create failed")
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
		Treasurer:            treasurers,
		IsDecidedAndExecuted: isDecidedAndExecuted,
	}, nil
}

func (e *Entity) PostCreateTreasurerRequest(req *request.CreateTreasurerRequest, treasurerRequest *model.TreasurerRequest, vault *model.Vault, treasurers []model.Treasurer) (bool, error) {
	threshold, _ := strconv.ParseFloat(vault.Threshold, 64)
	percentage := float64(1) / float64(len(treasurers)) * 100

	if percentage >= threshold {
		// execute action
		switch req.Type {
		case "add":
			_, err := e.AddTreasurerToVault(&request.AddTreasurerToVaultRequest{
				GuildId:       req.GuildId,
				VaultId:       vault.Id,
				UserDiscordID: req.UserDiscordId,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.AddTreasurerToVault failed")
				return false, err
			}
		case "remove":
			_, err := e.RemoveTreasurerFromVault(&request.AddTreasurerToVaultRequest{
				GuildId:       req.GuildId,
				VaultId:       vault.Id,
				UserDiscordID: req.UserDiscordId,
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
				Target:    req.UserDiscordId,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.TransferVaultToken failed")
				return false, err
			}
		}

		_, err := e.CreateTreasurerSubmission(&request.CreateTreasurerSubmission{
			Type:      req.Type,
			VaultId:   vault.Id,
			Sumitter:  req.Requester,
			Choice:    consts.TreasurerSubmissionStatusApproved,
			RequestId: treasurerRequest.Id,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.PostCreateTreasurerRequest] - e.CreateTreasurerSubmission failed")
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (e *Entity) validateTreasurer(treasurers []model.Treasurer, userDiscordId string) bool {
	for _, treasurer := range treasurers {
		if treasurer.UserDiscordId == userDiscordId {
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
	vault, err := e.repo.Vault.GetById(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.Vault.GetById failed")
		return nil, err
	}

	modelSubmission := model.TreasurerSubmission{
		VaultId:   req.VaultId,
		RequestId: req.RequestId,
		Submitter: req.Sumitter,
		Status:    req.Choice,
	}

	// get pending submission
	_, err = e.repo.TreasurerSubmission.GetPendingSubmission(&modelSubmission)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// update pending submission
	submission, err := e.repo.TreasurerSubmission.UpdatePendingSubmission(&modelSubmission)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetPendingSubmission failed")
		return nil, err
	}

	// check if total submission >= threshold
	// get all submission of this vault
	submissions, err := e.repo.TreasurerSubmission.GetByRequestId(req.RequestId, req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerSubmission.GetByRequestId failed")
		return nil, err
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
		},
		TotalSubmissions: submissions,
	}

	if percentage >= threshold {
		resp.VoteResult.IsApproved = true
	}

	// notify treasurer about process voting
	treasurers, err := e.repo.Treasurer.GetByVaultId(req.VaultId)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.Treasurer.GetByVaultId failed")
		return nil, err
	}
	for _, treasurer := range treasurers {
		msg := discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "<:bell:1087564962941124679> Mochi notifications",
					Description: fmt.Sprintf("<@%s> %s the request #%d in %s vault. This request will be approved if `%d/%d` treasurers approve (%s)",
						req.Sumitter,
						req.Choice,
						req.RequestId,
						vault.Name,
						len(submissions)-int(allowedRejectVote),
						len(submissions),
						vault.Threshold+"%",
					),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Approved",
							Value:  fmt.Sprintf("<:check:1077631110047080478> `%d/%d`", totalApprovedSubmission, len(submissions)),
							Inline: true,
						},
						{
							Name:   "Rejected",
							Value:  fmt.Sprintf("<:revoke:1077631119073230970> `%d`", totalRejectedSubmisison),
							Inline: true,
						},
						{
							Name:   "Waiting",
							Value:  fmt.Sprintf("<:clock:1080757110146605086> `%d`", len(submissions)-totalApprovedSubmission-totalRejectedSubmisison),
							Inline: true,
						},
					},
					Color: 0x34AAFF,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL: "https://cdn.discordapp.com/attachments/1090195482506174474/1090905984299442246/image.png",
					},
					Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Type /feedback to report",
					},
				},
			},
		}
		err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
			continue
		}

		// DM result to user
		if resp.VoteResult.IsApproved {
			action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)
			msg := prepareMessageNotifyTreasurerResult(&request.CreateTreasurerResultRequest{Status: consts.TreasurerStatusSuccess, UserDiscordID: submissions[0].TreasurerRequest.UserDiscordId, Token: submissions[0].TreasurerRequest.Token, Amount: submissions[0].TreasurerRequest.Amount, Address: submissions[0].TreasurerRequest.Address}, action, submissions[0].Vault.Name, thumbnail)
			err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
				continue
			}
		} else {
			if int64(totalRejectedSubmisison) > allowedRejectVote {
				action, thumbnail := prepareParamNotifyTreasurerResult(req.Type)
				msg := prepareMessageNotifyTreasurerResult(&request.CreateTreasurerResultRequest{Status: consts.TreasurerStatusFail, UserDiscordID: submissions[0].TreasurerRequest.UserDiscordId, Token: submissions[0].TreasurerRequest.Token, Amount: submissions[0].TreasurerRequest.Amount, Address: submissions[0].TreasurerRequest.Address}, action, submissions[0].Vault.Name, thumbnail)
				err = e.svc.Discord.SendDM(treasurer.UserDiscordId, msg)
				if err != nil {
					e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.svc.Discord.SendDM failed")
					continue
				}
			}
		}
	}

	// update request status
	if resp.VoteResult.IsApproved {
		err = e.repo.TreasurerRequest.UpdateStatus(submission.RequestId, consts.TreasurerRequestStatusApproved)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateTreasurerSubmission] - e.repo.TreasurerRequest.UpdateStatus failed")
			return nil, err
		}
	}

	return resp, nil
}

func (e *Entity) RemoveTreasurerFromVault(req *request.AddTreasurerToVaultRequest) (*model.Treasurer, error) {
	treasurer, err := e.repo.Treasurer.Delete(&model.Treasurer{
		GuildId:       req.GuildId,
		VaultId:       req.VaultId,
		UserDiscordId: req.UserDiscordID,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.RemoveTreasurerFromVault] - e.repo.Treasurer.Create failed")
		return nil, err
	}

	_, err = e.repo.VaultTransaction.Create(&model.VaultTransaction{
		GuildId: req.GuildId,
		VaultId: req.VaultId,
		Action:  consts.TreasurerRemoveType,
		Target:  req.UserDiscordID,
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
	treasurers, err := e.repo.Treasurer.GetByGuildIdAndVaultId(guildId, vault.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.Treasurer.GetByGuildIdAndVaultName failed")
		return nil, err
	}

	// get current request
	currentRequest, err := e.repo.TreasurerRequest.GetCurrentRequest(vault.Id, guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildId": guildId, "vaultName": vaultName}).Errorf(err, "[entity.GetVaultDetail] - e.repo.TreasurerRequest.GetCurrentRequest failed")
		return nil, err
	}

	currentRequestResponse := make([]response.CurrentRequest, 0)
	for _, req := range currentRequest {
		totalApprovedSubmisison := 0
		for _, sub := range req.TreasurerSubmission {
			if sub.Status == consts.TreasurerSubmissionStatusApproved {
				totalApprovedSubmisison++
			}
		}
		currentRequestResponse = append(currentRequestResponse, response.CurrentRequest{
			Target:                  req.UserDiscordId,
			Action:                  util.Capitalize(req.Type),
			Token:                   req.Token,
			Amount:                  req.Amount,
			Address:                 req.Address,
			TotalSubmission:         int64(len(req.TreasurerSubmission)),
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
		Treasurer:           treasurers,
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

func (e *Entity) GetTreasurerRequest(requestId string) (*model.TreasurerRequest, error) {
	requestIdInt, err := strconv.Atoi(requestId)
	if err != nil {
		e.log.Fields(logger.Fields{"requestId": requestId}).Errorf(err, "[entity.GetTreasurerRequest] - strconv.Atoi failed")
		return nil, err
	}

	return e.repo.TreasurerRequest.GetById(int64(requestIdInt))
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
