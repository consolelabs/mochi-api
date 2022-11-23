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

//firebase "firebase.google.com/go"
//"github.com/appleboy/go-fcm"

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

func (f *APNSClient) PushNotificationToIos(pushToken string, price float64, trend string, token string) error {
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
	if err != nil || !res.Sent() {
		return fmt.Errorf("failed to push err: %s | apns err: %s", err, res.Reason)
	}
	return nil
}
