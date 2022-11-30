package apns

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

type APNSClient struct {
	client *apns2.Client
}

func NewService(cfg *config.Config) Service {
	decoded, _ := base64.StdEncoding.DecodeString(cfg.AppleAuthKey)
	authKey, _ := token.AuthKeyFromBytes(decoded)

	token := &token.Token{
		AuthKey: authKey,
		KeyID:   cfg.AppleKeyID,
		TeamID:  cfg.AppleTeamID,
	}
	c := apns2.NewTokenClient(token).Production()
	return &APNSClient{
		client: c,
	}
}

func (f *APNSClient) PushNotificationToIos(pushToken string, price float64, trend string, token string) (*apns2.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := fmt.Sprintf("%s has fallen to %v", token, price)
	if trend == "up" {
		msg = fmt.Sprintf("%s has reached %v", token, price)
	}
	notification := &apns2.Notification{
		DeviceToken: pushToken,
		Topic:       "so.console.mochi",
		Payload:     []byte(fmt.Sprintf(`{"aps":{"alert":"%s"}}`, msg)),
	}

	res, err := f.client.PushWithContext(ctx, notification)
	if err != nil {
		return res, err
	}
	if !res.Sent() {
		return res, fmt.Errorf(res.Reason)
	}
	return res, nil
}
