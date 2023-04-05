package entities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) NewGuildConfigWalletVerificationMessage(req model.GuildConfigWalletVerificationMessage) (*model.GuildConfigWalletVerificationMessage, error) {

	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get discord guild: %v", err.Error())
	}

	_, err = e.repo.GuildConfigWalletVerificationMessage.GetOne(req.GuildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("failed to get guild config verification: %v", err.Error())
	}

	verificationMsg := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Verify",
						Style:    discordgo.PrimaryButton,
						CustomID: "mochi_verify",
					},
				},
			},
		},
	}

	switch {
	case req.EmbeddedMessage != nil:
		var embeddedMsg discordgo.MessageEmbed

		err = json.Unmarshal([]byte(req.EmbeddedMessage), &embeddedMsg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal embedded message %v: %v", req.EmbeddedMessage, err.Error())
		}

		verificationMsg.Embed = &embeddedMsg

	case req.Content != "":
		verificationMsg.Content = req.Content

	default:
		verificationMsg.Embed = &discordgo.MessageEmbed{
			Title:       ":robot: Verification required",
			Description: "Verify your wallet. This is a read-only connection. Do not share your private keys. We will never ask for your seed phrase. We will never DM you.",
			Color:       15240072,
		}
	}

	m, err := e.discord.ChannelMessageSendComplex(req.VerifyChannelID, verificationMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %v", err.Error())
	}

	req.DiscordMessageID = m.ID

	if err := e.repo.GuildConfigWalletVerificationMessage.UpsertOne(req); err != nil {
		return nil, fmt.Errorf("failed to upsert guild config verification: %v", err.Error())
	}

	return &req, nil
}

func (e *Entity) GetGuildConfigWalletVerificationMessage(guildId string) (*model.GuildConfigWalletVerificationMessage, error) {
	_, err := e.repo.DiscordGuilds.GetByID(guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[e.repo.DiscordGuilds.GetByID] - failed to get guild by id")
		return nil, err
	}

	res, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildId)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildId}).Error(err, "[e.repo.GuildConfigWalletVerificationMessage.GetOne] - failed to get config by guild id")
		return nil, err
	}

	return res, nil
}

func (e *Entity) UpdateGuildConfigWalletVerificationMessage(req model.GuildConfigWalletVerificationMessage) (*model.GuildConfigWalletVerificationMessage, error) {

	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get discord guild: %v", err.Error())
	}

	verificationMsg, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(req.GuildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get guild config verification: %v", err.Error())
	}

	var embeddedMsg discordgo.MessageEmbed

	if req.EmbeddedMessage == nil && req.Content == "" {
		embeddedMsg = discordgo.MessageEmbed{
			Title:       ":robot: Verification required",
			Description: "Verify your wallet. This is a read-only connection. Do not share your private keys. We will never ask for your seed phrase. We will never DM you.",
			Color:       15240072,
		}
	} else {
		err = json.Unmarshal([]byte(req.EmbeddedMessage), &embeddedMsg)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal embedded message %v: %v", req.EmbeddedMessage, err.Error())
		}
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Verify",
					Style:    discordgo.PrimaryButton,
					CustomID: "mochi_verify",
				},
			},
		},
	}

	switch {
	case req.VerifyChannelID != verificationMsg.VerifyChannelID:
		if verificationMsg.DiscordMessageID != "" {
			err = e.discord.ChannelMessageDelete(verificationMsg.VerifyChannelID, verificationMsg.DiscordMessageID)
			if err != nil {
				return nil, fmt.Errorf("failed to delete discord message: %v", err.Error())
			}
		}
		fallthrough

	case verificationMsg.DiscordMessageID == "":

		m, err := e.discord.ChannelMessageSendComplex(req.VerifyChannelID, &discordgo.MessageSend{
			Content:    req.Content,
			Embed:      &embeddedMsg,
			Components: components,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to send new verification message: %v", err.Error())
		}
		req.DiscordMessageID = m.ID

	default:
		_, err = e.discord.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:         verificationMsg.DiscordMessageID,
			Channel:    verificationMsg.VerifyChannelID,
			Content:    &req.Content,
			Embed:      &embeddedMsg,
			Components: components,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to edit verification message: %v", err.Error())
		}
		req.DiscordMessageID = verificationMsg.DiscordMessageID
	}

	if err := e.repo.GuildConfigWalletVerificationMessage.UpsertOne(req); err != nil {
		return nil, fmt.Errorf("failed to upsert guild config verification: %v", err.Error())
	}

	return &req, nil
}

