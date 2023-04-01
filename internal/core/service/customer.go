package service

import (
	"context"
	"github.com/omid-h70/bank-service/internal/core/domain"
	"time"
)

type (
	// go:generate mockgen -destination=../../mocks/service/mockCustomerService.go -package=service github.com/omid-h70/bank-service/internal/core/service CustomerService
	CustomerService interface {
		GetMostActiveCustomersWithinTime(ctx context.Context, count int, time int) ([]domain.CustomerReportOut, error)
	}

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
