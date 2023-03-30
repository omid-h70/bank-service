package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"github.com/omid-h70/bank-service/internal/core/helper"
	"time"
)

type (
	CustomerServiceImpl struct {
		repo       domain.CustomerRepository
		ctxTimeout time.Duration
	}
)

// NewCustomerService creates new NewReportService with its dependencies
func NewCustomerService(repo domain.CustomerRepository, t time.Duration) CustomerServiceImpl {
	return CustomerServiceImpl{
		repo:       repo,
		ctxTimeout: t,
	}
}

func (t CustomerServiceImpl) GetMostActiveCustomersWithinTime(ctx context.Context, count int, time int) error {
	err, customerList := t.repo.FindMostActiveCustomersWithinTime(context.Background(), 5, 5)
	helper.GO_UNUSED(err, customerList)
	return nil
}