func (e *Entity) DeleteGuildConfigWalletVerificationMessage(guildID string) error {

	verificationMsg, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildID)
	if err != nil {
		return fmt.Errorf("failed to get guild config verification message: %v", err.Error())
	}

	if verificationMsg.DiscordMessageID != "" {
		err = e.discord.ChannelMessageDelete(verificationMsg.VerifyChannelID, verificationMsg.DiscordMessageID)
		// case user deleted channel
		if err != nil && !strings.Contains(err.Error(), "Not Found") {
			return fmt.Errorf("failed to delete discord message: %v", err.Error())
		}
	}

	if err := e.repo.GuildConfigWalletVerificationMessage.DeleteOne(guildID); err != nil {
		return fmt.Errorf("failed to delete guild config verification: %v", err.Error())
	}

	return nil
}

func (e *Entity) GenerateVerification(req request.GenerateVerificationRequest) (data string, statusCode int, err error) {
	_, err = e.repo.GuildConfigWalletVerificationMessage.GetOne(req.GuildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", http.StatusBadRequest, fmt.Errorf("this guild has not set verification config")
		}
		return "", http.StatusInternalServerError, fmt.Errorf("failed to get guild config verification: %v", err.Error())
	}

	uw, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(req.UserDiscordID, req.GuildID)
	switch err {
	case nil:
		if !req.IsReverify {
			return uw.Address, http.StatusConflict, fmt.Errorf("already have a verified wallet")
		}
	case gorm.ErrRecordNotFound:
		if req.IsReverify {
			return "", http.StatusBadRequest, fmt.Errorf("unverified user")
		}
	default:
		return "", http.StatusInternalServerError, err
	}

	code := uuid.New().String()
	if err := e.repo.DiscordWalletVerification.UpsertOne(
		model.DiscordWalletVerification{
			Code:          code,
			UserDiscordID: req.UserDiscordID,
			GuildID:       req.GuildID,
			CreatedAt:     time.Now(),
		},
	); err != nil {
		return "", http.StatusInternalServerError, err
	}

	return code, http.StatusOK, nil
}

func (e *Entity) VerifyWalletAddress(req request.VerifyWalletAddressRequest) (int, error) {
	verification, err := e.repo.DiscordWalletVerification.GetByValidCode(req.Code)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid code")
	}

	if err := util.VerifySig(req.WalletAddress, req.Signature, fmt.Sprintf(
		"This will help us connect your discord account to the wallet address.\n\nMochiBotCode=%s", req.Code)); err != nil {
		return http.StatusBadRequest, err
	}

	// case add wallet
	if verification.GuildID == "" {
		err = e.handleWalletAddition(req.WalletAddress, *verification)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.VerifyWalletAddress] entity.handleWalletAddition() failed")
			return http.StatusInternalServerError, err
		}
		return http.StatusOK, nil
	}

	_, err = e.repo.Users.GetOne(verification.UserDiscordID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"userID": verification.UserDiscordID}).Error(err, "[entity.VerifyWalletAddress] repo.Users.GetOne() failed")
		return http.StatusInternalServerError, err
	}
	if err == gorm.ErrRecordNotFound {
		err = e.repo.Users.UpsertMany([]model.User{{ID: verification.UserDiscordID}})
		if err != nil {
			e.log.Fields(logger.Fields{"userID": verification.UserDiscordID}).Error(err, "[entity.VerifyWalletAddress] repo.Users.UpsertMany() failed")
			return http.StatusInternalServerError, err
		}
	}

	uw, err := e.repo.UserWallet.GetOneByGuildIDAndAddress(verification.GuildID, req.WalletAddress)
	switch err {
	case nil:
		if uw.UserDiscordID != verification.UserDiscordID {
			// this address is already used by another user in this guild
			return http.StatusBadRequest, fmt.Errorf("this wallet address already belong to another user")
		}

	case gorm.ErrRecordNotFound:
		if err := e.repo.UserWallet.UpsertOne(model.UserWallet{
			UserDiscordID: verification.UserDiscordID,
			GuildID:       verification.GuildID,
			Address:       req.WalletAddress,
		}); err != nil {
			return http.StatusInternalServerError, fmt.Errorf("failed to upsert user wallet: %v", err.Error())
		}

	default:
		return http.StatusInternalServerError, fmt.Errorf("failed to get user wallet: %v", err.Error())
	}

	// assign role to user if guild has config
	err = e.VerifyWalletAssignRole(verification, req)
	if err != nil {
		e.log.Fields(logger.Fields{
			"req": req,
		}).Error(err, "[entities.VerifyWalletAddress] failed to assign role")
		return http.StatusInternalServerError, err
	}

	if err := e.repo.DiscordWalletVerification.DeleteByCode(verification.Code); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to delete verification: %v", err.Error())
	}

	return http.StatusOK, nil
}

