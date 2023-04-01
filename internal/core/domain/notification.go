package domain

const (
	FromCardNotification = iota
	ToCardNotification
)

type PushNotificationService interface {
	GetReceiverNotifyMessage(info AccountInfoOutput, templatePath string) string
	GetSenderNotifyMessage(info AccountInfoOutput, templatePath string) string
	SendNotifyMessage(sender string, receptor []string, msg string) error
}

type PushNotificationRepo interface {
	SendMessage(sender string, receptor []string, msg string) error
}

type SenderMessageTemplate struct {
	Texts    []string `xml:"String"`
	FromCard string   `xml:"FromCard"`
	Amount   string   `xml:"Amount"`
	Balance  string   `xml:"Balance"`
}

type ReceiverMessageTemplate struct {
	Texts   []string `xml:"String"`
	ToCard  string   `xml:"ToCard"`
	Amount  string   `xml:"Amount"`
	Balance string   `xml:"Balance"`
}
