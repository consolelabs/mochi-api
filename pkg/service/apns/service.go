package apns

import "github.com/sideshow/apns2"

type Service interface {
	PushNotificationToIos(pushToken string, price float64, trend string, token string) (*apns2.Response, error)
}