func (e *Entity) VerifyWalletAssignRole(verification *model.DiscordWalletVerification, req request.VerifyWalletAddressRequest) error {
	guildConfigVerification, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(verification.GuildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{
				"guild_id": verification.GuildID,
			}).Info("[entities.VerifyWalletAssignRole] Guild does not have config verification")
			return nil
		}
		e.log.Fields(logger.Fields{
			"guild_id": verification.GuildID,
		}).Error(err, "[entities.VerifyWalletAssignRole] Failed to get guild config verification")
		return err
	}

	// check guild has config role for verify wallet, if yes then assign role, if not return
	if guildConfigVerification.VerifyRoleID != "" {
		// assign role to user
		err := e.discord.GuildMemberRoleAdd(verification.GuildID, verification.UserDiscordID, guildConfigVerification.VerifyRoleID)
		if err != nil {
			// allow acceptable error like bot not have access to assign role
			if util.IsAcceptableErr(err) {
				e.log.Fields(logger.Fields{
					"guild_id":        verification.GuildID,
					"user_discord_id": verification.UserDiscordID,
					"verify_role_id":  guildConfigVerification.VerifyRoleID,
				}).Infof("[entities.VerifyWalletAssignRole] Acceptable errors: %v", err)
				return nil
			}
			e.log.Fields(logger.Fields{
				"guild_id":        verification.GuildID,
				"user_discord_id": verification.UserDiscordID,
				"verify_role_id":  guildConfigVerification.VerifyRoleID,
			}).Error(err, "[entities.VerifyWalletAssignRole] Failed to assign role to user")
			return err
		}
	}
	return nil
}

func (e *Entity) handleWalletAddition(walletAddress string, verification model.DiscordWalletVerification) error {
	req := request.TrackWalletRequest{
		UserID:    verification.UserDiscordID,
		Address:   walletAddress,
		Type:      "eth",
		IsOwner:   true,
		MessageID: verification.MessageID,
		ChannelID: verification.ChannelID,
	}
	err := e.TrackWallet(req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.handleWalletAddition] entity.TrackWallet() failed")
		return err
	}
	err = e.repo.DiscordWalletVerification.DeleteByCode(verification.Code)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.handleWalletAddition] repo.DiscordWalletVerification.DeleteByCode() failed")
		return err
	}
	return nil
}

func (e *Entity) AssignVerifiedRole(userDiscordID, guildID string) error {
	guildConfigVerification, err := e.repo.GuildConfigWalletVerificationMessage.GetOne(guildID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{
				"guild_id": guildID,
			}).Info("[entities.AssignVerifiedRole] Guild does not have config verification")
			return nil
		}
		e.log.Fields(logger.Fields{
			"guild_id": guildID,
		}).Error(err, "[entities.AssignVerifiedRole] Failed to get guild config verification")
		return err
	}

	// Skip assign if guild does not have config verified role
	if guildConfigVerification.VerifyRoleID == "" {
		e.log.Fields(logger.Fields{
			"guild_id": guildID,
		}).Info("[entities.AssignVerifiedRole] Guild does not have config verification channel")
		return nil
	}

	// Get user mochi profile
	profile, err := e.svc.MochiProfile.GetByDiscordID(userDiscordID)
	if err != nil {
		e.log.Fields(logger.Fields{
			"guild_id":        guildID,
			"user_discord_id": userDiscordID,
		}).Error(err, "[entities.AssignVerifiedRole] Failed to get user profile")
		return err
	}

	shouldAssignRole := false
	for _, acc := range profile.AssociatedAccounts {
		if acc.Platform == mochiprofile.PlatformEVM || acc.Platform == mochiprofile.PlatformSol {
			shouldAssignRole = true
			break
		}
	}

	if !shouldAssignRole {
		err = fmt.Errorf("user does not have any verified wallet")
		e.log.Fields(logger.Fields{
			"discord_id": userDiscordID,
			"guild_id":   guildID,
			"profile_id": profile.ID,
		}).Error(err, "[entities.AssignVerifiedRole] user does not have any verified wallet")
		return err
	}

	// assign role to user
	err = e.discord.GuildMemberRoleAdd(guildID, userDiscordID, guildConfigVerification.VerifyRoleID)
	if err != nil {
		// allow acceptable error like bot not have access to assign role
		if util.IsAcceptableErr(err) {
			e.log.Fields(logger.Fields{
				"guild_id":        guildID,
				"user_discord_id": userDiscordID,
				"verify_role_id":  guildConfigVerification.VerifyRoleID,
			}).Infof("[entities.VerifyWalletAssignRole] Acceptable errors: %v", err)
			return nil
		}
		e.log.Fields(logger.Fields{
			"guild_id":        guildID,
			"user_discord_id": userDiscordID,
			"verify_role_id":  guildConfigVerification.VerifyRoleID,
		}).Error(err, "[entities.VerifyWalletAssignRole] Failed to assign role to user")
		return err
	}
	return nil
}
