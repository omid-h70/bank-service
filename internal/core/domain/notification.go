package domain

type PushNotificationService interface {
	SendSuccessfulMessage(msg string) error
	SendFailureMessage(msg string) error
}
