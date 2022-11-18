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
	c := apns2.NewTokenClient(token).Development()
	return &APNSClient{
		client: c,
	}
}

func (f *APNSClient) PushNotificationToIos(pushToken string, price float64, trend string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	notification := &apns2.Notification{
		DeviceToken: pushToken,
		Topic:       "console.mochi.alert",
		Payload:     []byte(fmt.Sprintf(`{"aps":{"alert":"Price alert: %v %s"}}`, price, trend)),
	}
	_, err := f.client.PushWithContext(ctx, notification)
	if err != nil {
		return fmt.Errorf("failed to push notification: %s", err)
	}
	return nil
}
