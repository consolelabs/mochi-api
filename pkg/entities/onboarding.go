package entities

import (
	"crypto/ed25519"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/mochipay"
)

func (e *Entity) OnboardingStart(req request.OnboardingStartRequest) (*response.OnboardingStartData, error) {
	// Check onboarding status
	onboardingStatus, err := e.svc.MochiProfile.GetOnboardingStatus(req.ProfileId)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": req.ProfileId}).Error(err, "[Entity.OnboardingStart] svc.MochiProfile.GetOnboardingStatus() failed")
		return nil, err
	}
	if onboardingStatus.DidOnboarding {
		return &response.OnboardingStartData{
			UserAlreadyStarted: true,
			Reward:             nil,
		}, nil
	}

	// Get reward token info
	var (
		symbol  = "KEKK"
		chainId = "1"
	)
	kekkToken, err := e.svc.MochiPay.GetToken(symbol, chainId)
	if err != nil {
		e.log.
			Fields(logger.Fields{"symbol": symbol, "chainId": chainId}).
			Error(err, "[Entity.OnboardingStart] svc.MochiPay.GetToken() failed")
		return nil, err
	}

	// Prepare application auth
	privateKey, err := hex.DecodeString(e.cfg.MochiAppPrivateKey)
	if err != nil {
		e.log.Error(err, "[Entity.OnboardingStart] hex.DecodeString() failed")
		return nil, err
	}
	message := strconv.FormatInt(time.Now().Unix(), 10)
	signature := ed25519.Sign(privateKey, []byte(message))
	sigStr := hex.EncodeToString(signature)
	if err != nil {
		e.log.Error(err, "[Entity.OnboardingStart] hex.EncodeString() failed")
		return nil, err
	}

	// Transfer reward token
	rewardAmount := "10"
	appTransferReq := mochipay.ApplicationTransferRequest{
		AppId: "35",
		Header: mochipay.ApplicationBaseHeaderRequest{
			Application: "Mochi",
			Message:     message,
			Signature:   sigStr,
		},
		Metadata: mochipay.ApplicationTransferMetadata{
			RecipientIds: []string{req.ProfileId},
			Amounts:      []string{rewardAmount},
			TokenId:      kekkToken.Id,
			References:   "",
			Description:  "User onboarding reward",
			Platform:     req.Platform,
		},
	}
	_, err = e.svc.MochiPay.ApplicationTransfer(appTransferReq)
	if err != nil {
		e.log.
			Fields(logger.Fields{"appTransferRequest": appTransferReq}).
			Error(err, "[Entity.OnboardingStart] svc.MochiPay.ApplicationTransfer() failed")
		return nil, err
	}

	// Mark user already started onboarding
	if err := e.svc.MochiProfile.MarkUserDidOnboarding(req.ProfileId); err != nil {
		e.log.
			Fields(logger.Fields{"profileId": req.ProfileId}).
			Error(err, "[Entity.OnboardingStart] svc.MochiProfile.MarkUserDidOnboarding() failed")
		return nil, err
	}

	return &response.OnboardingStartData{
		UserAlreadyStarted: false,
		Reward: &response.OnboardingStartReward{
			Token:  kekkToken.Symbol,
			Amount: rewardAmount,
		},
	}, nil
}
