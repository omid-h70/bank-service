package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"time"
)

type (
	TransactionService interface {
		ExecuteCardTransfer(context.Context, domain.Transaction) ([]domain.AccountInfoOutput, error)
	}

	TransactionServiceImpl struct {
		notifyHandler domain.PushNotificationService
		repo          domain.TransactionRepository
		ctxTimeout    time.Duration
	}
)

func NewTransactionService(a domain.TransactionRepository, t time.Duration) TransactionServiceImpl {
	return TransactionServiceImpl{
		repo:       a,
		ctxTimeout: t,
	}
}

func (t TransactionServiceImpl) ExecuteCardTransfer(ctx context.Context, input domain.Transaction) ([]domain.AccountInfoOutput, error) {
	return t.repo.MakeTransferFromCardToCard(ctx, input)
}
