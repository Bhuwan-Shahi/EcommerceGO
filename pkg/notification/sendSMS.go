package notification

import "ecommerceGO/config"

type NotificationClient interface {
	SendSMS(phone, message string) error
}

type notificationClient struct {
	config config.Config
}

func (c notificationClient) SendSMS(phone, message string) error {
	return nil
}

func NewNotificationCleint(config config.Config) NotificationClient {
	return &notificationClient{
		config: config,
	}
}
