package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/core/domain"
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

func (t CustomerServiceImpl) GetMostActiveCustomersWithinTime(ctx context.Context, count int, timeMinute int) ([]domain.CustomerReportOut, error) {
	return t.repo.FindMostActiveCustomersWithinTime(context.Background(), count, timeMinute)
}
