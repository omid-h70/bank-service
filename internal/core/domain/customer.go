package domain

import "context"

type Customer struct {
	Id       string `json:"customer_id"`
	PhoneNum string `json:"phone_num"`
}

type CustomerService interface {
	GetMostActiveCustomersWithinTime(ctx context.Context, count int, time int) error
}

type CustomerReportOut struct {
}

type CustomerRepository interface {
	FindMostActiveCustomersWithinTime(ctx context.Context, count int, time int) ([]CustomerReportOut, error)
}
