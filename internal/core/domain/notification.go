package domain

type PushNotificationService interface {
	SendNotifyMessage(sender string, receptor []string, msg string) error
}

type PushNotificationRepo interface {
	SendMessage(sender string, receptor []string, msg string) error
}
