package apns

type Service interface {
	PushNotificationToIos(pushToken string, price float64, trend string) error
}
