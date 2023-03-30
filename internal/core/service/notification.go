package service

import (
	"github.com/omid-h70/bank-service/internal/core/domain"
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
	return notify.SendNotifyMessage(sender, receptor, msg)
}
