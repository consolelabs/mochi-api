package entities

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/mochipay"
)

func (e *Entity) OnboardingStart(req request.OnboardingStartRequest) error {
	// Check onboarding status
	onboardingStatus, err := e.svc.MochiProfile.GetOnboardingStatus(req.ProfileId)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": req.ProfileId}).Error(err, "[Entity.OnboardingStart] svc.MochiProfile.GetOnboardingStatus() failed")
		return err
	}
	if onboardingStatus.DidOnboarding {
		// if user already start onboarding, just return
		return nil
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
		return err
	}
	decimal := math.BigPow(10, kekkToken.Decimal)
	rewardAmount := 10
	amount := new(big.Int).Mul(big.NewInt(int64(rewardAmount)), decimal).String()

	// Prepare application auth
	privateKey, err := hex.DecodeString(e.cfg.MochiAppPrivateKey)
	if err != nil {
		e.log.Error(err, "[Entity.OnboardingStart] hex.DecodeString() failed")
		return err
	}
	message := "ApplicationTransfer"
	signature := ed25519.Sign(privateKey, []byte(message))
	sigStr := hex.EncodeToString(signature)
	if err != nil {
		e.log.Error(err, "[Entity.OnboardingStart] hex.EncodeString() failed")
		return err
	}

	// Transfer reward token
	appTransferReq := mochipay.ApplicationTransferRequest{
		AppId: "35",
		Header: mochipay.ApplicationBaseHeaderRequest{
			Application: "Mochi",
			Message:     message,
			Signature:   sigStr,
		},
		Metadata: mochipay.ApplicationTransferMetadata{
			RecipientIds: []string{req.ProfileId},
			Amounts:      []string{amount},
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
		return err
	}

	// Mark user already started onboarding
	if err := e.svc.MochiProfile.MarkUserDidOnboarding(req.ProfileId); err != nil {
		e.log.
			Fields(logger.Fields{"profileId": req.ProfileId}).
			Error(err, "[Entity.OnboardingStart] svc.MochiProfile.MarkUserDidOnboarding() failed")
		return err
	}

	// Send notification to user
	if err := e.sendOnboardingStartNotification(); err != nil {
		e.log.
			Fields(logger.Fields{"profileId": req.ProfileId}).
			Error(err, "[Entity.OnboardingStart] e.sendOnboardingStartNotification() failed")
		return nil
	}

	return nil
}

func (e *Entity) sendOnboardingStartNotification() error {
	msg := struct{ Message string }{Message: "Hello"}
	byteNotification, _ := json.Marshal(msg)
	if err := e.kafka.ProduceNotification(e.cfg.Kafka.NotificationTopic, byteNotification); err != nil {
		e.log.Fields(logger.Fields{"req": ""}).Error(err, "[entity.sendOnboardingStartNotification] - e.kafka.Produce failed")
		return err
	}
	return nil
}
