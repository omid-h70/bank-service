package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"time"
)

type (
	TransactionServiceImpl struct {
		notifyHandler domain.PushNotificationService
		repo          domain.TransactionRepository
		ctxTimeout    time.Duration
	}
)

// NewTransferService creates new NewTransferService with its dependencies
func NewTransferService(a domain.TransactionRepository, t time.Duration) TransactionServiceImpl {
	return TransactionServiceImpl{
		repo:       a,
		ctxTimeout: t,
	}
}

func (t TransactionServiceImpl) ExecuteCardTransfer(ctx context.Context, input domain.Transaction) error {
	err := t.repo.MakeTransferFromCardToCard(ctx, input)
	if err != nil {
		//t.notifyHandler.SendSuccessfulMessage("done")
	}
	return err
}
