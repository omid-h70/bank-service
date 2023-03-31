package service

import (
	"encoding/xml"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"log"
	"os"
	"time"
)

type (
	PushNotificationServiceImpl struct {
		repo       domain.PushNotificationRepo
		ctxTimeout time.Duration
	}
)

func NewPushNotificationService(repo domain.PushNotificationRepo, t time.Duration) PushNotificationServiceImpl {
	return PushNotificationServiceImpl{
		repo:       repo,
		ctxTimeout: t,
	}
}

func (notify PushNotificationServiceImpl) SendNotifyMessage(sender string, receptor []string, msg string) error {
	return notify.repo.SendMessage(sender, receptor, msg)
}

func (notify PushNotificationServiceImpl) GetSenderNotifyMessage() string {
	return ""
}

func (notify PushNotificationServiceImpl) GetReceiverNotifyMessage() string {
	return ""
}

func ReadTemplateFile(filePath string) {
	body, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	t := domain.MessageTemplate{}
	xml.Unmarshal(body, &t)
	fmt.Println(t, string(body))
}
