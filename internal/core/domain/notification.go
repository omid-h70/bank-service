package domain

type PushNotificationService interface {
	GetReceiverNotifyMessage() string
	GetSenderNotifyMessage() string
	SendNotifyMessage(sender string, receptor []string, msg string) error
}

type PushNotificationRepo interface {
	SendMessage(sender string, receptor []string, msg string) error
}

type MessageTemplate struct {
	Texts    []string `xml:"String"`
	FromCard string   `xml:"FromCard"`
	ToCard   string   `xml:"ToCard"`
	Amount   string   `xml:"Amount"`
}
